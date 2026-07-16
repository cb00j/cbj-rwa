package service

import (
	"context"
	"crypto/ecdsa"
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

func (s *OrderSyncService) handleFillOrPartialFill(ctx context.Context, data handlers.TradeUpdateMessageData, isFull bool) {
	// TODO
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
		log.ErrorZ(ctx, "callMarkExecuted: failed to create OrderCaller",
			zap.Error(err), zap.Uint64("order_id", order.ID))
		return
	}

	userAddr := onChainOrder.User
	escrowAmountWei := onChainOrder.Amount // on-chain escrow in wei (18 decimals)

	// Calculate refundAmount based on order side
	// All on-chain amounts use 18 decimals
	// filledQty and filledPrice are human-readable decimals from Alpaca
	refundAmount := big.NewInt(0)
	if order.Side == rwa.OrderSideBuy {
		// Buy: refundAmount = escrowAmount - (filledQty * filledPrice) in 18 decimals
		// actualCost = filledQty * filledPrice (USD value), convert to wei
		actualCost := order.FilledQuantity.Mul(order.FilledPrice)
		actualCostWei := decimalToWei(actualCost, 18)
		if actualCostWei.Cmp(escrowAmountWei) > 0 {
			refundAmount = new(big.Int).Sub(escrowAmountWei, actualCostWei)
		}
	}
	// Sell: no USDM refund needed (escrowed stock tokens are consumed)

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
