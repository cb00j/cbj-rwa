// apps/indexer/service/handlers/handle_order_backend_refunded.go
package handlers

import (
	"context"
	"errors"

	"github.com/acmestack/gorm-plus/gplus"
	contractRwa "github.com/cb00j/cbj-rwa/rwa-backend/libs/contracts/rwa"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/models/rwa"
	coreTypes "github.com/cb00j/cbj-rwa/rwa-backend/libs/core/types"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HandleOrderBackendRefunded struct {
	topic0        string
	orderFilterer *contractRwa.OrderFilterer
}

func NewHandleOrderBackendRefunded() (*HandleOrderBackendRefunded, error) {
	orderFilterer, err := contractRwa.NewOrderFilterer(common.Address{}, nil)
	if err != nil {
		return nil, err
	}

	orderAbi, err := contractRwa.OrderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	orderBackendRefundedEvent, ok := orderAbi.Events["OrderBackendRefunded"]
	if !ok {
		return nil, errors.New("orderBackendRefunded event not found in Order ABI")
	}

	return &HandleOrderBackendRefunded{
		topic0:        orderBackendRefundedEvent.ID.Hex(),
		orderFilterer: orderFilterer,
	}, nil
}

func (h *HandleOrderBackendRefunded) ContractType() coreTypes.ContractType {
	return coreTypes.ContractTypeOrder
}

func (h *HandleOrderBackendRefunded) Topic0() string {
	return h.topic0
}

func (h *HandleOrderBackendRefunded) HandleEvent(ctx context.Context, tx *gorm.DB, event *coreTypes.EventLogWithId) error {
	ethLog, err := event.ConvertToEthLog()
	if err != nil {
		log.ErrorZ(ctx, "failed to convert event to eth log", zap.Error(err))
		return err
	}

	parsedEvent, err := h.orderFilterer.ParseOrderBackendRefunded(*ethLog)
	if err != nil {
		log.ErrorZ(ctx, "failed to parse Order OrderBackendRefunded event", zap.Error(err))
		return err
	}

	orderIdStr := parsedEvent.OrderId.String()
	q, u := gplus.NewQuery[rwa.Order]()
	q.Eq(&u.ClientOrderID, orderIdStr)
	order, dbRes := gplus.SelectOne(q, gplus.Db(tx.WithContext(ctx)))
	if dbRes.Error != nil {
		log.ErrorZ(ctx, "HandleOrderBackendRefunded: order not found",
			zap.Error(dbRes.Error), zap.String("orderId", orderIdStr))
		return dbRes.Error
	}

	if err := tx.Model(order).Update("backend_refund_tx_hash", event.TxHash).Error; err != nil {
		log.ErrorZ(ctx, "HandleOrderBackendRefunded: failed to save backend_refund_tx_hash",
			zap.Error(err), zap.Uint64("dbOrderId", order.ID))
		return err
	}

	log.InfoZ(ctx, "HandleOrderBackendRefunded: refund confirmed on-chain, status left untouched",
		zap.String("orderId", orderIdStr),
		zap.Uint64("dbOrderId", order.ID),
		zap.String("status", string(order.Status)),
		zap.String("txHash", event.TxHash))

	return nil
}
