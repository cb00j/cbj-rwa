package service

import (
	"context"
	"crypto/ecdsa"

	"github.com/cb00j/cbj-rwa/rwa-backend/apps/indexer/config"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/indexer/service/handlers"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/trade"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func LoadModule() fx.Option {
	return fx.Module("service",
		fx.Provide(
			NewBlockService,
			NewEventListener,
			NewProcessTxService,
			provideHandlers,
		),
		fx.Invoke(registerEventListener),
	)
}

func provideHandlers(
	tradeService trade.TradeService,
	conf *config.Config,
	evmClient *evm_helper.EvmClient,
	db *gorm.DB, // 新增:AfterCommit要在事务外单独查/改库,需要主db连接,不能用tx
) ([]EventHandler, error) {
	var backendPK *ecdsa.PrivateKey
	if conf.Backend != nil && conf.Backend.PrivateKey != "" {
		pk, err := crypto.HexToECDSA(conf.Backend.PrivateKey)
		if err != nil {
			log.ErrorZ(context.Background(), "failed to parse backend private key", zap.Error(err))
		} else {
			backendPK = pk
		}
	}

	orderSubmittedHandler, err := handlers.NewHandleOrderSubmitted(
		tradeService, conf.Chain.UsdmAddress, evmClient, conf.Chain.ChainId, conf.Chain.OrderAddress, backendPK, db)
	if err != nil {
		return nil, err
	}

	orderCancelledHandler, err := handlers.NewHandleOrderCancelled(tradeService)
	if err != nil {
		return nil, err
	}

	cancelRequestedHandler, err := handlers.NewHandleCancelRequested(tradeService)
	if err != nil {
		return nil, err
	}

	orderExecutedHandler, err := handlers.NewHandleOrderExecuted(tradeService)
	if err != nil {
		return nil, err
	}

	orderBackendRefundedHandler, err := handlers.NewHandleOrderBackendRefunded()
	if err != nil {
		return nil, err
	}

	// TODO: Add Gate contract event handlers once Gate Go bindings are generated.

	return []EventHandler{
		orderSubmittedHandler,
		orderCancelledHandler,
		cancelRequestedHandler,
		orderExecutedHandler,
		orderBackendRefundedHandler,
	}, nil
}

func registerEventListener(
	lc fx.Lifecycle,
	eventListener *EventListener,
	blockService *BlockService,
	processService *ProcessTxService,
	conf *config.Config,
) {
	pollingCtx, pollingCancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := processService.InitCache(ctx); err != nil {
				log.ErrorZ(ctx, "failed to init process service cache", zap.Error(err))
				pollingCancel()
				return err
			}

			go func() {
				if err := eventListener.StartPolling(pollingCtx, blockService, processService); err != nil {
					log.ErrorZ(pollingCtx, "event listener polling stopped with error", zap.Error(err))
				}
			}()

			log.InfoZ(ctx, "event listener started",
				zap.Uint64("chainId", conf.Chain.ChainId), zap.String("orderAddress", conf.Chain.OrderAddress))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.InfoZ(ctx, "stopping event listener")
			pollingCancel()
			return nil
		},
	})
}
