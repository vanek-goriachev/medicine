//nolint:err113,revive,goconst // tests
package retry_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"medicine/pkg/retry"
)

func Test_WrappedFunctionRuns_ExactlyOnce_WhenNoErrors(t *testing.T) {
	t.Parallel()

	retrier := retry.NewRetrier(
		retry.MaxRetries(3),
		retry.Delay(0),
	)
	callsCounter := 0
	someFunc := func() error {
		callsCounter++

		return nil
	}

	wrappedFunc := retrier.Wrap(someFunc)
	err := wrappedFunc()

	assert.NoError(t, err)
	assert.Equal(t, 1, callsCounter)
}

func Test_WrappedFunctionRuns_UntilSuccess_WhenLessErrorsThanRetries(t *testing.T) {
	t.Parallel()

	retrier := retry.NewRetrier(
		retry.MaxRetries(3),
		retry.Delay(0),
	)
	callsCounter := 0
	someFunc := func() error {
		callsCounter++
		if callsCounter <= 2 {
			return errors.New("some error")
		}

		return nil
	}

	wrappedFunc := retrier.Wrap(someFunc)
	err := wrappedFunc()

	assert.NoError(t, err)
	assert.Equal(t, 3, callsCounter)
}

func Test_WrappedFunction_Fails_WhenMoreErrorsThanRetries(t *testing.T) {
	t.Parallel()

	retrier := retry.NewRetrier(
		retry.MaxRetries(3),
		retry.Delay(0),
	)
	callsCounter := 0
	someFunc := func() error {
		callsCounter++

		return fmt.Errorf("some error %d", callsCounter) //nolint:goconst // test code
	}

	wrappedFunc := retrier.Wrap(someFunc)
	err := wrappedFunc()

	assert.EqualError(t, err, "some error 3", "must return the last error of wrapped func")
	assert.Equal(t, 3, callsCounter)
}

func Test_WrappedFunctionRuns_FailsFewTimes_WithExpectedDelays(t *testing.T) {
	t.Parallel()

	retries := 4
	delay := time.Millisecond * 500

	retrier := retry.NewRetrier(
		retry.MaxRetries(uint16(retries)),
		retry.Delay(delay),
	)

	callsCounter := 0
	var previousCallTime time.Time
	someFunc := func() error {
		if callsCounter != 0 {
			assert.GreaterOrEqual(t, time.Since(previousCallTime), delay, "must wait the expected delay")
		}

		callsCounter += 1
		previousCallTime = time.Now()

		return fmt.Errorf("some error %d", callsCounter)
	}

	wrappedFunc := retrier.Wrap(someFunc)

	err := wrappedFunc()

	assert.EqualError(t, err, "some error 4", "must return the last error of wrapped func")
	assert.Equal(t, retries, callsCounter)
}
