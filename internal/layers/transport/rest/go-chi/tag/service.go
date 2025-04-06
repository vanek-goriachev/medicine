package tag

import (
	"net/http"

	"github.com/go-andiamo/chioas"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	goChiTooling "medicine/internal/tooling/go-chi"
)

type Service struct {
	mapper userActionsMapper
	//getByIDUA    tagGetByIDUserAction
	forceCreateUA tagCreateUserAction
}

func NewService(
	mapper userActionsMapper,
	//getByIDUA tagGetByIDUserAction,
	forceCreateUA tagCreateUserAction,
) *Service {
	return &Service{
		mapper: mapper,
		//getByIDUA:    getByIDUA,
		forceCreateUA: forceCreateUA,
	}
}

func (s *Service) GenerateOpenApiDefinition() chioas.Path {
	//getByIDHandler := goChiTooling.Handler[
	//	TagGetByIDIn,
	//	tagUA.TagGetByIDIn,
	//	tagUA.TagGetByIDOut,
	//	TagGetByIDOut,
	//](
	//	goChiTooling.ProcessRequestQueryArgs,
	//	s.mapper.TagGetByIDInFromChi,
	//	s.getByIDUA,
	//	s.mapper.TagGetByIDOutToChi,
	//)

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
			//"/get-by-id": {
			//	Methods: chioas.Methods{
			//		http.MethodGet: {
			//			Description: "Эндпоинт для получения Tag",
			//			Handler:     getByIDHandler,
			//			QueryParams: TagGetByIDInOpenApiDefinition,
			//			Responses: chioas.Responses{
			//				http.StatusOK: {Schema: TagGetByIDOutOpenApiDefinition},
			//			},
			//		},
			//	},
			//},
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
