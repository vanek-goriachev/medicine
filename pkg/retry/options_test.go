package retry_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"medicine/pkg/retry"
)

func defaultRetrier() *retry.Retrier {
	return &retry.Retrier{
		MaxRetires: 3,
		Delay:      1 * time.Second,
	}
}

func TestRetrierDefaultOptions(t *testing.T) {
	t.Parallel()
	assert.Equal(t, defaultRetrier(), retry.NewRetrier())
}

func TestMaxRetriesOption(t *testing.T) {
	t.Parallel()

	expectedRetrier := defaultRetrier()
	expectedRetrier.MaxRetires = 9

	retrier := retry.NewRetrier(
		retry.MaxRetries(9),
	)

	assert.Equal(t, expectedRetrier, retrier)
}

func TestDelayOption(t *testing.T) {
	t.Parallel()

	expectedRetrier := defaultRetrier()
	expectedRetrier.Delay = 10 * time.Second

	retrier := retry.NewRetrier(
		retry.Delay(10 * time.Second),
	)

	assert.Equal(t, expectedRetrier, retrier)
}
