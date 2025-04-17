package dto

import (
	"github.com/go-andiamo/chioas"
)

type MedicalFileInfo struct {
	ID        string `json:"id"`
	Extension string `json:"extension"`
	Name      string `json:"name"`
}

var MedicalFileInfoOpenApiDefinition = chioas.Schema{
	Name:               "medical-file-info",
	RequiredProperties: []string{"id", "extension", "name"},
	Properties: chioas.Properties{
		{
			Name:     "id",
			Type:     "string",
			Format:   "uuid",
			Required: true,
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "extension",
			Type:    "string",
			Example: ".png",
		},
		{
			Name:    "name",
			Type:    "string",
			Example: "Medical File Name",
		},
	},
}
