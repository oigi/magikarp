package redis

import "context"

func (c *ClientRedis) LPush(ctx context.Context, key string, value ...string) error {
	return c.Client.LPush(ctx, key, value).Err()
}

func (c *ClientRedis) RPush(ctx context.Context, key string, value ...string) error {
	return c.Client.RPush(ctx, key, value).Err()
}
