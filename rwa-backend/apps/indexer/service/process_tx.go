package service

import (
	"context"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/indexer/config"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/types"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ProcessTxService processes events with idempotency checks
type ProcessTxService struct {
	db                    *gorm.DB
	handlerMap            map[types.ContractType]map[string]EventHandler
	conf                  *config.Config
	latestResolvedEventId uint64
	initialized           bool
}

func NewProcessTxService(handlers []EventHandler, db *gorm.DB, conf *config.Config) *ProcessTxService {
	handlerMap := make(map[types.ContractType]map[string]EventHandler)
	for _, handler := range handlers {
		contractType := handler.ContractType()
		topic0 := handler.Topic0()

		if _, exists := handlerMap[contractType]; !exists {
			handlerMap[contractType] = make(map[string]EventHandler)
		}
		handlerMap[contractType][topic0] = handler
	}
	return &ProcessTxService{
		db:         db,
		conf:       conf,
		handlerMap: handlerMap,
	}
}

// InitCache initializes the service and loads the last processed event ID
func (s *ProcessTxService) InitCache(ctx context.Context) error {
	if s.initialized {
		return nil
	}

	q, u := gplus.NewQuery[rwa.EventClientRecord]()
	q.Eq(&u.ChainID, s.conf.Chain.ChainId)
	recordList, dbRes := gplus.SelectList(q, gplus.Db(s.db))
	if dbRes.Error != nil {
		log.ErrorZ(ctx, "failed to select event_client_record", zap.Uint64("chainId", s.conf.Chain.ChainId), zap.Error(dbRes.Error))
		return dbRes.Error
	}
	if len(recordList) > 0 {
		s.latestResolvedEventId = recordList[0].LastEventID
	} else {
		s.latestResolvedEventId = 0
	}
	s.initialized = true
	return nil
}

// ProcessEvent processes a single event with idempotency check using the provided transaction.
func (s *ProcessTxService) ProcessEvent(ctx context.Context, tx *gorm.DB, event *types.EventLogWithId, resolvedUpTo uint64) error {
	if event.EventId <= resolvedUpTo {
		log.InfoZ(ctx, "event already processed, skipping",
			zap.Uint64("eventId", event.EventId), zap.Uint64("resolvedUpTo", resolvedUpTo))
		return nil
	}

	if !lo.HasKey(s.handlerMap, event.ContractType) || !lo.HasKey(s.handlerMap[event.ContractType], event.Topics[0]) {
		log.InfoZ(ctx, "no hander for event",
			zap.String("txHash", event.TxHash), zap.String("blockNumber", event.BlockNumber),
			zap.Uint64("eventId", event.EventId), zap.String("contractType", string(event.ContractType)),
			zap.String("topic0", event.Topics[0]))
		return nil
	}

	handler := s.handlerMap[event.ContractType][event.Topics[0]]
	err := handler.HandleEvent(ctx, tx, event)
	if err != nil {
		log.ErrorZ(ctx, "failed to handle event",
			zap.String("txHash", event.TxHash), zap.String("blockNumber", event.BlockNumber),
			zap.Uint64("eventId", event.EventId), zap.Error(err))
		return err
	}

	log.InfoZ(ctx, "event processed successfully",
		zap.Uint64("eventId", event.EventId), zap.String("txHash", event.TxHash), zap.String("blockNumber", event.BlockNumber))

	return nil
}

// ProcessBatch processes all events and updates block progress in a single atomic transaction.
// Only after that transaction has actually committed does it call each handler's
// AfterCommit — this is what guarantees Alpaca (or any other external system)
// can never react to an order before that order is visible in the database.
func (s *ProcessTxService) ProcessBatch(ctx context.Context, events []*types.EventLogWithId, toBlock uint64, blockService *BlockService) error {
	maxEventID := s.latestResolvedEventId
	for _, event := range events {
		if event.EventId > maxEventID {
			maxEventID = event.EventId
		}
	}

	batchResolvedUpTo := s.latestResolvedEventId

	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, event := range events {
			err := s.ProcessEvent(ctx, tx, event, batchResolvedUpTo)
			if err != nil {
				return err
			}
			if event.EventId > batchResolvedUpTo {
				batchResolvedUpTo = event.EventId
			}
		}
		return blockService.UpdateLastProcessedBlockTx(ctx, tx, toBlock, maxEventID)
	})
	if err != nil {
		return err
	}

	// Transaction committed successfully — every row written above is now
	// visible to every other DB connection. Only now is it safe to run any
	// handler logic that talks to the outside world.
	s.latestResolvedEventId = batchResolvedUpTo

	for _, event := range events {
		if !lo.HasKey(s.handlerMap, event.ContractType) || !lo.HasKey(s.handlerMap[event.ContractType], event.Topics[0]) {
			continue
		}
		handler := s.handlerMap[event.ContractType][event.Topics[0]]
		if afterCommitErr := handler.AfterCommit(ctx, event); afterCommitErr != nil {
			// The DB write is already durably committed at this point — an
			// AfterCommit failure (e.g. Alpaca is down) must NOT be treated
			// as a batch failure, or we'd re-process an event whose row
			// already exists. Log it; recovery for this specific gap is the
			// reconciler's job (see the Pending+no-ExternalOrderID scan rule).
			log.ErrorZ(ctx, "AfterCommit failed for event; will need reconciler to pick this up",
				zap.Uint64("eventId", event.EventId), zap.String("txHash", event.TxHash), zap.Error(afterCommitErr))
		}
	}

	return nil
}
