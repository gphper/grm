/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 13:23:38
 */
package common

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type RedisLog struct {
	Logger *zap.Logger
}

func (rl RedisLog) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (rl RedisLog) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	value, ok := ctx.Value("username").(string)
	if ok {
		byteLog, _ := json.Marshal(cmd.Args())
		rl.Logger.Info(string(byteLog), zap.String("username", value))
	}
	return nil
}

func (rl RedisLog) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return context.Background(), nil
}

func (rl RedisLog) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
