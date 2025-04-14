package tag

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper              userActionsMapper
	forceCreateUA       tagCreateUserAction
	untagAllAndDeleteUA tagUntagAllAndDeleteUserAction
}

func NewService(
	mapper userActionsMapper,
	forceCreateUA tagCreateUserAction,
	untagAllAndDeleteUA tagUntagAllAndDeleteUserAction,
) *Service {
	return &Service{
		mapper:              mapper,
		forceCreateUA:       forceCreateUA,
		untagAllAndDeleteUA: untagAllAndDeleteUA,
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

	untagAllAndDeleteHandler := goChiTooling.Handler[
		TagUntagAllAndDeleteIn,
		tagUA.TagUntagAllAndDeleteIn,
		tagUA.TagUntagAllAndDeleteOut,
		TagUntagAllAndDeleteOut,
	](
		goChiTooling.ProcessRequestBody,
		s.mapper.TagUntagAllAndDeleteInFromChi,
		s.untagAllAndDeleteUA,
		s.mapper.TagUntagAllAndDeleteOutToChi,
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
			"/untag-all-and-delete": {
				Methods: chioas.Methods{
					http.MethodDelete: {
						Description: "Эндпоинт для удаления Tag с предварительным полным удалением " +
							"всех связей с другими сущностями",
						Handler: untagAllAndDeleteHandler,
						Request: &chioas.Request{Schema: TagUntagAllAndDeleteInOpenApiDefinition},
						Responses: chioas.Responses{
							http.StatusNoContent: {Schema: TagUntagAllAndDeleteOutOpenApiDefinition},
						},
					},
				},
			},
		},
	}
}
