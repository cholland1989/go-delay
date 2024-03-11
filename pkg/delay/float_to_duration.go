package delay

import (
	"math"
	"time"
)

// FloatToDuration prevents overflow when converting a float64 to a
// [time.Duration] by clamping the result between [math.MinInt64] and
// [math.MaxInt64].
func FloatToDuration(value float64) (duration time.Duration) {
	if value <= math.MinInt64 {
		return time.Duration(math.MinInt64)
	} else if value >= math.MaxInt64 {
		return time.Duration(math.MaxInt64)
	}
	return time.Duration(value)
}
