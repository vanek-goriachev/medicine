package tag

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
)

type tagChiMapper interface {
	FromChi(tag tagChi.Tag) (tagModels.Tag, error)
	MultipleFromChi(tag []tagChi.Tag) ([]tagModels.Tag, error)
	ToChi(tag tagModels.Tag) tagChi.Tag
	MultipleToChi(tags []tagModels.Tag) []tagChi.Tag
}
