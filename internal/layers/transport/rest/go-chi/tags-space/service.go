package tags_space

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper       userActionsMapper
	getByIDUA    tagsSpaceGetByIDUserAction
	listByUserUA tagsSpaceListByUserUserAction
	createUA     tagsSpaceCreateUserAction
}

func NewService(
	mapper userActionsMapper,
	getByIDUA tagsSpaceGetByIDUserAction,
	listByUserUA tagsSpaceListByUserUserAction,
	createUA tagsSpaceCreateUserAction,
) *Service {
	return &Service{
		mapper:       mapper,
		getByIDUA:    getByIDUA,
		listByUserUA: listByUserUA,
		createUA:     createUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	getByIDHandler := goChiTooling.Handler[
		TagsSpaceGetByIDIn,
		tagsSpaceUA.TagsSpaceGetByIDIn,
		tagsSpaceUA.TagsSpaceGetByIDOut,
		TagsSpaceGetByIDOut,
	](
		goChiTooling.ProcessRequestQueryArgs,
		s.mapper.TagsSpaceGetByIDInFromChi,
		s.getByIDUA,
		s.mapper.TagsSpaceGetByIDOutToChi,
	)

	listByUserHandler := goChiTooling.Handler[
		TagsSpaceListByUserIn,
		tagsSpaceUA.TagsSpaceListByUserIn,
		tagsSpaceUA.TagsSpaceListByUserOut,
		TagsSpaceListByUserOut,
	](
		goChiTooling.NoParser,
		s.mapper.TagsSpaceListByUserInFromChi,
		s.listByUserUA,
		s.mapper.TagsSpaceListByUserOutToChi,
	)

	createHandler := goChiTooling.Handler[
		TagsSpaceCreateIn,
		tagsSpaceUA.TagsSpaceCreateIn,
		tagsSpaceUA.TagsSpaceCreateOut,
		TagsSpaceCreateOut,
	](
		goChiTooling.ProcessRequestBody,
		s.mapper.TagsSpaceCreateInFromChi,
		s.createUA,
		s.mapper.TagsSpaceCreateOutToChi,
	)

	return chioas.Path{
		Paths: chioas.Paths{
			"/get-by-id": {
				Methods: chioas.Methods{
					http.MethodGet: {
						Description: "Эндпоинт для получения TagsSpace",
						Handler:     getByIDHandler,
						QueryParams: TagsSpaceGetByIDInOpenApiDefinition,
						Responses: chioas.Responses{
							http.StatusOK: {Schema: TagsSpaceGetByIDOutOpenApiDefinition},
						},
					},
				},
			},
			"/create": {
				Methods: chioas.Methods{
					http.MethodPost: {
						Description: "Эндпоинт для создания TagsSpace",
						Handler:     createHandler,
						Request:     &chioas.Request{Schema: TagsSpaceCreateInOpenApiDefinition},
						Responses: chioas.Responses{
							http.StatusCreated: {Schema: TagsSpaceCreateOutOpenApiDefinition},
						},
					},
				},
			},
			"/list-by-user": {
				Methods: chioas.Methods{
					http.MethodGet: {
						Description: "Эндпоинт для получения списка TagsSpace текущего пользователя",
						Handler:     listByUserHandler,
						Responses: chioas.Responses{
							http.StatusOK: {Schema: TagsSpaceListByUserOutOpenApiDefinition},
						},
					},
				},
			},
		},
	}
}
