package authorization

type UnauthorizedError struct {
	err error
}

func NewUnauthorizedError(
	err error,
) *UnauthorizedError {
	return &UnauthorizedError{
		err: err,
	}
}

func (e *UnauthorizedError) Error() string {
	return "unauthorized " + e.err.Error()
}

func (e *UnauthorizedError) Unwrap() error {
	return e.err
}
