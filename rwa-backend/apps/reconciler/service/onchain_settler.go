package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"time"

	config "github.com/cb00j/cbj-rwa/rwa-backend/apps/reconciler/config"
	contractRwa "github.com/cb00j/cbj-rwa/rwa-backend/libs/contracts/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// OnchainSettler holds the one thing both alpaca-stream (live path) and
// reconciler (retry path) need in common: how to actually call
// markExecuted/cancelOrder on OrderContract. Keeping this in one place means
// a bugfix here (like the two below) only has to happen once.
type OnchainSettler struct {
	db         *gorm.DB
	evmClient  *evm_helper.EvmClient
	conf       *config.Config
	privateKey *ecdsa.PrivateKey
}

func NewOnchainSettler(db *gorm.DB, evmClient *evm_helper.EvmClient, conf *config.Config) (*OnchainSettler, error) {
	s := &OnchainSettler{db: db, evmClient: evmClient, conf: conf}

	if conf.Backend == nil || conf.Backend.PrivateKey == "" {
		log.WarnZ(context.Background(), "OnchainSettler: backend private key not configured; on-chain calls will be skipped")
		return s, nil
	}
	pkHex := conf.Backend.PrivateKey
	if len(pkHex) > 1 && pkHex[:2] == "0x" {
		pkHex = pkHex[2:]
	}
	pk, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse backend private key: %w", err)
	}
	s.privateKey = pk
	return s, nil
}

