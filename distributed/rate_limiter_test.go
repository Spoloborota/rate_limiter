package distributed_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"

	ratelimiter "github.com/Spoloborota/rate_limiter/distributed"
)

func TestRateLimiter(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	limiter := ratelimiter.NewRateLimiter(client, 5, time.Second*10)

	ctx := context.Background()
	key := "test_key"
	client.Del(ctx, key)

	for i := 0; i < 5; i++ {
		allowed, err := limiter.Allow(ctx, key)
		require.NoError(t, err)
		require.True(t, allowed)
	}

	allowed, err := limiter.Allow(ctx, key)
	require.NoError(t, err)
	require.False(t, allowed)
}
