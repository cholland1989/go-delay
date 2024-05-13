// Package retry provides wrapper functions to simplify the most common use
// cases.
package retry

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/sleep"
)

// ConstantBackoff attempts the specified function up to the specified number
// of retries, with constant backoff after each failed attempt.
func ConstantBackoff(retries int, duration time.Duration, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			sleep.RandomJitter(duration, jitter)
		}
	}
	return err
}

// ConstantBackoffWithContext attempts the specified function up to the
// specified number of retries, or until the context is canceled, with constant
// backoff after each attempt.
func ConstantBackoffWithContext(ctx context.Context, retries int, duration time.Duration, jitter float64, function RetryableFunc) (err error) {
	for attempt := 0; attempt <= retries; attempt++ {
		err = function()
		if err == nil {
			break
		}
		if attempt < retries {
			err = sleep.RandomJitterWithContext(ctx, duration, jitter)
			if err != nil {
				break
			}
		}
	}
	return err
}
