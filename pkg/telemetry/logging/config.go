package logging

type Config struct {
	// Level
	// PrintToStdout

	// Processor
	// Provider
	// UnifiedAgent
}

func (Config) Validate() error {
	return nil
}
