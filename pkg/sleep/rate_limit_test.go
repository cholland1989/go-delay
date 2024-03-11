package sleep

import (
	"testing"
	"time"
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
			if delay < params.expected {
				test.Fatalf("expected %v, got %v", params.expected, delay)
			}
		})
	}
}
