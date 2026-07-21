package service

import (
	config "github.com/cb00j/cbj-rwa/rwa-backend/apps/alpaca-stream/confg"
	"github.com/cb00j/cbj-rwa/rwa-backend/apps/alpaca-stream/ws"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/kafka_helper"
)

type AlpacaWebScoketService struct {
	config             *config.Config
	orderSyncService   *OrderSyncService
	barkafakaService   *kafka_helper.BarKafkaService
	tradeUpdatesClient *ws.Client
	marketDataClient   *ws.Client
	subscriptionMgr    *ws.SubscriptionManager
	marketDataSubMgr   *ws.SubscriptionManager
}
