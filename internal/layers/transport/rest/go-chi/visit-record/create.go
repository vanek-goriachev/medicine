package visit_record

import (
	"github.com/go-andiamo/chioas"
	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file"
)

type VisitRecordCreateIn struct {
	Name     string `schema:"name"`
	Datetime string `schema:"datetime"`

	UploadedMedicalFile []medicalFileChi.UploadedMedicalFile `schema:"uploaded_medical_files"`

	TagIDs []string `schema:"tag_ids"`
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
			Name:      "uploaded_medical_files",
			Type:      "array",
			ItemType:  "object",
			SchemaRef: "uploaded-medical-file",
		},
		{
			Name:     "tag_ids",
			Type:     "array",
			ItemType: "string",
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
	RequiredProperties: []string{"visit_record"},
	Properties: chioas.Properties{
		{
			Name:      "visit_record",
			Type:      "object",
			SchemaRef: "tags-space",
		},
	},
}
