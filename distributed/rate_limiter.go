package distributed

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	TxPipeline() redis.Pipeliner
}

type RateLimiter struct {
	client   Redis
	limit    int
	duration time.Duration
}

func NewRateLimiter(client Redis, limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		client:   client,
		limit:    limit,
		duration: duration,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	pipe := rl.client.TxPipeline()

	counter := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, rl.duration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to exec pipe: %w", err)
	}

	if counter.Val() > int64(rl.limit) {
		return false, nil
	}

	return true, nil
}
