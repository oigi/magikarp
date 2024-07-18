package redis

import (
	"context"
)

func (c *ClientRedis) ListSize(ctx context.Context, key string) (int64, error) {
	l, err := c.Client.LLen(ctx, key).Result()

	if err != nil {
		return -1, err
	}

	return l, nil
}

func (c *ClientRedis) LPop(ctx context.Context, key string) (string, error) {
	return c.Client.LPop(ctx, key).Result()
}

func (c *ClientRedis) LPops(ctx context.Context, key string, times int) ([]string, error) {
	// 使用 Lua 脚本实现一次调用完成多次弹出操作
	script := `
		local result = {}
		for i = 1, ARGV[1] do
			local val = redis.call("LPOP", KEYS[1])
			if val then
				table.insert(result, val)
			else
				break
			end
		end
		return result
	`
	// 执行 Lua 脚本
	cmd := c.Client.Eval(ctx, script, []string{key}, times)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	// 解析结果
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	// 将结果转换为字符串切片
	var values []string
	for _, v := range result.([]interface{}) {
		values = append(values, v.(string))
	}

	return values, nil
}
