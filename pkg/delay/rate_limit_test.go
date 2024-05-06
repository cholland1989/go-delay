package delay

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func ExampleRateLimit() {
	fmt.Println(RateLimit(1, time.Second))
	fmt.Println(RateLimit(2, time.Second))
	fmt.Println(RateLimit(4, time.Second))
	fmt.Println(RateLimit(8, time.Second))
	fmt.Println(RateLimit(10, time.Second))
	// Output:
	// 1s
	// 500ms
	// 250ms
	// 125ms
	// 100ms
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
			delay := RateLimit(params.actions, params.period)
			require.Equal(test, params.expected, delay)
		})
	}
}

func BenchmarkRateLimit(benchmark *testing.B) {
	for count := 0; count < benchmark.N; count++ {
		RateLimit(count, time.Millisecond)
	}
}
