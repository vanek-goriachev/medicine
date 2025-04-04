package retry

import "time"

type RetrierOption func(retrier *Retrier)

func MaxRetries(maxRetries uint16) RetrierOption {
	return func(retrier *Retrier) {
		retrier.MaxRetires = maxRetries
	}
}

func Delay(delay time.Duration) RetrierOption {
	return func(retrier *Retrier) {
		retrier.Delay = delay
	}
}
