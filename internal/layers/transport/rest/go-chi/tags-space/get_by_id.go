package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type TagsSpaceGetByIDIn struct {
	ID string `json:"id"`
}

var TagsSpaceGetByIDInOpenApiDefinition = chioas.Schema{
	Name: "TagsSpaceGetByIDIn",
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000001",
		},
	},
}

type TagsSpaceGetByIDOut struct {
	TagsSpace TagsSpace `json:"tags_space"`
}

var TagsSpaceGetByIDOutOpenApiDefinition = chioas.Schema{
	Name: "TagsSpaceGetByIDOut",
	Properties: chioas.Properties{
		{
			Name:      "tags_space",
			Type:      "object",
			SchemaRef: "tags-space",
		},
	},
}
