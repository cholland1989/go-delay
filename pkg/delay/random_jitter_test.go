package delay

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func ExampleRandomJitter() {
	random.Seed(0)
	fmt.Println(RandomJitter(time.Second, 0.5))
	fmt.Println(RandomJitter(time.Second, 0.5))
	fmt.Println(RandomJitter(time.Second, 0.5))
	fmt.Println(RandomJitter(time.Second, 0.5))
	fmt.Println(RandomJitter(time.Second, 0.5))
	// Output:
	// 554.80385ms
	// 1.255034914s
	// 844.043734ms
	// 1.44565616s
	// 1.132412793s
}

func TestRandomJitter(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		duration time.Duration
		jitter   float64
	}{
		"Random Jitter (0.0)":  {time.Millisecond, 0.0},
		"Random Jitter (0.25)": {time.Millisecond, 0.25},
		"Random Jitter (0.5)":  {time.Millisecond, 0.5},
		"Random Jitter (0.75)": {time.Millisecond, 0.75},
		"Random Jitter (1.0)":  {time.Millisecond, 1.0},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			delay := RandomJitter(params.duration, params.jitter)
			lower := params.duration - time.Duration(params.jitter*float64(params.duration))
			require.LessOrEqual(test, lower, delay)
			upper := params.duration + time.Duration(params.jitter*float64(params.duration))
			require.GreaterOrEqual(test, upper, delay)
		})
	}
}

func BenchmarkRandomJitter(benchmark *testing.B) {
	for count := 0; count < benchmark.N; count++ {
		RandomJitter(time.Millisecond, float64(count))
	}
}
