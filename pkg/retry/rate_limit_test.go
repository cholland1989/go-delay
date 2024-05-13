package retry

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func HealthCheck() error {
	return io.EOF
}

func ExampleRateLimit() {
	_ = RateLimit(10, time.Second, 0.5, HealthCheck)
	// Output:
}

func ExampleRateLimitWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = RateLimitWithContext(ctx, 10, time.Second, 0.5, HealthCheck)
	// Output:
}

func TestRateLimit(test *testing.T) {
	test.Parallel()
	err := RateLimit(10, time.Millisecond, 0.5, NewRetryableFunc(true))
	require.ErrorIs(test, err, io.EOF)
}

func TestRateLimitWithContext(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	err := RateLimitWithContext(ctx, 10, time.Millisecond, 0.5, NewRetryableFunc(true))
	require.ErrorIs(test, err, io.EOF)
	cancel()
	err = RateLimitWithContext(ctx, 10, time.Millisecond, 0.5, NewRetryableFunc(true))
	require.ErrorIs(test, err, context.Canceled)
}
