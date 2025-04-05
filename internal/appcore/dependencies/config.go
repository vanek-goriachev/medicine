package dependencies

import (
	"medicine/internal/appcore/dependencies/db"
	"medicine/internal/tooling/iam"
	"medicine/pkg/telemetry"
)

type DepsConfig struct {
	IAM       iam.Config       `yaml:"iam"`
	Telemetry telemetry.Config `yaml:"telemetry"`
	DB        db.Config        `yaml:"db"`
}
