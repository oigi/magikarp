package global

import (
	"github.com/oigi/Magikarp/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	VIPER  *viper.Viper
	CONFIG config.Server
	LOG    *zap.Logger

	BlackCache local_cache.Cache
)
