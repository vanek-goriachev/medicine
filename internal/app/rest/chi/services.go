package chi

import (
	"medicine/internal/appcore/collections"
	chiTag "medicine/internal/layers/transport/rest/go-chi/tag"
	chiTagsSpace "medicine/internal/layers/transport/rest/go-chi/tags-space"
)

type chiServices struct {
	tag       *chiTag.Service
	tagsSpace *chiTagsSpace.Service
}

func newChiServices(chiMappers *mappers, userActions *collections.UserActions) *chiServices {
	var s chiServices

	s.tag = chiTag.NewService(
		chiMappers.tagUA,
		//userActions.Tag.GetByID,
		userActions.Tag.ForceCreate,
	)

	s.tagsSpace = chiTagsSpace.NewService(
		chiMappers.tagsSpaceUA,
		userActions.TagsSpace.GetByID,
		userActions.TagsSpace.ListByUser,
		userActions.TagsSpace.Create,
	)

	return &s
}
