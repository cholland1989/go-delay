package sleep

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func DatabaseQuery() error {
	return nil
}

func ExampleExponentialBackoff() {
	for attempt := 0; attempt <= 3; attempt++ {
		err := DatabaseQuery()
		if err == nil {
			break
		}
		if attempt < 3 {
			ExponentialBackoff(time.Second, 2.0, 0.5, attempt)
		}
	}
	// Output:
}

func ExampleExponentialBackoffWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for attempt := 0; attempt <= 3; attempt++ {
		err := DatabaseQuery()
		if err == nil {
			break
		}
		if attempt < 3 {
			err = ExponentialBackoffWithContext(ctx, time.Second, 2.0, 0.5, attempt)
			if err != nil {
				break
			}
		}
	}
	// Output:
}

func TestExponentialBackoff(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration   time.Duration
		multiplier float64
		attempt    int
		expected   time.Duration
	}{
		"Exponential Backoff (0)": {time.Millisecond, 2.0, 0, 2 * time.Millisecond},
		"Exponential Backoff (1)": {time.Millisecond, 2.0, 1, 4 * time.Millisecond},
		"Exponential Backoff (2)": {time.Millisecond, 2.0, 2, 8 * time.Millisecond},
		"Exponential Backoff (3)": {time.Millisecond, 2.0, 3, 16 * time.Millisecond},
		"Exponential Backoff (4)": {time.Millisecond, 2.0, 4, 32 * time.Millisecond},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			timestamp := time.Now()
			ExponentialBackoff(params.duration, params.multiplier, 0.0, params.attempt)
			delay := time.Since(timestamp)
			require.LessOrEqual(test, params.expected, delay)
		})
	}
}

func TestExponentialBackoffWithContext(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration   time.Duration
		multiplier float64
		attempt    int
		expected   time.Duration
	}{
		"Exponential Backoff (0)": {time.Millisecond, 2.0, 0, 2 * time.Millisecond},
		"Exponential Backoff (1)": {time.Millisecond, 2.0, 1, 4 * time.Millisecond},
		"Exponential Backoff (2)": {time.Millisecond, 2.0, 2, 8 * time.Millisecond},
		"Exponential Backoff (3)": {time.Millisecond, 2.0, 3, 16 * time.Millisecond},
		"Exponential Backoff (4)": {time.Millisecond, 2.0, 4, 32 * time.Millisecond},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			for _, ctx := range []context.Context{nil, context.Background()} {
				timestamp := time.Now()
				err := ExponentialBackoffWithContext(ctx, params.duration, params.multiplier, 0.0, params.attempt)
				delay := time.Since(timestamp)
				require.NoError(test, err)
				require.LessOrEqual(test, params.expected, delay)
			}
		})
	}
}

func TestExponentialBackoffWithContext_WithCancel(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timestamp := time.Now()
	err := ExponentialBackoffWithContext(ctx, time.Second, 1.0, 0.0, 0)
	delay := time.Since(timestamp)
	require.ErrorIs(test, err, context.Canceled)
	require.GreaterOrEqual(test, time.Millisecond, delay)
}
