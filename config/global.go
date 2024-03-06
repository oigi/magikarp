package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	REDIS  *redis.Client
	VIPER  *viper.Viper
	CONFIG *Config
	LOG    *zap.Logger
)
