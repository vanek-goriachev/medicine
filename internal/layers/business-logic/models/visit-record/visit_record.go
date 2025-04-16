package visit_record

import (
	entityID "medicine/pkg/entity-id"
	"time"
)

type VisitRecord struct {
	ID       entityID.EntityID
	Name     string
	Datetime time.Time
}

type VisitRecordLinkedEntities struct {
	TagIDs         []entityID.EntityID
	MedicalFileIDs []entityID.EntityID
}
