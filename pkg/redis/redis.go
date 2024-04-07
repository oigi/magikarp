package redis

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type ClientRedis struct {
	Client *redis.Client
}

func InitRedis() (map[string]*redis.Client, error) {
	redisConfig := config.CONFIG.Redis
	redisClients := make(map[string]*redis.Client)
	errs := make(chan error)

	// 初始化多个Redis客户端
	for k, v := range redisConfig.Dbs {
		client := redis.NewClient(&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password,
			DB:       v,
		})

		// 并发进行Ping操作
		go func(client *redis.Client, dbName string) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := client.Ping(ctx).Result()
			if err != nil {
				errs <- err
				return
			}

			redisClients[dbName] = client
		}(client, k)
	}

	// 等待所有Ping操作完成或出现错误
	for range redisConfig.Dbs {
		if err := <-errs; err != nil {
			return redisClients, err
		}
	}

	return redisClients, nil
}
