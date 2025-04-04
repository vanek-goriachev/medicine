package retry

import "time"

type Retrier struct {
	MaxRetires uint16
	Delay      time.Duration
}

func (r *Retrier) Wrap(f func() error) func() error {
	return func() error {
		var err error

		failsCounter := uint16(0)
		for failsCounter < r.MaxRetires {
			err = f()

			if err != nil {
				time.Sleep(r.Delay)

				failsCounter++

				continue
			}

			break
		}

		return err
	}
}

func NewRetrier(opts ...RetrierOption) *Retrier {
	//nolint:mnd // Its not magic numbers, just default ones
	retrier := &Retrier{
		// default values
		MaxRetires: 3,
		Delay:      1 * time.Second,
	}

	for _, opt := range opts {
		opt(retrier)
	}

	return retrier
}
