package go_chi

import (
	"encoding/json"
	"net/http"

	userModels "medicine/pkg/user"
	userActionPkg "medicine/pkg/user-action"
)

//nolint:revive // This function is pretty transparent so its length is not a problem
func HandlerWithBody[dtoInT, domainInT, domainOutT, outT any](
	parse func(r *http.Request) (dtoInT, error),
	inputMapper func(dto dtoInT) (domainInT, error),
	userAction userActionPkg.UserAction[domainInT, domainOutT],
	outputMapper func(domainOutT) outT,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close() //nolint:errcheck // Can't really do anything with error here

		jsEnc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")

		user, err := userModels.GetFromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		inDTO, err := parse(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		in, err := inputMapper(inDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		out, err := userAction.Act(r.Context(), user, in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = jsEnc.Encode(outputMapper(out))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
