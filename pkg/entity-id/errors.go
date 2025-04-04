package entity_id

import "errors"

type CantParsedEntityIDError struct {
	rawID string
}

func NewCantParsedEntityIDError(rawID string) *CantParsedEntityIDError {
	return &CantParsedEntityIDError{rawID: rawID}
}

func (e *CantParsedEntityIDError) Error() string {
	return "cant parse entity id: " + e.rawID
}

func (e *CantParsedEntityIDError) Is(rawOther error) bool {
	var other *CantParsedEntityIDError

	if !errors.As(rawOther, &other) {
		return false
	}

	return other.rawID == e.rawID
}
