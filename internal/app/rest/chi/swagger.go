package chi

import (
	"github.com/go-andiamo/chioas"
	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file"
	visitRecordChi "medicine/internal/layers/transport/rest/go-chi/visit-record"

	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
	tagsSpaceChi "medicine/internal/layers/transport/rest/go-chi/tags-space"
)

func generateApiSpec(services *chiServices) chioas.Definition {
	apiDefinition := chioas.Definition{
		AutoHeadMethods: true,
		DocOptions: chioas.DocOptions{
			UIStyle:         chioas.Swagger,
			ServeDocs:       true,
			HideHeadMethods: true,
		},
		Paths: chioas.Paths{
			"/api/v1": {
				Paths: chioas.Paths{
					"/tags-space":   services.tagsSpace.GenerateOpenApiDefinition(),
					"/tag":          services.tag.GenerateOpenApiDefinition(),
					"/visit-record": services.visitRecord.GenerateOpenApiDefinition(),
				},
			},
		},
		Components: &chioas.Components{
			Schemas: chioas.Schemas{
				tagChi.TagOpenApiDefinition,
				tagsSpaceChi.TagsSpaceOpenApiDefinition,

				medicalFileChi.UploadedMedicalFileOpenApiDefinition,
				medicalFileChi.MedicalFileInfoOpenApiDefinition,

				visitRecordChi.VisitRecordOpenApiDefinition,
				visitRecordChi.VisitRecordLinkedEntitiesOpenApiDefinition,
			},
		},
	}

	return apiDefinition
}
