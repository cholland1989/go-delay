package sleep

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// RateLimit pauses the current goroutine for the minimum duration per action
// with random jitter.
func RateLimit(actions int, period time.Duration, jitter float64) {
	RandomJitter(delay.RateLimit(actions, period), jitter)
}

// RateLimitWithContext pauses the current goroutine for the minimum duration
// per action with random jitter, or until the context is canceled.
func RateLimitWithContext(ctx context.Context, actions int, period time.Duration, jitter float64) (err error) {
	return RandomJitterWithContext(ctx, delay.RateLimit(actions, period), jitter)
}
