package chi

import (
	"medicine/internal/appcore/collections"
	medicalFileMapper "medicine/internal/layers/adapters/transport/rest/go-chi/medical-file"
	chiTagMapper "medicine/internal/layers/adapters/transport/rest/go-chi/tag"
	chiTagsSpaceMapper "medicine/internal/layers/adapters/transport/rest/go-chi/tags-space"
	visitRecordMapper "medicine/internal/layers/adapters/transport/rest/go-chi/visit-record"
)

type mappers struct {
	tag   *chiTagMapper.ChiMapper
	tagUA *chiTagMapper.UserActionsChiMapper

	tagsSpace   *chiTagsSpaceMapper.ChiMapper
	tagsSpaceUA *chiTagsSpaceMapper.UserActionsChiMapper

	medicalFile *medicalFileMapper.ChiMapper

	visitRecord   *visitRecordMapper.ChiMapper
	visitRecordUA *visitRecordMapper.UserActionsChiMapper
}

func newChiMappers(commonMappers *collections.CommonMappers) *mappers {
	var m mappers

	m.tag = chiTagMapper.NewChiMapper(commonMappers.EntityIDMapper)
	m.tagUA = chiTagMapper.NewUserActionsChiMapper(commonMappers.EntityIDMapper, m.tag)

	m.tagsSpace = chiTagsSpaceMapper.NewChiMapper(commonMappers.EntityIDMapper, m.tag)
	m.tagsSpaceUA = chiTagsSpaceMapper.NewUserActionsChiMapper(commonMappers.EntityIDMapper, m.tagsSpace)

	m.medicalFile = medicalFileMapper.NewChiMapper(commonMappers.EntityIDMapper)

	m.visitRecord = visitRecordMapper.NewChiMapper(
		commonMappers.Datetime,
		commonMappers.EntityIDMapper,
	)
	m.visitRecordUA = visitRecordMapper.NewUserActionsChiMapper(
		commonMappers.Datetime,
		commonMappers.EntityIDMapper,
		m.medicalFile,
		m.visitRecord,
	)

	return &m
}
