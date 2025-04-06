package go_chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	userModels "medicine/pkg/user"
	userActionPkg "medicine/pkg/user-action"
)

var _ customIdentifiers.Identifier

//nolint:revive // This function is pretty transparent so its length is not a problem
func Handler[dtoInT, domainInT, domainOutT, outT any](
	inputMapper func(dto dtoInT) (domainInT, error),
	userAction userActionPkg.UserAction[domainInT, domainOutT],
	outputMapper func(domainOutT) outT,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close() //nolint:errcheck // Can't really do anything with error here

		jsEnc := json.NewEncoder(w)
		jsDec := json.NewDecoder(r.Body)

		w.Header().Set("Content-Type", "application/json")

		user, err := userModels.GetFromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		in, err := processInput[dtoInT, domainInT](jsDec, inputMapper)
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

//nolint:ireturn // Return type is not an interface
func processInput[dtoInType, domainInType any](
	jsDec *json.Decoder,
	inputMapper func(dto dtoInType) (domainInType, error),
) (domainInType, error) {
	var zeroAnswer domainInType

	var inDTO dtoInType

	err := jsDec.Decode(&inDTO)
	if err != nil {
		return zeroAnswer, fmt.Errorf("error decoding input: %w", err)
	}

	in, err := inputMapper(inDTO)
	if err != nil {
		return zeroAnswer, fmt.Errorf("error mapping input to domain model: %w", err)
	}

	return in, nil
}
