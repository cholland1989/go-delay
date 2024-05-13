package retry

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/sleep"
)

// LinearBackoff attempts the specified function up to the specified number of
// retries, with linear backoff after each failed attempt.
func LinearBackoff(retries int, duration time.Duration, multiplier float64, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			sleep.LinearBackoff(duration, multiplier, jitter, attempt)
		}
	}
	return err
}

// LinearBackoffWithContext attempts the specified function up to the specified
// number of retries, or until the context is canceled, with linear backoff
// after each attempt.
func LinearBackoffWithContext(ctx context.Context, retries int, duration time.Duration, multiplier float64, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			err = sleep.LinearBackoffWithContext(ctx, duration, multiplier, jitter, attempt)
			if err != nil {
				break
			}
		}
	}
	return err
}
