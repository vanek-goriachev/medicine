package appcore

type Dependencies struct {
	// Telemetry
	// IAM
	// DB
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func (*Dependencies) Initialize() error {
	// Initialize all dependencies
	return nil
}
