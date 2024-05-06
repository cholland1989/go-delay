package delay

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func ExampleExponentialBackoff() {
	fmt.Println(ExponentialBackoff(time.Second, 2.0, 0))
	fmt.Println(ExponentialBackoff(time.Second, 2.0, 1))
	fmt.Println(ExponentialBackoff(time.Second, 2.0, 2))
	fmt.Println(ExponentialBackoff(time.Second, 2.0, 3))
	fmt.Println(ExponentialBackoff(time.Second, 2.0, 4))
	// Output:
	// 2s
	// 4s
	// 8s
	// 16s
	// 32s
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
			delay := ExponentialBackoff(params.duration, params.multiplier, params.attempt)
			require.Equal(test, params.expected, delay)
		})
	}
}

func BenchmarkExponentialBackoff(benchmark *testing.B) {
	for count := 0; count < benchmark.N; count++ {
		ExponentialBackoff(time.Millisecond, float64(count), benchmark.N-count)
	}
}
