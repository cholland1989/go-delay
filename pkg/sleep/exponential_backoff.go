// Package sleep provides wrapper functions to simplify the most common use
// case.
package sleep

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// ExponentialBackoff pauses the current goroutine for the specified duration
// with exponential backoff and random jitter.
func ExponentialBackoff(duration time.Duration, multiplier float64, jitter float64, attempt int) {
	RandomJitter(delay.ExponentialBackoff(duration, multiplier, attempt), jitter)
}

// ExponentialBackoffWithContext pauses the current goroutine for the specified
// duration with exponential backoff and random jitter, or until the context is
// canceled.
func ExponentialBackoffWithContext(ctx context.Context, duration time.Duration, multiplier float64, jitter float64, attempt int) (err error) {
	return RandomJitterWithContext(ctx, delay.ExponentialBackoff(duration, multiplier, attempt), jitter)
}
