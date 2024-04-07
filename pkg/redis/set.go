package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (c *ClientRedis) Set(ctx context.Context, key string, value string, expire time.Duration) error {
	return c.Client.Set(ctx, key, value, expire).Err()
}

func (c *ClientRedis) SetObj(ctx context.Context, key string, value any, expire time.Duration) error {
	valueJson, err := json.Marshal(value)

	if err != nil {
		return err
	}

	if err := c.Set(ctx, key, string(valueJson), expire); err != nil {
		return err
	}

	return nil
}