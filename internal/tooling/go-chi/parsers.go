package go_chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
)

var _ customIdentifiers.Identifier

//nolint:ireturn // Return type is not an interface
func ProcessRequestBody[dtoInType any](r *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	var zero dtoInType

	err := json.NewDecoder(r.Body).Decode(&inDTO)
	if err != nil {
		return zero, fmt.Errorf("error decoding request body: %w", err)
	}

	return inDTO, nil
}

//nolint:ireturn // Return type is not an interface
func ProcessRequestQueryArgs[dtoInType any](r *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	var zero dtoInType

	err := schema.NewDecoder().Decode(&inDTO, r.URL.Query())
	if err != nil {
		return zero, fmt.Errorf("error decoding request query params: %w", err)
	}

	return inDTO, nil
}

//nolint:ireturn // Return type is not an interface
func NoParser[dtoInType any](_ *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	return inDTO, nil
}
