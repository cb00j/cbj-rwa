package service

import (
	"context"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/indexer/config"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/types"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
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

	// Check event_client_record exists
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
