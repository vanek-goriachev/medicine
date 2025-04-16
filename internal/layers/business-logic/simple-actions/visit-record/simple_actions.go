package visit_record

type SimpleActions struct {
	idGenerator        EntityIDGenerator
	visitRecordFactory VisitRecordFactory

	fileAtomicActions MedicalFileAtomicActions
	atomicActions     AtomicActions
}

func NewSimpleActions(
	idGenerator EntityIDGenerator,
	visitRecordFactory VisitRecordFactory,
	fileAtomicActions MedicalFileAtomicActions,
	atomicActions AtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:        idGenerator,
		visitRecordFactory: visitRecordFactory,
		fileAtomicActions:  fileAtomicActions,
		atomicActions:      atomicActions,
	}
}
