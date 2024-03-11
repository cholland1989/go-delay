package sleep

import (
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// LinearBackoff pauses the current goroutine for the specified duration with
// linear backoff and random jitter.
func LinearBackoff(duration time.Duration, multiplier float64, jitter float64, attempt int) {
	time.Sleep(delay.RandomJitter(delay.LinearBackoff(duration, multiplier, attempt), jitter))
}
