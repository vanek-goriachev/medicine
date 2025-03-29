package appcore

type Core struct {
	Dependencies *Dependencies
	// Gateways
	// Mappers
	// Actions
	// UserScenarios
}

func NewCore(dependencies *Dependencies) *Core {
	return &Core{
		Dependencies: dependencies,
	}
}

func (*Core) Initialize() {
	// Initialize core
}
