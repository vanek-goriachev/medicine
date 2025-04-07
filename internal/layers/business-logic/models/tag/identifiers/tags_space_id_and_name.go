package identifiers

import (
	"fmt"

	entityID "medicine/pkg/entity-id"
)

type TagsSpaceIDAndNameIdentifier struct {
	Name        string
	TagsSpaceID entityID.EntityID
}

func (TagsSpaceIDAndNameIdentifier) Identifier() {}

func (u TagsSpaceIDAndNameIdentifier) String() string {
	return fmt.Sprintf(
		"identifier: tags_space_id=%q name=%q",
		u.TagsSpaceID,
		u.Name,
	)
}

func (u TagsSpaceIDAndNameIdentifier) Equals(otherRaw any) bool {
	other, ok := otherRaw.(TagsSpaceIDAndNameIdentifier)
	if !ok {
		return false
	}

	return u.Name == other.Name && u.TagsSpaceID == other.TagsSpaceID
}
