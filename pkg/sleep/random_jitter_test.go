package sleep

import (
	"context"
	"testing"
	"time"
)

func NetworkRequest() error {
	return nil
}

func ExampleRandomJitter() {
	for attempt := 0; attempt < 5; attempt++ {
		err := NetworkRequest()
		if err == nil {
			break
		}
		RandomJitter(time.Second, 0.5)
	}
	// Output:
}

func ExampleRandomJitterWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for attempt := 0; attempt < 5; attempt++ {
		err := NetworkRequest()
		if err == nil {
			break
		}
		RandomJitterWithContext(ctx, time.Second, 0.5)
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
			if delay < lower {
				test.Fatalf("%v less than %v", delay, lower)
			}
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
				RandomJitterWithContext(ctx, params.duration, params.jitter)
				delay := time.Since(timestamp)
				lower := params.duration - time.Duration(params.jitter*float64(params.duration))
				if delay < lower {
					test.Fatalf("%v less than %v", delay, lower)
				}
			}
		})
	}
}

func TestRandomJitterWithContext_WithCancel(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timestamp := time.Now()
	RandomJitterWithContext(ctx, time.Second, 0.0)
	delay := time.Since(timestamp)
	if delay > time.Millisecond {
		test.Fatalf("%v greater than %v", delay, time.Millisecond)
	}
}
