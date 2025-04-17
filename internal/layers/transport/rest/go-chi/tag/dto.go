package tag

import "github.com/go-andiamo/chioas"

type Tag struct {
	ID          string `json:"id"`
	TagsSpaceID string `json:"tags_space_id"`
	Name        string `json:"name"`
}

var TagOpenApiDefinition = chioas.Schema{
	Name:               "tag",
	RequiredProperties: []string{"id", "tags_space_id", "name"},
	Properties: chioas.Properties{
		{
			Name:     "id",
			Type:     "string",
			Format:   "uuid",
			Required: true,
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:     "tags_space_id",
			Type:     "string",
			Format:   "uuid",
			Required: true,
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "name",
			Type:    "string",
			Example: "Tag Name",
		},
	},
}
