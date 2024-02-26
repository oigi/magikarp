package initialize

import (
	"context"
	"github.com/oigi/Magikarp/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
