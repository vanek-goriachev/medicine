package tag

import (
	"errors"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
)

type TagAlreadyExistError struct {
	identifier customIdentifiers.Identifier
}

func NewTagAlreadyExistError(identifier customIdentifiers.Identifier) *TagAlreadyExistError {
	return &TagAlreadyExistError{identifier: identifier}
}

func (e *TagAlreadyExistError) Error() string {
	return "Tag with identifier " + e.identifier.String() + " already exists"
}

func (e *TagAlreadyExistError) Is(otherRaw error) bool {
	var other *TagAlreadyExistError

	if !errors.As(otherRaw, &other) {
		return false
	}

	return e.identifier.Equals(other.identifier)
}

func (e *TagAlreadyExistError) Unwrap() error {
	return e
}
