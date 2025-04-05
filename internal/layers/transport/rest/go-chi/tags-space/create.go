package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type CreateTagsSpaceIn struct {
	Name string `json:"name"`
}

var CreateTagsSpaceInOpenApiDefinition = chioas.Schema{
	Name: "CreateTagsSpaceIn",
	Properties: chioas.Properties{
		{
			Name: "name",
			Type: "string",
		},
	},
}

type CreateTagsSpaceOut struct {
	TagsSpace TagsSpace `json:"tags_space"`
}

var CreateTagsSpaceOutOpenApiDefinition = chioas.Schema{
	Name: "CreateTagsSpaceOut",
	Properties: chioas.Properties{
		{
			Name:      "tags_space",
			Type:      "object",
			SchemaRef: "tags-space",
		},
	},
}
