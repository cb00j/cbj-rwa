package service

import (
	"context"
	"time"

	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	reconcileInterval  = 30 * time.Second
	maxSettlementRetry = 5
)

// OnchainReconciler periodically scans for orders whose database status says
// settlement is done (cancelled/filled), but the corresponding on-chain tx
// hash is still empty — meaning the on-chain call actually failed or never
// fired. It retries via OnchainSettler with backoff, and flags an order for
// manual review once the retry budget is exhausted.
type OnchainReconciler struct {
	db       *gorm.DB
	settler  *OnchainSettler
	stopCh   chan struct{}
}

func NewOnchainReconciler(lc fx.Lifecycle, db *gorm.DB, settler *OnchainSettler) *OnchainReconciler {
	s := &OnchainReconciler{db: db, settler: settler, stopCh: make(chan struct{})}

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

func (s *OnchainReconciler) run() {
	ctx := context.WithValue(context.Background(), log.TraceID, "onchain_reconciler")
	ticker := time.NewTicker(reconcileInterval)
	defer ticker.Stop()

	log.InfoZ(ctx, "OnchainReconciler started", zap.Duration("interval", reconcileInterval))

	for {
		select {
		case <-s.stopCh:
			log.InfoZ(ctx, "OnchainReconciler stopped")
			return
		case <-ticker.C:
			s.scanAndRetry(ctx)
		}
	}
}

func (s *OnchainReconciler) scanAndRetry(ctx context.Context) {
	var stuck []rwa.Order

	err := s.db.WithContext(ctx).
		Where("needs_manual_review = ?", false).
		Where(
			s.db.Where("status = ? AND execute_tx_hash = ''", rwa.OrderStatusFilled).
				Or("status IN (?, ?) AND filled_quantity = 0 AND cancel_tx_hash = ''",
					rwa.OrderStatusCancelled, rwa.OrderStatusExpired).
				Or("status IN (?, ?) AND filled_quantity > 0 AND execute_tx_hash = ''",
					rwa.OrderStatusCancelled, rwa.OrderStatusExpired),
		).
		Find(&stuck).Error
	if err != nil {
		log.ErrorZ(ctx, "OnchainReconciler: failed to query stuck orders", zap.Error(err))
		return
	}
	if len(stuck) == 0 {
		return
	}
	log.InfoZ(ctx, "OnchainReconciler: found stuck orders", zap.Int("count", len(stuck)))

	for i := range stuck {
		order := &stuck[i]
		if !s.dueForRetry(order) {
			continue
		}
		s.retryOne(ctx, order)
	}
}

func (s *OnchainReconciler) dueForRetry(order *rwa.Order) bool {
	if order.LastSettlementAttemptAt == nil {
		return true
	}
	return time.Since(*order.LastSettlementAttemptAt) >= backoffDuration(order.SettlementAttempts)
}

func backoffDuration(attempts int) time.Duration {
	switch {
	case attempts <= 1:
		return 1 * time.Minute
	case attempts == 2:
		return 5 * time.Minute
	case attempts == 3:
		return 15 * time.Minute
	default:
		return 30 * time.Minute
	}
}

func (s *OnchainReconciler) retryOne(ctx context.Context, order *rwa.Order) {
	if order.SettlementAttempts >= maxSettlementRetry {
		s.markNeedsManualReview(ctx, order, "settlement retry budget exhausted")
		return
	}

	log.WarnZ(ctx, "OnchainReconciler: retrying stuck settlement",
		zap.Uint64("order_id", order.ID), zap.String("status", string(order.Status)),
		zap.Int("attempt", order.SettlementAttempts+1))

	// Same rule as alpaca-stream's handleTerminalState: any fill at all means
	// it must be settled via markExecuted, never a plain cancelOrder.
	if order.FilledQuantity.IsPositive() {
		s.settler.CallMarkExecuted(ctx, order)
	} else {
		s.settler.CallCancelOrder(ctx, order)
	}
}

func (s *OnchainReconciler) markNeedsManualReview(ctx context.Context, order *rwa.Order, reason string) {
	log.ErrorZ(ctx, "OnchainReconciler: order exceeded max retries, needs manual review",
		zap.Uint64("order_id", order.ID), zap.String("client_order_id", order.ClientOrderID), zap.String("reason", reason))

	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", order.ID).
		Update("needs_manual_review", true).Error; err != nil {
		log.ErrorZ(ctx, "OnchainReconciler: failed to flag order for manual review", zap.Error(err), zap.Uint64("order_id", order.ID))
	}
	// TODO: 接入真正的告警渠道(Slack / 企业微信 / PagerDuty)
}
