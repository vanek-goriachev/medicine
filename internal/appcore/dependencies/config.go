package dependencies

import (
	"medicine/internal/appcore/dependencies/db"
	"medicine/pkg/telemetry"
)

type IAMConfig struct{}

type DepsConfig struct {
	IAM       IAMConfig        `yaml:"iam"`
	Telemetry telemetry.Config `yaml:"telemetry"`
	DB        db.Config        `yaml:"db"`
}
