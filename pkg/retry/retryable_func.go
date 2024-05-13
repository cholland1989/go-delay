package retry

// RetryableFunc defines a retryable function.
type RetryableFunc func() error
