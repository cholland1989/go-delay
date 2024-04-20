package sleep

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// LinearBackoff pauses the current goroutine for the specified duration with
// linear backoff and random jitter.
func LinearBackoff(duration time.Duration, multiplier float64, jitter float64, attempt int) {
	RandomJitter(delay.LinearBackoff(duration, multiplier, attempt), jitter)
}

// LinearBackoffWithContext pauses the current goroutine for the specified
// duration with linear backoff and random jitter, or until the context is
// canceled.
func LinearBackoffWithContext(ctx context.Context, duration time.Duration, multiplier float64, jitter float64, attempt int) (err error) {
	return RandomJitterWithContext(ctx, delay.LinearBackoff(duration, multiplier, attempt), jitter)
}
