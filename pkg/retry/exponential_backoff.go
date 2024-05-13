package retry

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/sleep"
)

// ExponentialBackoff attempts the specified function up to the specified
// number of retries, with exponential backoff after each failed attempt.
func ExponentialBackoff(retries int, duration time.Duration, multiplier float64, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			sleep.ExponentialBackoff(duration, multiplier, jitter, attempt)
		}
	}
	return err
}

// ExponentialBackoffWithContext attempts the specified function up to the
// specified number of retries, or until the context is canceled, with
// exponential backoff after each attempt.
func ExponentialBackoffWithContext(ctx context.Context, retries int, duration time.Duration, multiplier float64, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			err = sleep.ExponentialBackoffWithContext(ctx, duration, multiplier, jitter, attempt)
			if err != nil {
				break
			}
		}
	}
	return err
}
