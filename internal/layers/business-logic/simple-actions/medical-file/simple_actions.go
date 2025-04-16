package medical_file

type SimpleActions struct {
	idGenerator EntityIDGenerator

	medicalFileFactory MedicalFileFactory

	atomicActions AtomicActions
}

func NewSimpleActions(
	idGenerator EntityIDGenerator,
	medicalFileFactory MedicalFileFactory,
	atomicActions AtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:        idGenerator,
		medicalFileFactory: medicalFileFactory,
		atomicActions:      atomicActions,
	}
}
