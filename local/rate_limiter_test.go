package local_test

import (
	"testing"
	"time"

	ratelimiter "github.com/Spoloborota/rate_limiter/local"
	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	tests := []struct {
		name           string
		rate           int
		interval       time.Duration
		expectedAllows []bool
	}{
		{
			name:           "5 операций в секунду",
			rate:           5,
			interval:       time.Second,
			expectedAllows: []bool{true, true, true, true, true, false, false},
		},
		{
			name:           "2 операции за полсекунды",
			rate:           2,
			interval:       500 * time.Millisecond,
			expectedAllows: []bool{true, true, false},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rl := ratelimiter.NewRateLimiter(test.rate, test.interval)

			for i := 0; i < len(test.expectedAllows); i++ {
				require.Equal(t, test.expectedAllows[i], rl.Allow())
			}
		})
	}
}

func TestRateLimiter_ResetsAfterInterval(t *testing.T) {
	rl := ratelimiter.NewRateLimiter(2, time.Second)

	require.True(t, rl.Allow())
	require.True(t, rl.Allow())
	require.False(t, rl.Allow())

	time.Sleep(time.Second)

	require.True(t, rl.Allow())
	require.True(t, rl.Allow())
	require.False(t, rl.Allow())
}
