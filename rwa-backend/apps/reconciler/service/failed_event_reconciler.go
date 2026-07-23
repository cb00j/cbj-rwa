package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	trade "github.com/cb00j/cbj-rwa/rwa-backend/libs/core/trade"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	failedEventInterval  = 30 * time.Second
	maxFailedEventRetry  = 5
)

// rawTradeUpdate mirrors just the fields ProcessFillEvent needs out of
// Alpaca's trade_updates JSON payload (the same shape persistFailedEvent
// stored verbatim in failed_events.event_data). Keeping this local — rather
// than importing alpaca-stream's handlers package — is what lets reconciler
// stay a standalone app that doesn't depend on another app's internals.
type rawTradeUpdate struct {
	Event       string `json:"event"`
	ExecutionID string `json:"execution_id"`
	Timestamp   string `json:"timestamp"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
	Order       struct {
		ID             string `json:"id"`
		ClientOrderID  string `json:"client_order_id"`
		FilledAvgPrice string `json:"filled_avg_price"`
		FilledQty      string `json:"filled_qty"`
	} `json:"order"`
}

// FailedEventReconciler retries rows in failed_events — messages whose DB
// transaction failed the first time round (e.g. transient DB error), as
// opposed to OnchainReconciler which retries on-chain calls that failed
// *after* the DB write already succeeded.
type FailedEventReconciler struct {
	db      *gorm.DB
	settler *OnchainSettler
	stopCh  chan struct{}
}

func NewFailedEventReconciler(lc fx.Lifecycle, db *gorm.DB, settler *OnchainSettler) *FailedEventReconciler {
	s := &FailedEventReconciler{db: db, settler: settler, stopCh: make(chan struct{})}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			close(s.stopCh)
			return nil
		},
	})

	return s
}

func (s *FailedEventReconciler) run() {
	ctx := context.WithValue(context.Background(), log.TraceID, "failed_event_reconciler")
	ticker := time.NewTicker(failedEventInterval)
	defer ticker.Stop()

	log.InfoZ(ctx, "FailedEventReconciler started", zap.Duration("interval", failedEventInterval))

	for {
		select {
		case <-s.stopCh:
			log.InfoZ(ctx, "FailedEventReconciler stopped")
			return
		case <-ticker.C:
			s.scanAndRetry(ctx)
		}
	}
}

func (s *FailedEventReconciler) scanAndRetry(ctx context.Context) {
	q, u := gplus.NewQuery[rwa.FailedEvent]()
	q.Eq(&u.Status, "pending")
	events, dbRes := gplus.SelectList(q, gplus.Db(s.db.WithContext(ctx)))
	if dbRes.Error != nil {
		log.ErrorZ(ctx, "FailedEventReconciler: failed to query failed_events", zap.Error(dbRes.Error))
		return
	}
	if len(events) == 0 {
		return
	}
	log.InfoZ(ctx, "FailedEventReconciler: found pending failed events", zap.Int("count", len(events)))

	for _, event := range events {
		s.retryOne(ctx, event)
	}
}

func (s *FailedEventReconciler) retryOne(ctx context.Context, event *rwa.FailedEvent) {
	var raw rawTradeUpdate
	if err := json.Unmarshal([]byte(event.EventData), &raw); err != nil {
		s.markDead(ctx, event, "failed to unmarshal stored event_data: "+err.Error())
		return
	}

	input := trade.FillEventInput{
		ClientOrderID:   raw.Order.ClientOrderID,
		ExecutionID:     raw.ExecutionID,
		Price:           raw.Price,
		Qty:             raw.Qty,
		Timestamp:       raw.Timestamp,
		FilledAvgPrice:  raw.Order.FilledAvgPrice,
		FilledQty:       raw.Order.FilledQty,
		ExternalOrderID: raw.Order.ID,
		IsFull:          raw.Event == "fill",
	}

	order, err := trade.ProcessFillEvent(ctx, s.db, input)
	if err != nil {
		s.incrementRetry(ctx, event, err.Error())
		return
	}

	log.InfoZ(ctx, "FailedEventReconciler: reprocessed failed event successfully",
		zap.Uint64("failed_event_id", event.ID), zap.String("client_order_id", event.ClientOrderID),
		zap.String("resulting_status", string(order.Status)))

	s.markDone(ctx, event)

	// If this reprocessing brought the order to Filled, it still needs
	// on-chain settlement — hand off to the same settler OnchainReconciler
	// uses, so it gets picked up on the very next reconcile pass regardless.
	if order.Status == rwa.OrderStatusFilled {
		go s.settler.CallMarkExecuted(ctx, order)
	}
}

func (s *FailedEventReconciler) incrementRetry(ctx context.Context, event *rwa.FailedEvent, errMsg string) {
	newCount := event.RetryCount + 1
	status := "pending"
	if newCount >= maxFailedEventRetry {
		status = "dead"
		log.ErrorZ(ctx, "FailedEventReconciler: retry budget exhausted, marking dead",
			zap.Uint64("failed_event_id", event.ID), zap.String("client_order_id", event.ClientOrderID), zap.String("error", errMsg))
	} else {
		log.WarnZ(ctx, "FailedEventReconciler: retry failed, will try again later",
			zap.Uint64("failed_event_id", event.ID), zap.Int("attempt", newCount), zap.String("error", errMsg))
	}

	s.db.WithContext(ctx).Model(&rwa.FailedEvent{}).
		Where("id = ?", event.ID).
		Updates(map[string]interface{}{
			"retry_count":   newCount,
			"status":        status,
			"error_message": errMsg,
		})
}

func (s *FailedEventReconciler) markDone(ctx context.Context, event *rwa.FailedEvent) {
	s.db.WithContext(ctx).Model(&rwa.FailedEvent{}).
		Where("id = ?", event.ID).
		Update("status", "done")
}

func (s *FailedEventReconciler) markDead(ctx context.Context, event *rwa.FailedEvent, reason string) {
	log.ErrorZ(ctx, "FailedEventReconciler: event cannot be processed, marking dead",
		zap.Uint64("failed_event_id", event.ID), zap.String("reason", reason))
	s.db.WithContext(ctx).Model(&rwa.FailedEvent{}).
		Where("id = ?", event.ID).
		Updates(map[string]interface{}{"status": "dead", "error_message": reason})
}
