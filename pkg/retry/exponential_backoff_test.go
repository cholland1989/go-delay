package retry

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func DatabaseQuery() error {
	return nil
}

func ExampleExponentialBackoff() {
	err := ExponentialBackoff(3, time.Second, 2.0, 0.5, DatabaseQuery)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func ExampleExponentialBackoffWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := ExponentialBackoffWithContext(ctx, 3, time.Second, 2.0, 0.5, DatabaseQuery)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func TestExponentialBackoff(test *testing.T) {
	test.Parallel()
	err := ExponentialBackoff(0, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = ExponentialBackoff(1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
}

func TestExponentialBackoffWithContext(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	err := ExponentialBackoffWithContext(ctx, 0, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = ExponentialBackoffWithContext(ctx, 1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
	cancel()
	err = ExponentialBackoffWithContext(ctx, 1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, context.Canceled)
}
