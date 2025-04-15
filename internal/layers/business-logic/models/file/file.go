package file

import (
	entityID "medicine/pkg/entity-id"
)

type FileDataType *[]byte

type File struct {
	ID   entityID.EntityID
	Data FileDataType
}
