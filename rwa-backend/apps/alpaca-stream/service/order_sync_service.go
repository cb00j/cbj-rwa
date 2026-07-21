package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	config "github.com/cb00j/cbj-rwa/rwa-backend/apps/alpaca-stream/confg"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/alpaca-stream/handlers"
	contractRwa "github.com/cb00j/cbj-rwa/rwa-backend/libs/contracts/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/kafka_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// OrderSyncService handles the business logic for trade update events from Alpaca.
type OrderSyncService struct {
	db             *gorm.DB
	evmClient      *evm_helper.EvmClient
	conf           *config.Config
	privateKey     *ecdsa.PrivateKey
	orderUpdatePub *kafka_helper.OrderUpdateKafkaService
}

// NewOrderSyncService creates a new OrderSyncService.
func NewOrderSyncService(db *gorm.DB, evmClient *evm_helper.EvmClient, conf *config.Config, privateKey *ecdsa.PrivateKey, orderUpdatePub *kafka_helper.OrderUpdateKafkaService) *OrderSyncService {
	svc := &OrderSyncService{
		db:             db,
		evmClient:      evmClient,
		conf:           conf,
		orderUpdatePub: orderUpdatePub,
	}

	// Parse backend private key if configured
	if conf.Backend != nil && conf.Backend.PrivateKey != "" {
		pk, err := crypto.HexToECDSA(conf.Backend.PrivateKey)
		if err != nil {
			log.ErrorZ(context.Background(), "NewOrderSyncService: failed to parse backend private key", zap.Error(err))
		} else {
			svc.privateKey = pk
		}

	}
	return svc
}

