package delay

import (
	"math"
	"time"
)

// LinearBackoff calculates the specified duration with linear backoff.
func LinearBackoff(duration time.Duration, multiplier float64, attempt int) (delay time.Duration) {
	return FloatToDuration(float64(duration) * multiplier * (math.Max(float64(attempt), 0.0) + 1.0))
}
