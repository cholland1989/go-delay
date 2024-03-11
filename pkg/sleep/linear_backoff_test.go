package sleep

import (
	"testing"
	"time"
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
			if delay < params.expected {
				test.Fatalf("expected %v, got %v", params.expected, delay)
			}
		})
	}
}
