package sleep

import (
	"context"
	"time"

	"github.com/cholland1989/go-delay/pkg/delay"
)

// RandomJitter pauses the current goroutine for the specified duration with
// random jitter.
func RandomJitter(duration time.Duration, jitter float64) {
	time.Sleep(delay.RandomJitter(duration, jitter))
}

// RandomJitterWithContext pauses the current goroutine for the specified
// duration with random jitter, or until the context is canceled.
func RandomJitterWithContext(ctx context.Context, duration time.Duration, jitter float64) (err error) {
	if ctx == nil {
		RandomJitter(duration, jitter)
	} else {
		timer := time.NewTimer(delay.RandomJitter(duration, jitter))
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return ctx.Err()
		case <-timer.C:
		}
	}
	return nil
}
