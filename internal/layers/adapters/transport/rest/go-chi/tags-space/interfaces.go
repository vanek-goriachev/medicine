package tags_space

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
	tagsSpaceChi "medicine/internal/layers/transport/rest/go-chi/tags-space"
)

type tagChiMapper interface {
	FromChi(tag tagChi.Tag) (tagModels.Tag, error)
	MultipleFromChi(tag []tagChi.Tag) ([]tagModels.Tag, error)
	ToChi(tag tagModels.Tag) tagChi.Tag
	MultipleToChi(tags []tagModels.Tag) []tagChi.Tag
}

// tagsSpaceChiMapper implemented by ChiMapper.
type tagsSpaceChiMapper interface {
	ToChi(tagsSpace tagsSpaceModels.TagsSpace) tagsSpaceChi.TagsSpace
	MultipleToChi(tagsSpaces []tagsSpaceModels.TagsSpace) []tagsSpaceChi.TagsSpace
	FromChi(chiTagsSpace tagsSpaceChi.TagsSpace) (tagsSpaceModels.TagsSpace, error)
	MultipleFromChi(chiTagsSpaces []tagsSpaceChi.TagsSpace) ([]tagsSpaceModels.TagsSpace, error)
}
