package visit_record

import (
	"time"

	entityID "medicine/pkg/entity-id"
)

type VisitRecord struct {
	Datetime time.Time
	Name     string
	ID       entityID.EntityID
}

type VisitRecordLinkedEntities struct {
	TagIDs         []entityID.EntityID
	MedicalFileIDs []entityID.EntityID
}
