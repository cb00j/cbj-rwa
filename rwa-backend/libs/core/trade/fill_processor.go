package trade

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// FillEventInput is a deliberately flat, Alpaca-shape-agnostic version of a
// fill/partial_fill trade update. Both the live WS handler (which parses
// Alpaca's raw JSON) and the reconciler (which re-reads a stored
// failed_events row containing that same raw JSON) build one of these and
// hand it to ProcessFillEvent, so the actual business logic — VWAP,
// execution bookkeeping, status determination — lives in exactly one place.
type FillEventInput struct {
	ClientOrderID   string
	ExecutionID     string
	Price           string
	Qty             string
	Timestamp       string
	FilledAvgPrice  string // Alpaca's authoritative running total, if present
	FilledQty       string // Alpaca's authoritative running total, if present
	ExternalOrderID string
	IsFull          bool
}

// ProcessFillEvent applies a single fill/partial_fill event to the matching
// order inside one DB transaction, and returns the updated order so the
// caller can decide whether an on-chain settlement call is now needed
// (order.Status == Filled).
//
// Fixes applied while extracting this from alpaca-stream's
// handleFillOrPartialFill:
//  1. The idempotency check inverted its own condition (`if data.ExecutionID
//     == ""` instead of `!= ""`), meaning it only ever ran when there was NO
//     execution id to check against — duplicate fill events were never
//     actually caught. Fixed below.
//  2. Status (Filled vs PartiallyFilled) was computed BEFORE the authoritative
//     Alpaca filled_qty/filled_avg_price overwrite, so a stale locally-summed
//     quantity could produce a status that contradicted the final
//     FilledQuantity value written to the row. Status is now computed AFTER
//     the authoritative overwrite.
func ProcessFillEvent(ctx context.Context, db *gorm.DB, input FillEventInput) (*rwa.Order, error) {
	execPrice, err := decimal.NewFromString(input.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price %q: %w", input.Price, err)
	}
	execQty, err := decimal.NewFromString(input.Qty)
	if err != nil {
		return nil, fmt.Errorf("failed to parse qty %q: %w", input.Qty, err)
	}

	var result *rwa.Order

	txErr := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		q, u := gplus.NewQuery[rwa.Order]()
		q.Eq(&u.ClientOrderID, input.ClientOrderID)
		order, dbRes := gplus.SelectOne(q, gplus.Db(tx))
		if dbRes.Error != nil {
			if errors.Is(dbRes.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("order not found for client_order_id: %s", input.ClientOrderID)
			}
			return dbRes.Error
		}

		// Idempotency: skip if we've already recorded this exact execution (fixed: was `== ""`).
		if input.ExecutionID != "" {
			eq, eu := gplus.NewQuery[rwa.OrderExecution]()
			eq.Eq(&eu.ExecutionID, input.ExecutionID)
			existing, execRes := gplus.SelectOne(eq, gplus.Db(tx))
			if execRes.Error == nil && existing != nil {
				result = order
				return nil // already processed, nothing to do
			}
			if execRes.Error != nil && !errors.Is(execRes.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("failed to check existing execution: %w", execRes.Error)
			}
		}

		executedAt := parseTimestampOrNow(input.Timestamp)
		execution := rwa.OrderExecution{
			OrderID:     order.ID,
			ExecutionID: input.ExecutionID,
			Quantity:    execQty,
			Price:       execPrice,
			Provider:    "alpaca",
			ExecutedAt:  *executedAt,
		}
		if dbRes := gplus.Insert(&execution, gplus.Db(tx)); dbRes.Error != nil {
			return fmt.Errorf("failed to insert order execution: %w", dbRes.Error)
		}

		newFilledQty := order.FilledQuantity.Add(execQty)
		if newFilledQty.IsPositive() {
			order.FilledPrice = order.FilledPrice.Mul(order.FilledQuantity).Add(execPrice.Mul(execQty)).Div(newFilledQty)
		}
		order.FilledQuantity = newFilledQty
		order.RemainingQuantity = order.Quantity.Sub(newFilledQty)

		// Authoritative overwrite from Alpaca, if provided.
		if input.FilledAvgPrice != "" {
			if avg, err := decimal.NewFromString(input.FilledAvgPrice); err == nil {
				order.FilledPrice = avg
			}
		}
		if input.FilledQty != "" {
			if fq, err := decimal.NewFromString(input.FilledQty); err == nil {
				order.FilledQuantity = fq
				order.RemainingQuantity = order.Quantity.Sub(fq)
			}
		}

		// Status computed AFTER the authoritative overwrite (fixed: was before it).
		if order.FilledQuantity.GreaterThanOrEqual(order.Quantity) {
			order.Status = rwa.OrderStatusFilled
		} else {
			order.Status = rwa.OrderStatusPartiallyFilled
		}

		if order.ExternalOrderID == "" && input.ExternalOrderID != "" {
			order.ExternalOrderID = input.ExternalOrderID
		}

		if err := tx.Save(order).Error; err != nil {
			return fmt.Errorf("failed to update order: %w", err)
		}
		result = order
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}
	return result, nil
}

func parseTimestampOrNow(ts string) *time.Time {
	if ts == "" {
		now := time.Now()
		return &now
	}
	t, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		now := time.Now()
		return &now
	}
	return &t
}
