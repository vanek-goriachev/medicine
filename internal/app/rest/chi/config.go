package chi

import "time"

type Config struct {
	Port        int           `yaml:"port"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}
