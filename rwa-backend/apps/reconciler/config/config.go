package config

import (
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/bootstrap"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/evm_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/database"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
)

// Config is intentionally narrow — reconciler only needs enough to talk to
// the database and to sign/send on-chain transactions. It does not touch
// Alpaca or Kafka directly.
type Config struct {
	AppName string                    `json:"appName" yaml:"appName"`
	Db      *database.DbConf          `json:"db" yaml:"db"`
	RpcInfo evm_helper.RpcInfoMap     `json:"rpcInfo" yaml:"rpcInfo"`
	Chain   *ChainConfig              `json:"chain" yaml:"chain"`
	Backend *BackendConfig            `json:"backend" yaml:"backend"`
	Logger  *log.Conf                 `json:"logger" yaml:"logger"`
}

type ChainConfig struct {
	ChainId      uint64 `json:"chainId" yaml:"chainId"`
	OrderAddress string `json:"orderAddress" yaml:"orderAddress"`
}

type BackendConfig struct {
	// Must be the address that actually holds BACKEND_ROLE on OrderContract —
	// see the earlier AccessControlUnauthorizedAccount incident.
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}

func NewConfig(configFile string) (*Config, error) {
	return bootstrap.LoadConfig[Config](configFile)
}