// HandleNew updates the order status to accepted when Alpaca acknowledges the order.
func (s *OrderSyncService) HandleNew(ctx context.Context, data handlers.TradeUpdateMessageData) {
	clientOrderID, err := extractClientOrderID(data)
	if err != nil {
		log.ErrorZ(ctx, "HandleNew: failed to extract client_order_id", zap.Error(err))
		return
	}
	order, err := s.findOrderByClientOrderID(ctx, clientOrderID)
	if err != nil {
		log.ErrorZ(ctx, "HandleNew: failed to find order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	// Idempotency: if already accepted or further along, skip
	if order.Status == rwa.OrderStatusAccepted || order.Status == rwa.OrderStatusFilled || order.Status == rwa.OrderStatusPartiallyFilled {
		log.InfoZ(ctx, "HandleNew: order already in accepted or later state, skipping",
			zap.String("client_order_id", clientOrderID),
			zap.String("current_status", string(order.Status)))
		return
	}

	order.Status = rwa.OrderStatusAccepted
	now := time.Now()
	order.AcceptedAt = &now

	// Extract external order ID from Alpaca data if available
	if data.Order.ID != "" {
		order.ExternalOrderID = data.Order.ID
	}

	if err := s.db.WithContext(ctx).Save(order).Error; err != nil {
		log.ErrorZ(ctx, "HandleNew: failed to update order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	log.InfoZ(ctx, "HandleNew: order status updated to accepted",
		zap.String("client_order_id", clientOrderID),
		zap.Uint64("order_id", order.ID),
		zap.String("external_order_id", order.ExternalOrderID))

	s.publishOrderUpdate(ctx, order, "new")
}

// HandleFill processes a full fill event from Alpaca.
func (s *OrderSyncService) HandleFill(ctx context.Context, data handlers.TradeUpdateMessageData) {
	s.handleFillOrPartialFill(ctx, data, true)
}

// HandlePartialFill processes a partial fill event from Alpaca.
func (s *OrderSyncService) HandlePartialFill(ctx context.Context, data handlers.TradeUpdateMessageData) {
	s.handleFillOrPartialFill(ctx, data, false)
}

// HandleCanceled updates the order status to cancelled.
func (s *OrderSyncService) HandleCanceled(ctx context.Context, data handlers.TradeUpdateMessageData) {
	s.handleTerminalState(ctx, data, rwa.OrderStatusCancelled, "cancelled", func(order *rwa.Order) {
		order.CancelledAt = parseTimestampOrNow(data.Timestamp)
	})
}

// HandleRejected updates the order status to rejected with a reason.
func (s *OrderSyncService) HandleRejected(ctx context.Context, data handlers.TradeUpdateMessageData) {
	clientOrderID, err := extractClientOrderID(data)
	if err != nil {
		log.ErrorZ(ctx, "HandleRejected: failed to extract client_order_id", zap.Error(err))
		return
	}

	order, err := s.findOrderByClientOrderID(ctx, clientOrderID)
	if err != nil {
		log.ErrorZ(ctx, "HandleRejected: failed to find order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	// Idempotency check
	if order.Status == rwa.OrderStatusRejected {
		log.InfoZ(ctx, "HandleRejected: order already rejected, skipping",
			zap.String("client_order_id", clientOrderID))
		return
	}

	// If order was already filled, don't try to cancel on-chain
	if order.Status == rwa.OrderStatusFilled {
		log.WarnZ(ctx, "HandleRejected: order already filled, skipping reject",
			zap.String("client_order_id", clientOrderID))
		return
	}

	order.Status = rwa.OrderStatusRejected

	// Add rejection reason to notes if available
	rejectionReason := data.Order.RejectReason
	if rejectionReason != "" {
		if order.Notes != "" {
			order.Notes += "; "
		}
		order.Notes += "Rejected by Alpaca: " + rejectionReason
	} else {
		if order.Notes != "" {
			order.Notes += "; "
		}
		order.Notes += "Rejected by Alpaca"
	}

	if err := s.db.WithContext(ctx).Save(order).Error; err != nil {
		log.ErrorZ(ctx, "HandleRejected: failed to update order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	log.WarnZ(ctx, "HandleRejected: order status updated to rejected",
		zap.String("client_order_id", clientOrderID),
		zap.Uint64("order_id", order.ID),
		zap.String("rejection_reason", rejectionReason))

	s.publishOrderUpdate(ctx, order, "rejected")

	// Call on-chain cancelOrder to refund escrowed assets to user
	go s.callCancelOrder(ctx, order)
}

// HandleDoneForDay handles the done_for_day event.
// GTC orders are not cancelled at the end of the day, but this event indicates that the market is closed for the day. This handler does not modify the order status, but can be used for logging or other purposes.
func (s *OrderSyncService) HandleDoneForDay(ctx context.Context, data handlers.TradeUpdateMessageData) {
	clientOrderID, err := extractClientOrderID(data)
	if err != nil {
		log.ErrorZ(ctx, "HandleDoneForDay: failed to extract client_order_id", zap.Error(err))
		return
	}

	order, err := s.findOrderByClientOrderID(ctx, clientOrderID)
	if err != nil {
		log.ErrorZ(ctx, "HandleDoneForDay: failed to find order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	log.InfoZ(ctx, "HandleDoneForDay: order done for day, will resume next trading day",
		zap.String("client_order_id", clientOrderID),
		zap.Uint64("order_id", order.ID),
		zap.String("current_status", string(order.Status)),
		zap.String("symbol", order.Symbol),
		zap.String("timestamp", data.Timestamp))
}

// HandleExpired updates the order status to expired.
func (s *OrderSyncService) HandleExpired(ctx context.Context, data handlers.TradeUpdateMessageData) {
	s.handleTerminalState(ctx, data, rwa.OrderStatusExpired, "expired", func(order *rwa.Order) {
		order.ExpiredAt = parseTimestampOrNow(data.Timestamp)
	})
}

func (s *OrderSyncService) handleFillOrPartialFill(ctx context.Context, data handlers.TradeUpdateMessageData, isFull bool) {
	eventLabel := "HandleFill"
	if !isFull {
		eventLabel = "HandlePartialFill"
	}

	clientOrderID, err := extractClientOrderID(data)
	if err != nil {
		log.ErrorZ(ctx, eventLabel+": failed to extract client_order_id", zap.Error(err))
		return
	}

	order, err := s.findOrderByClientOrderID(ctx, clientOrderID)
	if err != nil {
		log.ErrorZ(ctx, eventLabel+": failed to find order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	execPrice, err := decimal.NewFromString(data.Price)
	if err != nil {
		log.ErrorZ(ctx, eventLabel+": failed to parse price",
			zap.Error(err), zap.String("price", data.Price))
		return
	}

	execQty, err := decimal.NewFromString(data.Qty)
	if err != nil {
		log.ErrorZ(ctx, eventLabel+": failed to parse qty",
			zap.Error(err), zap.String("qty", data.Qty))
		return
	}

	// Begin transaction for atomicity
	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// indempotency check: if execution_id already exists, skip
		if data.ExecutionID != "" {
			q, u := gplus.NewQuery[rwa.OrderExecution]()
			q.Eq(&u.ExecutionID, data.ExecutionID)
			existing, dbRes := gplus.SelectOne(q, gplus.Db(tx))
			if dbRes.Error == nil && existing != nil {
				log.InfoZ(ctx, eventLabel+": execution_id already exists, skipping duplicate event",
					zap.String("execution_id", data.ExecutionID),
					zap.String("client_order_id", clientOrderID))
				return nil
			}

			// If dbRes.Error is not a record not found error, it indicates a database error, so we should return it
			if dbRes.Error != nil && !errors.Is(dbRes.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("failed to check existing execution: %w", dbRes.Error)
			}
		}

		// Create OrderExecution record
		execution := rwa.OrderExecution{
			OrderID:     order.ID,
			ExecutionID: data.ExecutionID,
			Quantity:    execQty,
			Price:       execPrice,
			Provider:    "alpaca",
			ExecutedAt:  *parseTimestampOrNow(data.Timestamp),
		}

		if dbRes := gplus.Insert(&execution, gplus.Db(tx)); dbRes.Error != nil {
			return fmt.Errorf("failed to insert order execution: %w", dbRes.Error)
		}

		// Update order filled quantity and price
		// New filled quantity = previous filled + this execution qty
		newFilledQty := order.FilledQuantity.Add(execQty)

		// Compute volume-weighted average price (VWAP) for filled price
		// VWAP = (old_filled_qty * old_filled_price + exec_qty * exec_price) / new_filled_qty
		if newFilledQty.IsPositive() {
			order.FilledPrice = order.FilledPrice.Mul(order.FilledQuantity).Add(execPrice.Mul(execQty)).Div(newFilledQty)
		}
		order.FilledQuantity = newFilledQty
		order.RemainingQuantity = order.Quantity.Sub(newFilledQty)

		// Use Alpaca's authoritative filled_avg_price if available
		if data.Order.FilledAvgPrice != "" {
			if avgPrice, err := decimal.NewFromString(data.Order.FilledAvgPrice); err == nil {
				order.FilledPrice = avgPrice
			}
		}

		// Use Alpaca's authoritative filled_qty for total filled if available
		if data.Order.FilledQty != "" {
			if fq, err := decimal.NewFromString(data.Order.FilledQty); err == nil {
				order.FilledQuantity = fq
				order.RemainingQuantity = order.Quantity.Sub(fq)
			}
		}

		// Determine status based on fill completeness
		if order.FilledQuantity.GreaterThanOrEqual(order.Quantity) {
			order.Status = rwa.OrderStatusFilled
		} else {
			order.Status = rwa.OrderStatusPartiallyFilled
		}

		// Extract external order ID if not already set
		if order.ExternalOrderID == "" && data.Order.ID != "" {
			order.ExternalOrderID = data.Order.ID
		}

		if err := tx.Save(order).Error; err != nil {
			return fmt.Errorf("failed to update order: %w", err)
		}

		return nil
	})

	if txErr != nil {
		log.ErrorZ(ctx, eventLabel+": transaction failed",
			zap.Error(txErr),
			zap.String("client_order_id", clientOrderID),
			zap.String("execution_id", data.ExecutionID),
			zap.String("event", data.Event),
			zap.String("price", data.Price),
			zap.String("qty", data.Qty),
			zap.String("alpaca_order_id", data.Order.ID),
			zap.String("filled_avg_price", data.Order.FilledAvgPrice),
			zap.String("filled_qty", data.Order.FilledQty),
			zap.String("timestamp", data.Timestamp),
		)

		// Persist failed event to failed_events table for later recovery
		s.persistFailedEvent(ctx, clientOrderID, data, txErr)
		return
	}

	log.InfoZ(ctx, eventLabel+": order updated successfully",
		zap.String("client_order_id", clientOrderID),
		zap.Uint64("order_id", order.ID),
		zap.String("status", string(order.Status)),
		zap.String("filled_quantity", order.FilledQuantity.String()),
		zap.String("filled_price", order.FilledPrice.String()),
		zap.String("exec_price", execPrice.String()),
		zap.String("exec_qty", execQty.String()),
		zap.String("execution_id", data.ExecutionID))

	if isFull {
		s.publishOrderUpdate(ctx, order, "fill")
	} else {
		s.publishOrderUpdate(ctx, order, "partial_fill")
	}

	// When fully filled, call on-chain markExecuted asynchronously
	if order.Status == rwa.OrderStatusFilled {
		go s.callMarkExecuted(ctx, order)
	}

}

// findOrderByClientOrderID looks up an order in the database by client_order_id.
func (s *OrderSyncService) findOrderByClientOrderID(ctx context.Context, clientOrderID string) (*rwa.Order, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if clientOrderID == "" {
		return nil, fmt.Errorf("client_order_id is empty")
	}

	q, u := gplus.NewQuery[rwa.Order]()
	q.Eq(&u.ClientOrderID, clientOrderID)
	order, dbRes := gplus.SelectOne(q, gplus.Db(s.db.WithContext(ctx)))
	if dbRes.Error != nil {
		if errors.Is(dbRes.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("order not found for client_order_id: %s", clientOrderID)
		}
		return nil, dbRes.Error
	}
	return order, nil
}

// publishOrderUpdate publishes an order status change to Kafka for real-time WS push.
func (s *OrderSyncService) publishOrderUpdate(ctx context.Context, order *rwa.Order, eventType string) {
	if s.orderUpdatePub == nil {
		return
	}
	s.orderUpdatePub.Publish(ctx, &kafka_helper.OrderUpdateEvent{
		AccountID:         order.AccountID,
		OrderID:           order.ID,
		ClientOrderID:     order.ClientOrderID,
		Symbol:            order.Symbol,
		Side:              string(order.Side),
		Status:            string(order.Status),
		FilledQuantity:    order.FilledQuantity.String(),
		FilledPrice:       order.FilledPrice.String(),
		RemainingQuantity: order.RemainingQuantity.String(),
		Quantity:          order.Quantity.String(),
		Event:             eventType,
		Timestamp:         time.Now().Unix(),
	})
}

// handleTerminalState is the shared logic for Canceled and Expired events.
func (s *OrderSyncService) handleTerminalState(
	ctx context.Context,
	data handlers.TradeUpdateMessageData,
	targetStatus rwa.OrderStatus,
	eventType string,
	setTimestamp func(order *rwa.Order),
) {
	label := fmt.Sprintf("Handle%s", targetStatus)
	clientOrderID, err := extractClientOrderID(data)
	if err != nil {
		log.ErrorZ(ctx, label+": failed to extract client_order_id", zap.Error(err))
		return
	}

	order, err := s.findOrderByClientOrderID(ctx, clientOrderID)
	if err != nil {
		log.ErrorZ(ctx, label+": failed to find order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	// Idempotency check
	if order.Status == targetStatus {
		log.InfoZ(ctx, label+": order already in target state, skipping",
			zap.String("client_order_id", clientOrderID))
		return
	}

	// If order was already filled, don't try to cancel on-chain
	if order.Status == rwa.OrderStatusFilled {
		log.WarnZ(ctx, label+": order already filled, skipping",
			zap.String("client_order_id", clientOrderID))
		return
	}

	wasPartiallyFilled := order.FilledQuantity.IsPositive()
	order.Status = targetStatus
	setTimestamp(order)

	if err := s.db.WithContext(ctx).Save(order).Error; err != nil {
		log.ErrorZ(ctx, label+": failed to update order",
			zap.Error(err), zap.String("client_order_id", clientOrderID))
		return
	}

	log.InfoZ(ctx, label+": order status updated",
		zap.String("client_order_id", clientOrderID),
		zap.Uint64("order_id", order.ID),
		zap.String("status", string(targetStatus)),
		zap.Bool("was_partially_filled", wasPartiallyFilled))

	s.publishOrderUpdate(ctx, order, eventType)

	// On-chain settlement:
	// - If partially filled: call markExecuted (settles filled portion + refunds unfilled)
	// - If not filled at all: call cancelOrder (refunds entire escrow)
	if wasPartiallyFilled {
		go s.callMarkExecuted(ctx, order)
	} else {
		go s.callCancelOrder(ctx, order)
	}
}

// callMarkExecuted sends a markExecuted transaction to the on-chain OrderContract,
// then mints the appropriate tokens to the user:
//   - Buy order: mint stock CBJToken(symbol) to user for filledQty, refund excess USDM
//   - Sell order: mint USDM to user for filledQty * filledPrice
//
// It runs asynchronously and logs errors without blocking the main flow.
func (s *OrderSyncService) callMarkExecuted(parentCtx context.Context, order *rwa.Order) {
	// Use a fresh context independent of the WebSocket connection context,
	// which may be cancelled on disconnect.
	ctx, cancel := context.WithTimeout(parentCtx, 120*time.Second)
	defer cancel()

	if traceID, ok := parentCtx.Value(log.TraceID).(string); ok {
		ctx = context.WithValue(ctx, log.TraceID, traceID)
	}

	if s.privateKey == nil {
		log.WarnZ(ctx, "callMarkExecuted: backend private key not configured, skipping on-chain call",
			zap.Uint64("order_id", order.ID))
		return
	}
	if s.conf.Chain == nil {
		log.WarnZ(ctx, "callMarkExecuted: chain config not set, skipping on-chain call",
			zap.Uint64("order_id", order.ID))
		return
	}

	chainId := s.conf.Chain.ChainId
	orderAddress := s.conf.Chain.OrderAddress

	// Parse clientOrderID as the on-chain orderId (uint256)
	onChainOrderId, err := strconv.ParseUint(order.ClientOrderID, 10, 64)
	if err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: failed to parse clientOrderID as uint",
			zap.Error(err),
			zap.String("client_order_id", order.ClientOrderID),
			zap.Uint64("order_id", order.ID))
		return
	}

	ethClient := s.evmClient.MustGetHttpClient(chainId)
	contractAddr := common.HexToAddress(orderAddress)
	orderId := new(big.Int).SetUint64(onChainOrderId)

	orderCaller, err := contractRwa.NewOrderCaller(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: failed to create OrderCaller",
			zap.Error(err), zap.Uint64("order_id", order.ID))
		return
	}

	onChainOrder, err := orderCaller.GetOrder(&bind.CallOpts{Context: ctx}, orderId)
	if err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: failed to get Order",
			zap.Error(err), zap.Uint64("order_id", order.ID))
		return
	}

	escrowAmountWei := onChainOrder.Amount // on-chain escrow in wei (18 decimals)

	// Calculate refundAmount based on order side
	// All on-chain amounts use 18 decimals
	// filledQty and filledPrice are human-readable decimals from Alpaca
	refundAmount := big.NewInt(0)
	mintAmount := big.NewInt(0)
	if order.Side == rwa.OrderSideBuy {
		// Buy: refundAmount = escrowAmount - (filledQty * filledPrice) in 18 decimals
		// actualCost = filledQty * filledPrice (USD value), convert to wei
		actualCost := order.FilledQuantity.Mul(order.FilledPrice)
		actualCostWei := decimalToWei(actualCost, 18)
		if escrowAmountWei.Cmp(actualCostWei) > 0 {
			refundAmount = new(big.Int).Sub(escrowAmountWei, actualCostWei)
		}
		mintAmount = decimalToWei(order.FilledQuantity, 18) // mint stock tokens for filledQty
	} else {
		// Sell: mintAmount = filledQty * filledPrice
		proceeds := order.FilledQuantity.Mul(order.FilledPrice)
		mintAmount = decimalToWei(proceeds, 18)
		// if order is partial filled,should refund the unsold stock tokens
		filledQtyWei := decimalToWei(order.FilledQuantity, 18)
		if escrowAmountWei.Cmp(filledQtyWei) > 0 {
			refundAmount = new(big.Int).Sub(escrowAmountWei, filledQtyWei)
		}
	}

	// Create OrderTransactor and send markExecuted
	orderTransactor, err := contractRwa.NewOrderTransactor(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: failed to create OrderTransactor",
			zap.Error(err), zap.Uint64("order_id", order.ID))
		return
	}

	auth := bind.NewKeyedTransactor(s.privateKey, new(big.Int).SetUint64(chainId))

	tx, err := orderTransactor.MarkExecuted(auth, orderId, refundAmount, mintAmount)
	if err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: markExecuted failed",
			zap.Error(err),
			zap.Uint64("order_id", order.ID),
			zap.String("client_order_id", order.ClientOrderID),
			zap.String("refund_amount", refundAmount.String()))
		return
	}
	txHash := tx.Hash().Hex()
	log.InfoZ(ctx, "callMarkExecuted: markExecuted tx sent",
		zap.Uint64("order_id", order.ID),
		zap.String("client_order_id", order.ClientOrderID),
		zap.String("tx_hash", txHash),
		zap.String("refund_amount", refundAmount.String()))

	// Save the execute tx hash and refund amount to the order record
	refundDec := weiToDecimal(refundAmount, 18)
	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", order.ID).
		Updates(map[string]interface{}{
			"execute_tx_hash": txHash,
			"refund_amount":   refundDec,
		}).Error; err != nil {
		log.ErrorZ(ctx, "callMarkExecuted: failed to save execute_tx_hash",
			zap.Error(err), zap.Uint64("order_id", order.ID))
	}
}

// callCancelOrder sends a cancelOrder transaction to the on-chain OrderContract
// to refund the user's escrowed assets (USDM for buy orders, stock tokens for sell orders).
// Called when Alpaca confirms cancellation, rejection, or expiration.
func (s *OrderSyncService) callCancelOrder(parentCtx context.Context, order *rwa.Order) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if traceID, ok := parentCtx.Value(log.TraceID).(string); ok {
		ctx = context.WithValue(ctx, log.TraceID, traceID)
	}

	if s.privateKey == nil {
		log.WarnZ(ctx, "callCancelOrder: backend private key not configured, skipping",
			zap.Uint64("order_id", order.ID))
		return
	}
	if s.conf.Chain == nil {
		log.WarnZ(ctx, "callCancelOrder: chain config not set, skipping",
			zap.Uint64("order_id", order.ID))
		return
	}

	onChainOrderId, err := strconv.ParseUint(order.ClientOrderID, 10, 64)
	if err != nil {
		log.ErrorZ(ctx, "callCancelOrder: failed to parse clientOrderID as uint",
			zap.Error(err),
			zap.String("client_order_id", order.ClientOrderID),
			zap.Uint64("order_id", order.ID))
		return
	}

	chainId := s.conf.Chain.ChainId
	ethClient := s.evmClient.MustGetHttpClient(chainId)
	contractAddr := common.HexToAddress(s.conf.Chain.OrderAddress)
	orderId := new(big.Int).SetUint64(onChainOrderId)

	orderTransactor, err := contractRwa.NewOrderTransactor(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "callCancelOrder: failed to create OrderTransactor",
			zap.Error(err), zap.Uint64("order_id", order.ID))
		return
	}

	auth := bind.NewKeyedTransactor(s.privateKey, new(big.Int).SetUint64(chainId))

	tx, err := orderTransactor.CancelOrder(auth, orderId)
	if err != nil {
		log.ErrorZ(ctx, "callCancelOrder: contract call failed",
			zap.Error(err),
			zap.Uint64("order_id", order.ID),
			zap.String("client_order_id", order.ClientOrderID),
			zap.Uint64("on_chain_order_id", onChainOrderId))
		return
	}
	txHash := tx.Hash().Hex()
	log.InfoZ(ctx, "callCancelOrder: cancel tx sent",
		zap.Uint64("order_id", order.ID),
		zap.String("client_order_id", order.ClientOrderID),
		zap.String("tx_hash", txHash))

	// Save the cancel tx hash to the order record
	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", order.ID).
		Update("cancel_tx_hash", txHash).Error; err != nil {
		log.ErrorZ(ctx, "callCancelOrder: failed to save cancel_tx_hash",
			zap.Error(err),
			zap.Uint64("order_id", order.ID),
			zap.String("tx_hash", txHash))
	}
}

