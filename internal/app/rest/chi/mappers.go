package chi

import (
	"medicine/internal/appcore/collections"
	chiTagMapper "medicine/internal/layers/adapters/transport/rest/go-chi/tag"
	chiTagsSpaceMapper "medicine/internal/layers/adapters/transport/rest/go-chi/tags-space"
)

type mappers struct {
	tag         *chiTagMapper.ChiMapper
	tagsSpace   *chiTagsSpaceMapper.ChiMapper
	tagsSpaceUA *chiTagsSpaceMapper.UserActionsChiMapper
}

func newChiMappers(commonMappers *collections.CommonMappers) *mappers {
	var m mappers

	m.tag = chiTagMapper.NewChiMapper(commonMappers.EntityIDMapper)

	m.tagsSpace = chiTagsSpaceMapper.NewChiMapper(commonMappers.EntityIDMapper, m.tag)
	m.tagsSpaceUA = chiTagsSpaceMapper.NewUserActionsChiMapper(commonMappers.EntityIDMapper, m.tagsSpace)

	return &m
}
