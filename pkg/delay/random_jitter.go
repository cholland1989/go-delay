package delay

import (
	"math/rand"
	"sync"
	"time"
)

// Seed random number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec

// Guard access to random number generator.
var mutex = new(sync.Mutex)

// RandomJitter calculates the specified duration plus or minus random jitter.
func RandomJitter(duration time.Duration, jitter float64) (delay time.Duration) {
	if jitter != 0.0 {
		mutex.Lock()
		defer mutex.Unlock()
		return FloatToDuration(float64(duration) * (1.0 + jitter*(1.0-2.0*random.Float64())))
	}
	return duration
}
