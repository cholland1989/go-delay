package retry

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func LockFile() error {
	return nil
}

func ExampleLinearBackoff() {
	err := LinearBackoff(3, time.Second, 2.0, 0.5, LockFile)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func ExampleLinearBackoffWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := LinearBackoffWithContext(ctx, 3, time.Second, 2.0, 0.5, LockFile)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func TestLinearBackoff(test *testing.T) {
	test.Parallel()
	err := LinearBackoff(0, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = LinearBackoff(1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
}

func TestLinearBackoffWithContext(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	err := LinearBackoffWithContext(ctx, 0, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = LinearBackoffWithContext(ctx, 1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
	cancel()
	err = LinearBackoffWithContext(ctx, 1, time.Millisecond, 2.0, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, context.Canceled)
}
