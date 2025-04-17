package visit_record

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	visitRecordUA "medicine/internal/layers/business-logic/user-actions/visit-record"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper               userActionsMapper
	createUA             visitRecordCreateUserAction
	attachMedicalFilesUA visitRecordAttachMedicalFilesUserAction
}

func NewService(
	mapper userActionsMapper,
	createUA visitRecordCreateUserAction,
	attachMedicalFilesUA visitRecordAttachMedicalFilesUserAction,
) *Service {
	return &Service{
		mapper:               mapper,
		createUA:             createUA,
		attachMedicalFilesUA: attachMedicalFilesUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	createHandler := goChiTooling.Handler[
		VisitRecordCreateIn,
		visitRecordUA.VisitRecordCreateIn,
		visitRecordUA.VisitRecordCreateOut,
		VisitRecordCreateOut,
	](
		goChiTooling.ProcessRequestBody,
		s.mapper.VisitRecordCreateInFromChi,
		s.createUA,
		s.mapper.VisitRecordCreateOutToChi,
	)

	attachmedicalFilesHandler := goChiTooling.Handler[
		VisitRecordAttachMedicalFilesIn,
		visitRecordUA.VisitRecordAttachMedicalFilesIn,
		visitRecordUA.VisitRecordAttachMedicalFilesOut,
		VisitRecordAttachMedicalFilesOut,
	](
		ParseVisitRecordAttachMedicalFilesRequest,
		s.mapper.VisitRecordAttachMedicalFilesInFromChi,
		s.attachMedicalFilesUA,
		s.mapper.VisitRecordAttachMedicalFilesOutToChi,
	)

	return chioas.Path{
		Paths: chioas.Paths{
			"/create": {
				Methods: chioas.Methods{
					http.MethodPost: {
						Description: "Эндпоинт для создания VisitRecord",
						Handler:     createHandler,
						Request: &chioas.Request{
							Schema: VisitRecordCreateInOpenApiDefinition,
						},
						Responses: chioas.Responses{
							http.StatusCreated: {Schema: VisitRecordCreateOutOpenApiDefinition},
						},
					},
				},
			},
			"/attach-medical-files": {
				Methods: chioas.Methods{
					http.MethodPost: {
						Description: "Эндпоинт для добавления медицинских файлов к VisitRecord",
						Handler:     attachmedicalFilesHandler,
						Request: &chioas.Request{
							ContentType: "multipart/form-data",
							Schema:      VisitRecordAttachMedicalFilesInOpenApiDefinition,
						},
						Responses: chioas.Responses{
							http.StatusCreated: {},
						},
					},
				},
			},
		},
	}
}
