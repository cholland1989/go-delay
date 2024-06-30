package retry

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func NetworkRequest() error {
	return nil
}

func ExampleConstantBackoff() {
	err := ConstantBackoff(3, time.Second, 0.5, NetworkRequest)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func ExampleConstantBackoffWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := ConstantBackoffWithContext(ctx, 3, time.Second, 0.5, NetworkRequest)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func TestConstantBackoff(test *testing.T) {
	test.Parallel()
	err := ConstantBackoff(0, time.Millisecond, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = ConstantBackoff(1, time.Millisecond, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
}

func TestConstantBackoffWithContext(test *testing.T) {
	test.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	err := ConstantBackoffWithContext(ctx, 0, time.Millisecond, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, io.EOF)
	err = ConstantBackoffWithContext(ctx, 1, time.Millisecond, 0.5, NewRetryableFunc(false))
	require.NoError(test, err)
	cancel()
	err = ConstantBackoffWithContext(ctx, 1, time.Millisecond, 0.5, NewRetryableFunc(false))
	require.ErrorIs(test, err, context.Canceled)
}
