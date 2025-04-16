package visit_record

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	visitRecordUA "medicine/internal/layers/business-logic/user-actions/visit-record"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper   userActionsMapper
	createUA visitRecordCreateUserAction
}

func NewService(
	mapper userActionsMapper,
	createUA visitRecordCreateUserAction,
) *Service {
	return &Service{
		mapper:   mapper,
		createUA: createUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	createHandler := goChiTooling.Handler[
		VisitRecordCreateIn,
		visitRecordUA.VisitRecordCreateIn,
		visitRecordUA.VisitRecordCreateOut,
		VisitRecordCreateOut,
	](
		goChiTooling.ProcessRequestMultipartFormData,
		s.mapper.VisitRecordCreateInFromChi,
		s.createUA,
		s.mapper.VisitRecordCreateOutToChi,
	)

	return chioas.Path{
		Paths: chioas.Paths{
			"/create": {
				Methods: chioas.Methods{
					http.MethodPost: {
						Description: "Эндпоинт для создания VisitRecord",
						Handler:     createHandler,
						Request: &chioas.Request{
							ContentType: "multipart/form-data",
							Schema:      VisitRecordCreateInOpenApiDefinition,
						},
						Responses: chioas.Responses{
							http.StatusCreated: {Schema: VisitRecordCreateOutOpenApiDefinition},
						},
					},
				},
			},
		},
	}
}