// CallMarkExecuted settles a filled (or partially-filled-then-terminated)
// order on-chain: refunds unfilled escrow and mints the appropriate asset.
//
// NOTE on fixes made while extracting this from alpaca-stream's
// order_sync_service.go:
//  1. Sell orders previously never computed a mintAmount at all — the buyer
//     side branch existed but sell had no equivalent, so a filled sell order
//     minted zero USDM to the seller. Fixed below.
//  2. The buy-side refund comparison used `> 0` where it needed `< 0`
//     (refund is owed when actual cost is LESS than the escrowed worst-case
//     amount, not more). Fixed below.
func (s *OnchainSettler) CallMarkExecuted(parentCtx context.Context, order *rwa.Order) {
	ctx, cancel := context.WithTimeout(parentCtx, 120*time.Second)
	defer cancel()
	if traceID, ok := parentCtx.Value(log.TraceID).(string); ok {
		ctx = context.WithValue(ctx, log.TraceID, traceID)
	}

	if s.privateKey == nil {
		log.WarnZ(ctx, "CallMarkExecuted: backend private key not configured, skipping", zap.Uint64("order_id", order.ID))
		return
	}
	if s.conf.Chain == nil {
		log.WarnZ(ctx, "CallMarkExecuted: chain config not set, skipping", zap.Uint64("order_id", order.ID))
		return
	}

	chainId := s.conf.Chain.ChainId
	onChainOrderId, err := strconv.ParseUint(order.ClientOrderID, 10, 64)
	if err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: failed to parse clientOrderID as uint",
			zap.Error(err), zap.String("client_order_id", order.ClientOrderID), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	ethClient := s.evmClient.MustGetHttpClient(chainId)
	contractAddr := common.HexToAddress(s.conf.Chain.OrderAddress)
	orderId := new(big.Int).SetUint64(onChainOrderId)

	orderCaller, err := contractRwa.NewOrderCaller(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: failed to create OrderCaller", zap.Error(err), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	onChainOrder, err := orderCaller.GetOrder(&bind.CallOpts{Context: ctx}, orderId)
	if err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: failed to get on-chain order", zap.Error(err), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}
	escrowAmountWei := onChainOrder.Amount

	refundAmount := big.NewInt(0)
	mintAmount := big.NewInt(0)

	if order.Side == rwa.OrderSideBuy {
		// Buy: escrow was price*qty at the worst acceptable price. Refund the
		// difference if the real fill cost less than that.
		actualCost := order.FilledQuantity.Mul(order.FilledPrice)
		actualCostWei := decimalToWei(actualCost, 18)
		if actualCostWei.Cmp(escrowAmountWei) < 0 { // fixed: was `> 0`
			refundAmount = new(big.Int).Sub(escrowAmountWei, actualCostWei)
		}
		mintAmount = decimalToWei(order.FilledQuantity, 18) // mint stock tokens for filledQty
	} else {
		// Sell: mint USDM for filledQty * filledPrice (fixed: previously missing entirely).
		proceeds := order.FilledQuantity.Mul(order.FilledPrice)
		mintAmount = decimalToWei(proceeds, 18)
		// If only partially filled before cancel/expire, refund the unsold
		// stock-token share of the escrow.
		filledQtyWei := decimalToWei(order.FilledQuantity, 18)
		if escrowAmountWei.Cmp(filledQtyWei) > 0 {
			refundAmount = new(big.Int).Sub(escrowAmountWei, filledQtyWei)
		}
	}

	orderTransactor, err := contractRwa.NewOrderTransactor(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: failed to create OrderTransactor", zap.Error(err), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	auth := bind.NewKeyedTransactor(s.privateKey, new(big.Int).SetUint64(chainId))

	tx, err := orderTransactor.MarkExecuted(auth, orderId, refundAmount, mintAmount)
	if err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: markExecuted failed",
			zap.Error(err), zap.Uint64("order_id", order.ID), zap.String("client_order_id", order.ClientOrderID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	txHash := tx.Hash().Hex()
	log.InfoZ(ctx, "CallMarkExecuted: markExecuted tx sent",
		zap.Uint64("order_id", order.ID), zap.String("tx_hash", txHash))

	refundDec := weiToDecimal(refundAmount, 18)
	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", order.ID).
		Updates(map[string]interface{}{
			"execute_tx_hash": txHash,
			"refund_amount":   refundDec,
		}).Error; err != nil {
		log.ErrorZ(ctx, "CallMarkExecuted: failed to save execute_tx_hash", zap.Error(err), zap.Uint64("order_id", order.ID))
	}
	s.recordSuccess(ctx, order.ID)
}

// CallCancelOrder refunds the full escrow for an order that never filled at all.
func (s *OnchainSettler) CallCancelOrder(parentCtx context.Context, order *rwa.Order) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if traceID, ok := parentCtx.Value(log.TraceID).(string); ok {
		ctx = context.WithValue(ctx, log.TraceID, traceID)
	}

	if s.privateKey == nil {
		log.WarnZ(ctx, "CallCancelOrder: backend private key not configured, skipping", zap.Uint64("order_id", order.ID))
		return
	}
	if s.conf.Chain == nil {
		log.WarnZ(ctx, "CallCancelOrder: chain config not set, skipping", zap.Uint64("order_id", order.ID))
		return
	}

	onChainOrderId, err := strconv.ParseUint(order.ClientOrderID, 10, 64)
	if err != nil {
		log.ErrorZ(ctx, "CallCancelOrder: failed to parse clientOrderID as uint",
			zap.Error(err), zap.String("client_order_id", order.ClientOrderID), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	chainId := s.conf.Chain.ChainId
	ethClient := s.evmClient.MustGetHttpClient(chainId)
	contractAddr := common.HexToAddress(s.conf.Chain.OrderAddress)
	orderId := new(big.Int).SetUint64(onChainOrderId)

	orderTransactor, err := contractRwa.NewOrderTransactor(contractAddr, ethClient)
	if err != nil {
		log.ErrorZ(ctx, "CallCancelOrder: failed to create OrderTransactor", zap.Error(err), zap.Uint64("order_id", order.ID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	auth := bind.NewKeyedTransactor(s.privateKey, new(big.Int).SetUint64(chainId))

	tx, err := orderTransactor.CancelOrder(auth, orderId)
	if err != nil {
		log.ErrorZ(ctx, "CallCancelOrder: contract call failed",
			zap.Error(err), zap.Uint64("order_id", order.ID), zap.String("client_order_id", order.ClientOrderID))
		s.recordFailure(ctx, order.ID, err.Error())
		return
	}

	txHash := tx.Hash().Hex()
	log.InfoZ(ctx, "CallCancelOrder: cancel tx sent", zap.Uint64("order_id", order.ID), zap.String("tx_hash", txHash))

	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", order.ID).
		Update("cancel_tx_hash", txHash).Error; err != nil {
		log.ErrorZ(ctx, "CallCancelOrder: failed to save cancel_tx_hash", zap.Error(err), zap.Uint64("order_id", order.ID))
	}
	s.recordSuccess(ctx, order.ID)
}

// recordFailure / recordSuccess maintain the retry-tracking fields added in
// the settlement_attempts migration, so ReconciliationService knows what
// still needs retrying and how many times it's already tried.
func (s *OnchainSettler) recordFailure(ctx context.Context, orderID uint64, errMsg string) {
	now := time.Now()
	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"settlement_attempts":        gorm.Expr("settlement_attempts + 1"),
			"last_settlement_error":      errMsg,
			"last_settlement_attempt_at": now,
		}).Error; err != nil {
		log.ErrorZ(ctx, "OnchainSettler: failed to record settlement failure", zap.Error(err), zap.Uint64("order_id", orderID))
	}
}

func (s *OnchainSettler) recordSuccess(ctx context.Context, orderID uint64) {
	if err := s.db.WithContext(ctx).Model(&rwa.Order{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"settlement_attempts":   0,
			"last_settlement_error": "",
		}).Error; err != nil {
		log.ErrorZ(ctx, "OnchainSettler: failed to reset settlement tracking", zap.Error(err), zap.Uint64("order_id", orderID))
	}
}

func decimalToWei(d decimal.Decimal, decimals int) *big.Int {
	factor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	wei := d.Mul(factor)
	result := new(big.Int)
	result.SetString(wei.StringFixed(0), 10)
	return result
}

func weiToDecimal(wei *big.Int, decimals int) decimal.Decimal {
	factor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(decimals)))
	return decimal.NewFromBigInt(wei, 0).Div(factor)
}
