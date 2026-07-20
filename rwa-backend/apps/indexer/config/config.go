package config

import (
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/bootstrap"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/trade"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/database"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
)

type Config struct {
	AppName string                `json:"appName" yaml:"appName"`
	Chain   *ChainInfo            `json:"chain" yaml:"chain"`
	RpcInfo evm_helper.RpcInfoMap `json:"rpcInfo" yaml:"rpcInfo"`
	Logger  *log.Conf             `json:"logger" yaml:"logger"`
	Db      *database.DbConf      `json:"db" yaml:"db"`
	Alpaca  *trade.AlpacaConfig   `json:"alpaca" yaml:"alpaca"`
	Indexer *IndexerConfig        `json:"indexer" yaml:"indexer"`
	Backend *BackendConfig        `json:"backend" yaml:"backend"`
}

// BackendConfig contains the backend wallet private key for signing transactions
type BackendConfig struct {
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}

type ChainInfo struct {
	ChainId        uint64 `json:"chainId" yaml:"chainId"`
	OrderAddress   string `json:"orderAddress" yaml:"orderAddress"`
	GatewayAddress string `json:"gatewayAddress" yaml:"gatewayAddress"`
	UsdmAddress    string `json:"usdmAddress" yaml:"usdmAddress"`
}

type IndexerConfig struct {
	PollInterval       int    `json:"pollInterval" yaml:"pollInterval"`
	BatchSize          int    `json:"batchSize" yaml:"batchSize"`
	StartBlock         uint64 `json:"startBlock" yaml:"startBlock"`
	ConfirmationBlocks uint64 `json:"confirmationBlocks" yaml:"confirmationBlocks"`
}

func NewConfig(configFile string) (*Config, error) {
	return bootstrap.LoadConfig[Config](configFile)
}
