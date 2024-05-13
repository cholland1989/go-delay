package retry

import "io"

func NewRetryableFunc(state bool) RetryableFunc {
	return func() error {
		state = !state
		if state {
			return io.EOF
		}
		return nil
	}
}
