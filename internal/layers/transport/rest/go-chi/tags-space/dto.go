package tags_space

import tagChi "medicine/internal/layers/transport/rest/go-chi/tag"

type TagsSpace struct {
	ID     string       `json:"id"`
	UserID string       `json:"user_id"`
	Name   string       `json:"name"`
	Tags   []tagChi.Tag `json:"tags"`
}
