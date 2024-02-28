package initialize

import (
    "context"
    "github.com/oigi/Magikarp/config"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)

func Redis() {
    redisConfig := config.CONFIG.Redis
    client := redis.NewClient(&redis.Options{
        Addr:     redisConfig.Addr,
        Password: redisConfig.Password,
        DB:       redisConfig.DB,
    })

    pong, err := client.Ping(context.Background()).Result()
    if err != nil {
        config.LOG.Error("redis connect ping failed, e:", zap.Error(err))
    } else {
        config.LOG.Info("redis connect ping response:", zap.String("pong", pong))
        config.REDIS = client
    }
}
