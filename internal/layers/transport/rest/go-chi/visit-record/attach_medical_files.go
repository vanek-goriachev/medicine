package visit_record

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-andiamo/chioas"

	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file/dto"
	goChiTooling "medicine/internal/tooling/go-chi"
)

var ErrCantParseVisitRecordID = errors.New("cant parse visit_record_id (probably field is missing)")

type VisitRecordAttachMedicalFilesIn struct {
	VisitRecordID        string
	UploadedMedicalFiles []medicalFileChi.UploadedMedicalFile
}

var VisitRecordAttachMedicalFilesInOpenApiDefinition = chioas.Schema{
	Name:               "VisitRecordAttachMedicalFilesIn",
	RequiredProperties: []string{"visit_record_id", "uploaded_medical_files"},
	Properties: chioas.Properties{
		{
			Name:     "visit_record_id",
			Type:     "string",
			Required: true,
			Format:   "uuid",
			Example:  "00000000-0000-0000-0000-000000000001",
		},
		{
			Name:     "uploaded_medical_files",
			Type:     "array",
			ItemType: "string",
			Format:   "binary",
		},
	},
}

// ParseVisitRecordAttachMedicalFilesRequest parses the request body and returns a VisitRecordAttachMedicalFilesIn
// A custom parser is required, because it's a little bit tricky to work with files.
func ParseVisitRecordAttachMedicalFilesRequest(r *http.Request) (VisitRecordAttachMedicalFilesIn, error) {
	var zero VisitRecordAttachMedicalFilesIn

	err := r.ParseMultipartForm(goChiTooling.GB)
	if err != nil {
		return zero, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	visitRecordIDs, ok := r.MultipartForm.Value["visit_record_id"]
	if !ok || len(visitRecordIDs) == 0 {
		return zero, ErrCantParseVisitRecordID
	}

	uploadedFiles, err := medicalFileChi.ParseUploadedMedicalFiles(r.MultipartForm.File["uploaded_medical_files"])
	if err != nil {
		return zero, fmt.Errorf("failed to parse uploaded medical files: %w", err)
	}

	return VisitRecordAttachMedicalFilesIn{
		VisitRecordID:        visitRecordIDs[0],
		UploadedMedicalFiles: uploadedFiles,
	}, nil
}

type VisitRecordAttachMedicalFilesOut struct{}

// No VisitRecordAttachMedicalFilesOutOpenApiDefinition needed
