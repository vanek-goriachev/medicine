package chi

import (
	"medicine/internal/appcore/collections"
	chiTag "medicine/internal/layers/transport/rest/go-chi/tag"
	chiTagsSpace "medicine/internal/layers/transport/rest/go-chi/tags-space"
	chiVisitRecord "medicine/internal/layers/transport/rest/go-chi/visit-record"
)

type chiServices struct {
	tag       *chiTag.Service
	tagsSpace *chiTagsSpace.Service

	visitRecord *chiVisitRecord.Service
}

func newChiServices(chiMappers *mappers, userActions *collections.UserActions) *chiServices {
	var s chiServices

	s.tag = chiTag.NewService(
		chiMappers.tagUA,
		userActions.Tag.ForceCreate,
		userActions.Tag.UntagAllAndDelete,
	)

	s.tagsSpace = chiTagsSpace.NewService(
		chiMappers.tagsSpaceUA,
		userActions.TagsSpace.GetByID,
		userActions.TagsSpace.ListAllAvailable,
		userActions.TagsSpace.Create,
		userActions.TagsSpace.Delete,
	)

	s.visitRecord = chiVisitRecord.NewService(
		chiMappers.visitRecordUA,
		userActions.VisitRecord.Create,
	)

	return &s
}
