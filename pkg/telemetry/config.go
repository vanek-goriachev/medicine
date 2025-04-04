package telemetry

import "medicine/pkg/telemetry/logging"

type Config struct {
	Logging logging.Config `yaml:"logging"`
	// Tracing
	// Metering
}
