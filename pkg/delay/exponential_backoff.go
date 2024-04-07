// Package delay provides utility functions for calculating backoff and jitter.
package delay

import (
	"math"
	"time"
)

// ExponentialBackoff calculates the specified duration with exponential
// backoff.
func ExponentialBackoff(duration time.Duration, multiplier float64, attempt int) (delay time.Duration) {
	return FloatToDuration(float64(duration) * math.Pow(multiplier, math.Max(float64(attempt), 0.0)+1.0))
}
