package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type TagsSpaceListByUserIn struct{}

// No TagsSpaceListByUserInOpenApiDefinition needed

type TagsSpaceListByUserOut struct {
	TagsSpaces []TagsSpace `json:"tags_spaces"`
}

var TagsSpaceListByUserOutOpenApiDefinition = chioas.Schema{
	Name: "TagsSpaceListByUserOut",
	Properties: chioas.Properties{
		{
			Name:      "tags_spaces",
			Type:      "array",
			ItemType:  "object",
			SchemaRef: "tags-space",
		},
	},
}
