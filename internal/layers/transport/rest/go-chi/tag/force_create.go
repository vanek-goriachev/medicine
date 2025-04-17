package tag

import (
	"github.com/go-andiamo/chioas"
)

type TagForceCreateIn struct {
	Name        string `json:"name"`
	TagsSpaceID string `json:"tags_space_id"`
}

var TagForceCreateInOpenApiDefinition = chioas.Schema{
	Name:               "TagForceCreateIn",
	RequiredProperties: []string{"name", "tags_space_id"},
	Properties: chioas.Properties{
		{
			Name: "name",
			Type: "string",
		},
		{
			Name:     "tags_space_id",
			Type:     "string",
			Format:   "uuid",
			Required: true,
			Example:  "00000000-0000-0000-0000-000000000000",
		},
	},
}

type TagForceCreateOut struct {
	Tag Tag `json:"tag"`
}

var TagForceCreateOutOpenApiDefinition = chioas.Schema{
	Name:               "TagForceCreateOut",
	RequiredProperties: []string{"tag"},
	Properties: chioas.Properties{
		{
			Name:      "tag",
			Type:      "object",
			SchemaRef: "tag",
		},
	},
}
