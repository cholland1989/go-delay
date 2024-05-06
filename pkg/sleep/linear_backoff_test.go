package sleep

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func LockFile() error {
	return nil
}

func ExampleLinearBackoff() {
	for attempt := 0; attempt < 5; attempt++ {
		err := LockFile()
		if err == nil {
			break
		}
		LinearBackoff(time.Second, 2.0, 0.5, attempt)
	}
	// Output:
}

func ExampleLinearBackoffWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for attempt := 0; attempt < 5; attempt++ {
		err := LockFile()
		if err == nil {
			break
		}
		err = LinearBackoffWithContext(ctx, time.Second, 2.0, 0.5, attempt)
		if err != nil {
			break
		}
	}
	// Output:
}

func TestLinearBackoff(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration   time.Duration
		multiplier float64
		attempt    int
		expected   time.Duration
	}{
		"Linear Backoff (0)": {time.Millisecond, 2.0, 0, 2 * time.Millisecond},
		"Linear Backoff (1)": {time.Millisecond, 2.0, 1, 4 * time.Millisecond},
		"Linear Backoff (2)": {time.Millisecond, 2.0, 2, 6 * time.Millisecond},
		"Linear Backoff (3)": {time.Millisecond, 2.0, 3, 8 * time.Millisecond},
		"Linear Backoff (4)": {time.Millisecond, 2.0, 4, 10 * time.Millisecond},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			timestamp := time.Now()
			LinearBackoff(params.duration, params.multiplier, 0.0, params.attempt)
			delay := time.Since(timestamp)
			require.LessOrEqual(test, params.expected, delay)
		})
	}
}

func TestLinearBackoffWithContext(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration   time.Duration
		multiplier float64
		attempt    int
		expected   time.Duration
	}{
		"Linear Backoff (0)": {time.Millisecond, 2.0, 0, 2 * time.Millisecond},
		"Linear Backoff (1)": {time.Millisecond, 2.0, 1, 4 * time.Millisecond},
		"Linear Backoff (2)": {time.Millisecond, 2.0, 2, 6 * time.Millisecond},
		"Linear Backoff (3)": {time.Millisecond, 2.0, 3, 8 * time.Millisecond},
		"Linear Backoff (4)": {time.Millisecond, 2.0, 4, 10 * time.Millisecond},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			for _, ctx := range []context.Context{nil, context.Background()} {
				timestamp := time.Now()
				err := LinearBackoffWithContext(ctx, params.duration, params.multiplier, 0.0, params.attempt)
				delay := time.Since(timestamp)
				require.NoError(test, err)
				require.LessOrEqual(test, params.expected, delay)
			}
		})
	}
}

func TestLinearBackoffWithContext_WithCancel(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timestamp := time.Now()
	err := LinearBackoffWithContext(ctx, time.Second, 1.0, 0.0, 0)
	delay := time.Since(timestamp)
	require.ErrorIs(test, err, context.Canceled)
	require.GreaterOrEqual(test, time.Millisecond, delay)
}
