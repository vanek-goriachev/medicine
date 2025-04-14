package tags_space

import (
	"github.com/go-andiamo/chioas"
)

type TagsSpaceDeleteIn struct {
	ID string `json:"id"`
}

var TagsSpaceDeleteInOpenApiDefinition = chioas.Schema{
	Name:               "TagsSpaceDeleteIn",
	RequiredProperties: []string{"id"},
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000000",
		},
	},
}

type TagsSpaceDeleteOut struct{}

var TagsSpaceDeleteOutOpenApiDefinition = chioas.Schema{
	Name:       "TagsSpaceDeleteOut",
	Properties: chioas.Properties{},
}
