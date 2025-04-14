package tags_space

import (
	"errors"
	"fmt"
	
	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	entityID "medicine/pkg/entity-id"
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

type TagsSpaceHaveTagsError struct {
	id entityID.EntityID
}

func NewTagsSpaceHaveTagsError(id entityID.EntityID) *TagsSpaceHaveTagsError {
	return &TagsSpaceHaveTagsError{
		id: id,
	}
}

func (e *TagsSpaceHaveTagsError) Error() string {
	return fmt.Sprintf("Tags space with id %q have tags", e.id.String())
}

func (e *TagsSpaceHaveTagsError) Is(otherRaw error) bool {
	var other *TagsSpaceHaveTagsError

	if !errors.As(otherRaw, &other) {
		return false
	}

	return e.id.Equals(other.id)
}

func (e *TagsSpaceHaveTagsError) Unwrap() error {
	return e
}
