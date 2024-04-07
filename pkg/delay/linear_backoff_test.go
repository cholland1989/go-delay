package delay

import (
	"fmt"
	"testing"
	"time"
)

func ExampleLinearBackoff() {
	fmt.Println(LinearBackoff(time.Second, 2.0, 0))
	fmt.Println(LinearBackoff(time.Second, 2.0, 1))
	fmt.Println(LinearBackoff(time.Second, 2.0, 2))
	fmt.Println(LinearBackoff(time.Second, 2.0, 3))
	fmt.Println(LinearBackoff(time.Second, 2.0, 4))
	// Output:
	// 2s
	// 4s
	// 6s
	// 8s
	// 10s
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
			delay := LinearBackoff(params.duration, params.multiplier, params.attempt)
			if delay != params.expected {
				test.Fatalf("expected %v, got %v", params.expected, delay)
			}
		})
	}
}

func BenchmarkLinearBackoff(benchmark *testing.B) {
	for count := 0; count < benchmark.N; count++ {
		LinearBackoff(time.Millisecond, float64(count), benchmark.N-count)
	}
}
