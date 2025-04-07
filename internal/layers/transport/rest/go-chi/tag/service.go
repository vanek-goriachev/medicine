package tag

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper        userActionsMapper
	forceCreateUA tagCreateUserAction
}

func NewService(
	mapper userActionsMapper,
	forceCreateUA tagCreateUserAction,
) *Service {
	return &Service{
		mapper:        mapper,
		forceCreateUA: forceCreateUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	forceCreateHandler := goChiTooling.Handler[
		TagForceCreateIn,
		tagUA.TagForceCreateIn,
		tagUA.TagForceCreateOut,
		TagForceCreateOut,
	](
		goChiTooling.ProcessRequestBody,
		s.mapper.TagForceCreateInFromChi,
		s.forceCreateUA,
		s.mapper.TagForceCreateOutToChi,
	)

	return chioas.Path{
		Paths: chioas.Paths{
			"/force-create": {
				Methods: chioas.Methods{
					http.MethodPost: {
						Description: "Эндпоинт для создания Tag",
						Handler:     forceCreateHandler,
						Request:     &chioas.Request{Schema: TagForceCreateInOpenApiDefinition},
						Responses: chioas.Responses{
							http.StatusCreated: {Schema: TagForceCreateOutOpenApiDefinition},
						},
					},
				},
			},
		},
	}
}
