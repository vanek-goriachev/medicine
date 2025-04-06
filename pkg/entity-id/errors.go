package entity_id

import "errors"

type InvalidEntityIDError struct {
	rawID string
}

func NewInvalidEntityIDError(rawID string) *InvalidEntityIDError {
	return &InvalidEntityIDError{rawID: rawID}
}

func (e *InvalidEntityIDError) Error() string {
	return "invalid entity id: " + e.rawID
}

func (e *InvalidEntityIDError) Is(rawOther error) bool {
	var other *InvalidEntityIDError

	if !errors.As(rawOther, &other) {
		return false
	}

	return other.rawID == e.rawID
}
