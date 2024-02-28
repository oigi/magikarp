package config

import (
    "github.com/redis/go-redis/v9"
    "github.com/songzhibin97/gkit/cache/local_cache"
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

var (
    REDIS  *redis.Client
    VIPER  *viper.Viper
    CONFIG Server
    LOG    *zap.Logger

    BlackCache local_cache.Cache
)
