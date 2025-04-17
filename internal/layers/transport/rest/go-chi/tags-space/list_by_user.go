package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type TagsSpaceListAllAvailableIn struct{}

// No TagsSpaceListAllAvailableInOpenApiDefinition needed

type TagsSpaceListAllAvailableOut struct {
	TagsSpaces []TagsSpace `json:"tags_spaces"`
}

var TagsSpaceListAllAvailableOutOpenApiDefinition = chioas.Schema{
	Name:               "TagsSpaceListAllAvailableOut",
	RequiredProperties: []string{"tags_spaces"},
	Properties: chioas.Properties{
		{
			Name:      "tags_spaces",
			Type:      "array",
			ItemType:  "object",
			SchemaRef: "tags-space",
		},
	},
}
