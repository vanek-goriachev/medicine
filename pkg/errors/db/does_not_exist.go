package db

import (
	"errors"
)

type DoesNotExistError struct {
	identifier Identifier
}

func NewDoesNotExistError(identifier Identifier) *DoesNotExistError {
	return &DoesNotExistError{
		identifier: identifier,
	}
}

func (e *DoesNotExistError) Error() string {
	return "does not exist for " + e.identifier.String()
}

func (e *DoesNotExistError) Is(otherRaw error) bool {
	var other *DoesNotExistError
	if !errors.As(otherRaw, &other) {
		return false
	}

	return e.identifier.Equals(other.identifier)
}
