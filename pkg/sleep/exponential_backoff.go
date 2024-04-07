// Package sleep provides wrapper functions to simplify the most common use
// case, passing the calculated delay to [time.Sleep].
package sleep

import (
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// ExponentialBackoff pauses the current goroutine for the specified duration
// with exponential backoff and random jitter.
func ExponentialBackoff(duration time.Duration, multiplier float64, jitter float64, attempt int) {
	time.Sleep(delay.RandomJitter(delay.ExponentialBackoff(duration, multiplier, attempt), jitter))
}
