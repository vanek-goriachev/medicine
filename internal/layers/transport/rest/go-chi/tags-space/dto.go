package tags_space

import (
	"github.com/go-andiamo/chioas"

	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
)

type TagsSpace struct {
	ID     string       `json:"id"`
	UserID string       `json:"user_id"`
	Name   string       `json:"name"`
	Tags   []tagChi.Tag `json:"tags"`
}

var TagsSpaceOpenApiDefinition = chioas.Schema{
	Name:               "tags-space",
	RequiredProperties: []string{"id", "user_id", "name", "tags"},
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "user_id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000001",
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
