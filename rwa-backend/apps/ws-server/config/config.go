package config

import (
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/bootstrap"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/kafka_helper"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/redis_cache"
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
)

type Config struct {
	AppName string                    `json:"appName" yaml:"appName"`
	Server  *ServerConfig             `json:"server" yaml:"server"`
	Redis   *redis_cache.RedisConfig  `json:"redis" yaml:"redis"`
	Kafka   *kafka_helper.KafkaConfig `json:"kafka" yaml:"kafka"`
	Logger  *log.Conf                 `json:"logger" yaml:"logger"`
}

type ServerConfig struct {
	Port     int    `json:"port" yaml:"port"`
	BasePath string `json:"basePath" yaml:"basePath"`
}

func NewConfig(configFile string) (*Config, error) {
	return bootstrap.LoadConfig[Config](configFile)
}
