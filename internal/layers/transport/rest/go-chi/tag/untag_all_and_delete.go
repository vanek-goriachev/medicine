package tag

import (
	"github.com/go-andiamo/chioas"
)

type TagUntagAllAndDeleteIn struct {
	ID string `json:"id"`
}

var TagUntagAllAndDeleteInOpenApiDefinition = chioas.Schema{
	Name:               "TagUntagAllAndDeleteIn",
	RequiredProperties: []string{"id"},
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000000",
		},
	},
}

type TagUntagAllAndDeleteOut struct{}

var TagUntagAllAndDeleteOutOpenApiDefinition = chioas.Schema{
	Name:       "TagUntagAllAndDeleteOut",
	Properties: chioas.Properties{},
}
