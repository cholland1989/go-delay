package sleep

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func NetworkRequest() error {
	return nil
}

func ExampleRandomJitter() {
	for attempt := 0; attempt <= 3; attempt++ {
		err := NetworkRequest()
		if err == nil {
			break
		}
		if attempt < 3 {
			RandomJitter(time.Second, 0.5)
		}
	}
	// Output:
}

func ExampleRandomJitterWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for attempt := 0; attempt <= 3; attempt++ {
		err := NetworkRequest()
		if err == nil {
			break
		}
		if attempt < 3 {
			err = RandomJitterWithContext(ctx, time.Second, 0.5)
			if err != nil {
				break
			}
		}
	}
	// Output:
}

func TestRandomJitter(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration time.Duration
		jitter   float64
	}{
		"Random Jitter (0)": {time.Millisecond, 0.0},
		"Random Jitter (1)": {time.Millisecond, 1.0},
		"Random Jitter (2)": {time.Millisecond, 2.0},
		"Random Jitter (3)": {time.Millisecond, 3.0},
		"Random Jitter (4)": {time.Millisecond, 4.0},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			timestamp := time.Now()
			RandomJitter(params.duration, params.jitter)
			delay := time.Since(timestamp)
			lower := params.duration - time.Duration(params.jitter*float64(params.duration))
			require.LessOrEqual(test, lower, delay)
		})
	}
}

func TestRandomJitterWithContext(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration time.Duration
		jitter   float64
	}{
		"Random Jitter (0)": {time.Millisecond, 0.0},
		"Random Jitter (1)": {time.Millisecond, 1.0},
		"Random Jitter (2)": {time.Millisecond, 2.0},
		"Random Jitter (3)": {time.Millisecond, 3.0},
		"Random Jitter (4)": {time.Millisecond, 4.0},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			for _, ctx := range []context.Context{nil, context.Background()} {
				timestamp := time.Now()
				err := RandomJitterWithContext(ctx, params.duration, params.jitter)
				delay := time.Since(timestamp)
				lower := params.duration - time.Duration(params.jitter*float64(params.duration))
				require.NoError(test, err)
				require.LessOrEqual(test, lower, delay)
			}
		})
	}
}

func TestRandomJitterWithContext_WithCancel(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timestamp := time.Now()
	err := RandomJitterWithContext(ctx, time.Second, 0.0)
	delay := time.Since(timestamp)
	require.ErrorIs(test, err, context.Canceled)
	require.GreaterOrEqual(test, time.Millisecond, delay)
}
