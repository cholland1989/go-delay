package sleep

import (
	"testing"
	"time"
)

func DatabaseQuery() error {
	return nil
}

func ExampleExponentialBackoff() {
	for attempt := 0; attempt < 5; attempt++ {
		err := DatabaseQuery()
		if err == nil {
			break
		}
		ExponentialBackoff(time.Second, 2.0, 0.5, attempt)
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
			if delay < params.expected {
				test.Fatalf("expected %v, got %v", params.expected, delay)
			}
		})
	}
}
