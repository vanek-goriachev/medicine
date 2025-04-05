package tags_space

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper   userActionsMapper
	createUA createTagsSpaceUserAction
}

func NewService(
	mapper userActionsMapper,
	createUA createTagsSpaceUserAction,
) *Service {
	return &Service{
		mapper:   mapper,
		createUA: createUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	createHandler := goChiTooling.Handler[
		CreateTagsSpaceIn,
		tagsSpaceUA.CreateTagsSpaceIn,
		tagsSpaceUA.CreateTagsSpaceOut,
		CreateTagsSpaceOut,
	](s.mapper.CreateTagsSpaceInFromChi, s.createUA, s.mapper.CreateTagsSpaceOutToChi)

	return chioas.Path{
		Methods: chioas.Methods{
			http.MethodPost: {
				Description: "Эндпоинт для создания TagsSpace",
				Handler:     createHandler,
				Request:     &chioas.Request{Schema: CreateTagsSpaceInOpenApiDefinition},
				Responses: chioas.Responses{
					http.StatusCreated: {Schema: CreateTagsSpaceOutOpenApiDefinition},
				},
			},
		},
	}
}
