package go_chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	userActionPkg "medicine/pkg/user-action"
)

var _ customIdentifiers.Identifier

func Handler[dtoInT, domainInT, domainOutT, outT any](
	inputMapper func(dto dtoInT) (domainInT, error),
	userAction userActionPkg.UserAction[domainInT, domainOutT],
	outputMapper func(domainOutT) outT,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsEnc := json.NewEncoder(w)
		jsDec := json.NewDecoder(r.Body)

		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")

		in, err := processInput[dtoInT, domainInT](jsDec, inputMapper)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		out, err := userAction.Act(r.Context(), in)
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

func processInput[dtoInType, domainInType any]( //nolint:ireturn // Return type is not an interface
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
