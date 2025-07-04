package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type LoggingHook struct {
	log *zap.Logger
}

func (h *LoggingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (h *LoggingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	if err := cmd.Err(); err != nil && err != redis.Nil {
		h.log.Error("redis error", zap.String("cmd", cmd.Name()), zap.Error(err))
	}
	return nil
}
