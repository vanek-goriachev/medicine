package tags_space

import (
	"github.com/go-andiamo/chioas"

	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
)

type TagsSpace struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	Tags []tagChi.Tag `json:"tags"`
}

var TagsSpaceOpenApiDefinition = chioas.Schema{
	Name:               "tags-space",
	RequiredProperties: []string{"id", "name", "tags"},
	Properties: chioas.Properties{
		{
			Name:     "id",
			Type:     "string",
			Format:   "uuid",
			Required: true,
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "name",
			Type:    "string",
			Example: "Tags Space Name",
		},
		{
			Name:      "tags",
			Type:      "array",
			ItemType:  "object",
			SchemaRef: "tag",
		},
	},
}