// extractClientOrderID extracts the client_order_id from the Alpaca trade update data.
func extractClientOrderID(data handlers.TradeUpdateMessageData) (string, error) {
	if data.Order.ClientOrderID == "" {
		return "", fmt.Errorf("client_order_id not found or empty in order data")
	}
	return data.Order.ClientOrderID, nil
}

// decimalToWei converts a decimal.Decimal to *big.Int with the given number of decimals.
// e.g., decimalToWei(1.5, 18) = 1500000000000000000
func decimalToWei(d decimal.Decimal, decimals int) *big.Int {
	// Multiply by 10^decimals, then convert to big.Int
	factor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	wei := d.Mul(factor)
	result := new(big.Int)
	result.SetString(wei.StringFixed(0), 10)
	return result
}

// weiToDecimal converts a *big.Int wei value to decimal.Decimal with the given decimals.
func weiToDecimal(wei *big.Int, decimals int) decimal.Decimal {
	factor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	return decimal.NewFromBigInt(wei, 0).Div(factor)
}

// parseTimestampOrNow attempts to parse an RFC3339 timestamp string.
// Falls back to time.Now() if parsing fails or the string is empty.
// Returns a pointer to time.Time for use with nullable time fields.
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

// persistFailedEvent saves a failed WebSocket trade update event to the failed_events table
// so it can be recovered or retried later.
func (s *OrderSyncService) persistFailedEvent(ctx context.Context, clientOrderID string, data handlers.TradeUpdateMessageData, originalErr error) {
	eventDataJSON, err := json.Marshal(data)
	if err != nil {
		log.ErrorZ(ctx, "persistFailedEvent: failed to marshal event data",
			zap.Error(err),
			zap.String("client_order_id", clientOrderID))
		return
	}

	failedEvent := rwa.FailedEvent{
		ClientOrderID: clientOrderID,
		EventType:     data.Event,
		ExecutionID:   data.ExecutionID,
		EventData:     string(eventDataJSON),
		ErrorMessage:  originalErr.Error(),
		Source:        "alpaca",
		Status:        "pending",
	}

	if dbRes := gplus.Insert(&failedEvent, gplus.Db(s.db.WithContext(ctx))); dbRes.Error != nil {
		log.ErrorZ(ctx, "persistFailedEvent: failed to insert failed event record",
			zap.Error(dbRes.Error),
			zap.String("client_order_id", clientOrderID),
			zap.String("event_type", data.Event))
		return
	}

	log.InfoZ(ctx, "persistFailedEvent: failed event persisted for later recovery",
		zap.Uint64("failed_event_id", failedEvent.ID),
		zap.String("client_order_id", clientOrderID),
		zap.String("event_type", data.Event),
		zap.String("execution_id", data.ExecutionID))
}
