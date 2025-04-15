package tags_space

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper             userActionsMapper
	getByIDUA          tagsSpaceGetByIDUserAction
	listAllAvailableUA tagsSpaceListAllAvailableUserAction
	createUA           tagsSpaceCreateUserAction
	deleteUA           tagsSpaceDeleteUserAction
}

func NewService(
	mapper userActionsMapper,
	getByIDUA tagsSpaceGetByIDUserAction,
	listByUserUA tagsSpaceListAllAvailableUserAction,
	createUA tagsSpaceCreateUserAction,
	deleteUA tagsSpaceDeleteUserAction,
) *Service {
	return &Service{
		mapper:             mapper,
		getByIDUA:          getByIDUA,
		listAllAvailableUA: listByUserUA,
		createUA:           createUA,
		deleteUA:           deleteUA,
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

	listAllAvailableHandler := goChiTooling.Handler[
		TagsSpaceListAllAvailableIn,
		tagsSpaceUA.TagsSpaceListAllAvailableIn,
		tagsSpaceUA.TagsSpaceListAllAvailableOut,
		TagsSpaceListAllAvailableOut,
	](
		goChiTooling.NoParser,
		s.mapper.TagsSpaceListAllAvailableInFromChi,
		s.listAllAvailableUA,
		s.mapper.TagsSpaceListAllAvailableOutToChi,
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

	deleteHandler := goChiTooling.Handler[
		TagsSpaceDeleteIn,
		tagsSpaceUA.TagsSpaceDeleteIn,
		tagsSpaceUA.TagsSpaceDeleteOut,
		TagsSpaceDeleteOut,
	](
		goChiTooling.ProcessRequestBody,
		s.mapper.TagsSpaceDeleteInFromChi,
		s.deleteUA,
		s.mapper.TagsSpaceDeleteOutToChi,
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
			"/delete": {
				Methods: chioas.Methods{
					http.MethodDelete: {
						Description: "Эндпоинт для удаления TagsSpace",
						Handler:     deleteHandler,
						Request:     &chioas.Request{Schema: TagsSpaceDeleteInOpenApiDefinition},
						Responses: chioas.Responses{
							http.StatusOK: {Schema: TagsSpaceDeleteOutOpenApiDefinition},
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
			"/list-all-available": {
				Methods: chioas.Methods{
					http.MethodGet: {
						Description: "Эндпоинт для получения списка TagsSpace текущего пользователя",
						Handler:     listAllAvailableHandler,
						Responses: chioas.Responses{
							http.StatusOK: {Schema: TagsSpaceListAllAvailableOutOpenApiDefinition},
						},
					},
				},
			},
		},
	}
}
