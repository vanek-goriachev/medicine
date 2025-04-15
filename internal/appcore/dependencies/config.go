package dependencies

import (
	"medicine/internal/appcore/dependencies/db"
	file_storage "medicine/internal/appcore/dependencies/file-storage"
	"medicine/internal/tooling/iam"
	"medicine/pkg/telemetry"
)

type DepsConfig struct {
	IAM         iam.Config          `yaml:"iam"`
	Telemetry   telemetry.Config    `yaml:"telemetry"`
	DB          db.Config           `yaml:"db"`
	FileStorage file_storage.Config `yaml:"file_storage"`
}
