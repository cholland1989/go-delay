package retry

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/sleep"
)

// RateLimit attempts the specified function repeatedly, limited to the
// specified number of actions per time period, returning an error if the
// function fails.
func RateLimit(actions int, period time.Duration, jitter float64, function RetryableFunc) (err error) {
	for err = function(); err == nil; err = function() {
		sleep.RateLimit(actions, period, jitter)
	}
	return err
}

// RateLimitWithContext attempts the specified function repeatedly, limited to
// the specified number of actions per time period, returning an error if the
// function fails or the context is canceled.
func RateLimitWithContext(ctx context.Context, actions int, period time.Duration, jitter float64, function RetryableFunc) (err error) {
	for err = function(); err == nil; err = function() {
		err = sleep.RateLimitWithContext(ctx, actions, period, jitter)
		if err != nil {
			break
		}
	}
	return err
}
