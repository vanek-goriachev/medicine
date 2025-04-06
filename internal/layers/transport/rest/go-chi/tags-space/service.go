package tags_space

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper    userActionsMapper
	createUA  createTagsSpaceUserAction
	getByIDUA tagsSpaceGetByIDUserAction
}

func NewService(
	mapper userActionsMapper,
	createUA createTagsSpaceUserAction,
	getByIDUA tagsSpaceGetByIDUserAction,
) *Service {
	return &Service{
		mapper:    mapper,
		createUA:  createUA,
		getByIDUA: getByIDUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	getByIDHandler := goChiTooling.Handler[
		TagsSpaceGetByIDIn,
		tagsSpaceUA.TagsSpaceGetByIDIn,
		tagsSpaceUA.TagsSpaceGetByIDOut,
		TagsSpaceGetByIDOut,
	](s.mapper.TagsSpaceGetByIDInFromChi, s.getByIDUA, s.mapper.TagsSpaceGetByIDOutToChi)

	createHandler := goChiTooling.Handler[
		TagsSpaceCreateIn,
		tagsSpaceUA.TagsSpaceCreateIn,
		tagsSpaceUA.TagsSpaceCreateOut,
		TagsSpaceCreateOut,
	](s.mapper.TagsSpaceCreateInFromChi, s.createUA, s.mapper.TagsSpaceCreateOutToChi)

	return chioas.Path{
		Methods: chioas.Methods{
			http.MethodGet: {
				Description: "Эндпоинт для получения TagsSpace",
				Handler:     getByIDHandler,
				Request:     &chioas.Request{Schema: TagsSpaceGetByIDInOpenApiDefinition},
				Responses: chioas.Responses{
					http.StatusOK: {Schema: TagsSpaceGetByIDOutOpenApiDefinition},
				},
			},
			http.MethodPost: {
				Description: "Эндпоинт для создания TagsSpace",
				Handler:     createHandler,
				Request:     &chioas.Request{Schema: TagsSpaceCreateInOpenApiDefinition},
				Responses: chioas.Responses{
					http.StatusCreated: {Schema: TagsSpaceCreateOutOpenApiDefinition},
				},
			},
		},
	}
}
