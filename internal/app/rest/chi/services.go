package chi

import (
	"medicine/internal/appcore/collections"
	chiTagsSpace "medicine/internal/layers/transport/rest/go-chi/tags-space"
)

type services struct {
	tagsSpace *chiTagsSpace.Service
}

func newChiServices(chiMappers *mappers, userActions *collections.UserActions) *services {
	var s services

	s.tagsSpace = chiTagsSpace.NewService(chiMappers.tagsSpaceUA, userActions.TagsSpace.Create)

	return &s
}
