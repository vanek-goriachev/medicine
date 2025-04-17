package visit_record

import (
	"github.com/go-andiamo/chioas"
)

type VisitRecord struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Datetime string `json:"datetime"`
}

var VisitRecordOpenApiDefinition = chioas.Schema{
	Name:               "visit-record",
	RequiredProperties: []string{"id", "name", "datetime"},
	Properties: chioas.Properties{
		{
			Name:    "id",
			Type:    "string",
			Example: "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:    "name",
			Type:    "string",
			Example: "Visit Record Name",
		},
		{
			Name:    "datetime",
			Type:    "string",
			Format:  "date-time",
			Example: "2020-01-01T00:00:00Z",
		},
	},
}

type VisitRecordLinkedEntities struct {
	TagIDs         []string `json:"tag_ids"`          //nolint:tagliatelle // False positive
	MedicalFileIDs []string `json:"medical_file_ids"` //nolint:tagliatelle // False positive
}

var VisitRecordLinkedEntitiesOpenApiDefinition = chioas.Schema{
	Name:               "visit-record-linked-entities",
	RequiredProperties: []string{"tag_ids", "medical_file_ids"},
	Properties: chioas.Properties{
		{
			Name:     "tag_ids",
			Type:     "array",
			ItemType: "string",
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:     "medical_file_ids",
			Type:     "array",
			ItemType: "string",
			Example:  "00000000-0000-0000-0000-000000000001",
		},
	},
}
