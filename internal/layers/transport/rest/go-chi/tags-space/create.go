package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type TagsSpaceCreateIn struct {
	Name string `json:"name"`
}

var TagsSpaceCreateInOpenApiDefinition = chioas.Schema{
	Name: "TagsSpaceCreateIn",
	Properties: chioas.Properties{
		{
			Name: "name",
			Type: "string",
		},
	},
}

type TagsSpaceCreateOut struct {
	TagsSpace TagsSpace `json:"tags_space"`
}

var TagsSpaceCreateOutOpenApiDefinition = chioas.Schema{
	Name: "TagsSpaceCreateOut",
	Properties: chioas.Properties{
		{
			Name:      "tags_space",
			Type:      "object",
			SchemaRef: "tags-space",
		},
	},
}
