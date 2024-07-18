package redis

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type ClientRedis struct {
	Client *redis.Client
}

func InitRedis() map[string]*ClientRedis {
	redisCaches := make(map[string]*ClientRedis)
	redisConfig := config.CONFIG.Redis
	for k, v := range redisConfig.Dbs {
		client := redis.NewClient(&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password,
			DB:       v,
			PoolSize: 20,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := client.Ping(ctx).Result()
		if err != nil {
			config.LOG.Error("", zap.Error(err))
			return nil
		}

		redisCaches[k] = &ClientRedis{
			Client: client,
		}
	}

	return redisCaches
}

//func InitRedis() map[string]*ClientRedis {
//	redisConfig := config.CONFIG.Redis
//	redisClients := make(map[string]*ClientRedis)
//	var mu sync.Mutex // 用于保护 redisClients 的并发访问
//
//	var wg sync.WaitGroup
//	var initError error // 用于收集初始化过程中的错误
//
//	for dbName, dbIndex := range redisConfig.Dbs {
//		wg.Add(1)
//		go func(dbName string, dbIndex int) {
//			defer wg.Done()
//
//			client := redis.NewClient(&redis.Options{
//				Addr:     redisConfig.Addr,
//				Password: redisConfig.Password,
//				DB:       dbIndex,
//			})
//
//			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//			defer cancel()
//
//			_, err := client.Ping(ctx).Result()
//			if err != nil {
//				mu.Lock()
//				initError = fmt.Errorf("failed to ping DB %s: %w", dbName, err)
//				mu.Unlock()
//				return
//			}
//
//			mu.Lock()
//			redisClients[dbName] = &ClientRedis{Client: client}
//			mu.Unlock()
//		}(dbName, dbIndex)
//	}
//
//	// 如果在初始化过程中有错误发生，则返回错误
//	if initError != nil {
//		config.LOG.Error("", zap.Error(initError))
//		return nil
//	}
//
//	return redisClients
//}
