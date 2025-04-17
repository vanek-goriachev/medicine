package noop_authorizer

import "medicine/internal/layers/business-logic/authorization"

type InvalidResourceError struct {
	resource authorization.Resource
}

func NewInvalidResourceError(resource authorization.Resource) *InvalidResourceError {
	return &InvalidResourceError{resource: resource}
}

func (e *InvalidResourceError) Error() string {
	return "invalid resource: " + string(e.resource)
}

func (e *InvalidResourceError) Unwrap() error {
	return e
}
