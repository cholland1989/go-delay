package sleep

import (
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// RandomJitter pauses the current goroutine for the specified duration with
// random jitter.
func RandomJitter(duration time.Duration, jitter float64) {
	time.Sleep(delay.RandomJitter(duration, jitter))
}
