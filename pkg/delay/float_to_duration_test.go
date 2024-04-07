package delay

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func ExampleFloatToDuration() {
	fmt.Println(FloatToDuration(float64(math.MinInt64) * 2.0))
	fmt.Println(FloatToDuration(float64(math.MinInt64) - 1.0))
	fmt.Println(FloatToDuration(0))
	fmt.Println(FloatToDuration(float64(math.MaxInt64) + 1.0))
	fmt.Println(FloatToDuration(float64(math.MaxInt64) * 2.0))
	// Output:
	// -2562047h47m16.854775808s
	// -2562047h47m16.854775808s
	// 0s
	// 2562047h47m16.854775807s
	// 2562047h47m16.854775807s
}

func TestFloatToDuration(test *testing.T) {
	test.Parallel()
	for name, params := range map[string]struct {
		value    float64
		expected time.Duration
	}{
		"Float To Duration (Minimum)":  {float64(math.MinInt64) - 1.0, time.Duration(math.MinInt64)},
		"Float To Duration (Negative)": {-1.0, time.Duration(-1.0)},
		"Float To Duration (Zero)":     {0.0, time.Duration(0.0)},
		"Float To Duration (Positive)": {1.0, time.Duration(1.0)},
		"Float To Duration (Maximum)":  {float64(math.MaxInt64) + 1.0, time.Duration(math.MaxInt64)},
	} {
		params := params
		test.Run(name, func(test *testing.T) {
			test.Parallel()
			duration := FloatToDuration(params.value)
			if duration != params.expected {
				test.Fatalf("expected %v, got %v", params.expected, duration)
			}
		})
	}
}

func BenchmarkFloatToDuration(benchmark *testing.B) {
	for count := 0; count < benchmark.N; count++ {
		FloatToDuration(float64(count))
	}
}
