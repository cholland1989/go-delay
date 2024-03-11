package delay

import (
	"math"
	"time"
)

// RateLimit calculates the minimum duration per action for the specified
// actions per time period.
func RateLimit(actions int, period time.Duration) (delay time.Duration) {
	return FloatToDuration(float64(period) / math.Max(float64(actions), 1.0))
}
