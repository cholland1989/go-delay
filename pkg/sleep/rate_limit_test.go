package sleep

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func HealthCheck() error {
	return nil
}

func ExampleRateLimit() {
	for attempt := 0; attempt < 5; attempt++ {
		err := HealthCheck()
		if err == nil {
			break
		}
		RateLimit(10, time.Second, 0.5)
	}
	// Output:
}

func ExampleRateLimitWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for attempt := 0; attempt < 5; attempt++ {
		err := HealthCheck()
		if err == nil {
			break
		}
		err = RateLimitWithContext(ctx, 10, time.Second, 0.5)
		if err != nil {
			break
		}
	}
	// Output:
}

func TestRateLimit(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		actions  int
		period   time.Duration
		expected time.Duration
	}{
		"Rate Limit (1)":  {1, time.Millisecond, time.Millisecond},
		"Rate Limit (2)":  {2, time.Millisecond, time.Millisecond / 2},
		"Rate Limit (4)":  {4, time.Millisecond, time.Millisecond / 4},
		"Rate Limit (8)":  {8, time.Millisecond, time.Millisecond / 8},
		"Rate Limit (10)": {10, time.Millisecond, time.Millisecond / 10},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			timestamp := time.Now()
			RateLimit(params.actions, params.period, 0.0)
			delay := time.Since(timestamp)
			require.LessOrEqual(test, params.expected, delay)
		})
	}
}

func TestRateLimitWithContext(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		actions  int
		period   time.Duration
		expected time.Duration
	}{
		"Rate Limit (1)":  {1, time.Millisecond, time.Millisecond},
		"Rate Limit (2)":  {2, time.Millisecond, time.Millisecond / 2},
		"Rate Limit (4)":  {4, time.Millisecond, time.Millisecond / 4},
		"Rate Limit (8)":  {8, time.Millisecond, time.Millisecond / 8},
		"Rate Limit (10)": {10, time.Millisecond, time.Millisecond / 10},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			for _, ctx := range []context.Context{nil, context.Background()} {
				timestamp := time.Now()
				err := RateLimitWithContext(ctx, params.actions, params.period, 0.0)
				delay := time.Since(timestamp)
				require.NoError(test, err)
				require.LessOrEqual(test, params.expected, delay)
			}
		})
	}
}

func TestRateLimitWithContext_WithCancel(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timestamp := time.Now()
	err := RateLimitWithContext(ctx, 1, time.Second, 0.0)
	delay := time.Since(timestamp)
	require.ErrorIs(test, err, context.Canceled)
	require.GreaterOrEqual(test, time.Millisecond, delay)
}
