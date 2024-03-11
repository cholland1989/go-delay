package sleep

import (
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// RateLimit pauses the current goroutine for the minimum duration per action
// with random jitter.
func RateLimit(actions int, period time.Duration, jitter float64) {
	time.Sleep(delay.RandomJitter(delay.RateLimit(actions, period), jitter))
}
