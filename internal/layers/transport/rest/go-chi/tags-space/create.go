package tags_space

import (
	"encoding/json"
	"net/http"

	userModels "medicine/pkg/user"
)

type CreateTagsSpaceIn struct {
	Name string `json:"name"`
}

type CreateTagsSpaceOut struct {
	TagsSpace TagsSpace `json:"tags_space"`
}

func (c *Service) Create(w http.ResponseWriter, r *http.Request) {
	jsEnc := json.NewEncoder(w)
	jsDec := json.NewDecoder(r.Body)
	ctx := r.Context()

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	user := userModels.FromContext(ctx)
	if user == (userModels.User{}) { //nolint:exhaustruct // Checking if struct is empty
		http.Error(w, "No user in context", http.StatusUnauthorized)

		return
	}

	var inDTO CreateTagsSpaceIn

	err := jsDec.Decode(&inDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	in := c.mapper.CreateTagsSpaceInFromChi(inDTO)

	out, err := c.createUA.Act(ctx, user, in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	outDTO := c.mapper.CreateTagsSpaceOutToChi(out)

	err = jsEnc.Encode(outDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
