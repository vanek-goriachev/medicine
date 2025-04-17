package go_chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
)

const GB = 1024 * 1024 * 1024

var _ customIdentifiers.Identifier

func ProcessRequestBody[dtoInType any](r *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	var zero dtoInType

	err := json.NewDecoder(r.Body).Decode(&inDTO)
	if err != nil {
		return zero, fmt.Errorf("error decoding request body: %w", err)
	}

	return inDTO, nil
}

func ProcessRequestQueryArgs[dtoInType any](r *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	var zero dtoInType

	err := schema.NewDecoder().Decode(&inDTO, r.URL.Query())
	if err != nil {
		return zero, fmt.Errorf("error decoding request query params: %w", err)
	}

	return inDTO, nil
}

func ProcessRequestMultipartFormData[dtoInType any](r *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	var zero dtoInType

	err := r.ParseMultipartForm(1 * GB)
	if err != nil {
		return zero, fmt.Errorf("error parsing multipart form data: %w", err)
	}

	err = schema.NewDecoder().Decode(&inDTO, r.Form)
	if err != nil {
		return zero, fmt.Errorf("error decoding request multipart form data: %w", err)
	}

	return inDTO, nil
}

func NoParser[dtoInType any](_ *http.Request) (dtoInType, error) {
	var inDTO dtoInType
	return inDTO, nil
}
