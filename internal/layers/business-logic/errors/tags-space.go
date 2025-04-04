package errors

import (
	"errors"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
)

type TagsSpaceAlreadyExistError struct {
	identifier customIdentifiers.Identifier
}

func NewTagsSpaceAlreadyExistError(identifier customIdentifiers.Identifier) *TagsSpaceAlreadyExistError {
	return &TagsSpaceAlreadyExistError{identifier: identifier}
}

func (e *TagsSpaceAlreadyExistError) Error() string {
	return "Tags space with identifier " + e.identifier.String() + " already exists"
}

func (e *TagsSpaceAlreadyExistError) Is(otherRaw error) bool {
	var other *TagsSpaceAlreadyExistError

	if !errors.As(otherRaw, &other) {
		return false
	}

	return e.identifier.Equals(other.identifier)
}

func (e *TagsSpaceAlreadyExistError) Unwrap() error {
	return e
}
