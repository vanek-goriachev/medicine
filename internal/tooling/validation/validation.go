package validation

import (
	"errors"
)

type ValidationError struct {
	err error
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{err: err}
}

func (e *ValidationError) Error() string {
	return e.err.Error()
}

func (e *ValidationError) Is(otherRaw error) bool {
	var other *ValidationError

	if !errors.As(otherRaw, &other) {
		return false
	}

	return errors.Is(e.err, other.err)
}
