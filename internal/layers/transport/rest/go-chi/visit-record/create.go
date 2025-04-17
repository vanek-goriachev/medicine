package visit_record

import (
	"github.com/go-andiamo/chioas"
)

type VisitRecordCreateIn struct {
	Name     string `json:"name"`
	Datetime string `json:"datetime"`

	TagIDs []string `json:"tag_ids"` //nolint:tagliatelle // False positive
}

var VisitRecordCreateInOpenApiDefinition = chioas.Schema{
	Name:               "VisitRecordCreateIn",
	RequiredProperties: []string{"name", "datetime"},
	Properties: chioas.Properties{
		{
			Name:    "name",
			Type:    "string",
			Example: "Visit Record Name",
		},
		{
			Name:    "datetime",
			Type:    "string",
			Example: "2006-01-02T15:04:05Z",
		},
		{
			Name:     "tag_ids",
			Type:     "array",
			ItemType: "string",
			Format:   "uuid",
			Example:  "00000000-0000-0000-0000-000000000001",
		},
	},
}

type VisitRecordCreateOut struct {
	VisitRecord               VisitRecord               `json:"visit_record"`
	VisitRecordLinkedEntities VisitRecordLinkedEntities `json:"visit_record_linked_entities"`
}

var VisitRecordCreateOutOpenApiDefinition = chioas.Schema{
	Name:               "VisitRecordCreateOut",
	RequiredProperties: []string{"visit_record", "visit_record_linked_entities"},
	Properties: chioas.Properties{
		{
			Name:      "visit_record",
			Type:      "object",
			SchemaRef: "visit-record",
		},
		{
			Name:      "visit_record_linked_entities",
			Type:      "object",
			SchemaRef: "visit-record-linked-entities",
		},
	},
}
