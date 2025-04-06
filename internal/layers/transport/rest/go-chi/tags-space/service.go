package tags_space

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper    userActionsMapper
	getByIDUA tagsSpaceGetByIDUserAction
	createUA  createTagsSpaceUserAction
}

func NewService(
	mapper userActionsMapper,
	getByIDUA tagsSpaceGetByIDUserAction,
	createUA createTagsSpaceUserAction,
) *Service {
	return &Service{
		mapper:    mapper,
		getByIDUA: getByIDUA,
		createUA:  createUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	getByIDHandler := goChiTooling.HandlerWithBody[
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

	createHandler := goChiTooling.HandlerWithBody[
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
		Methods: chioas.Methods{
			http.MethodGet: {
				Description: "Эндпоинт для получения TagsSpace",
				Handler:     getByIDHandler,
				QueryParams: TagsSpaceGetByIDInOpenApiDefinition,
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
