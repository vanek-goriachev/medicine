package db

import "fmt"

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SslMode  string `yaml:"ssl_mode"`
	Timezone string `yaml:"timezone"`
	Port     int    `yaml:"port"`
}

func (dbc *Config) AsDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbc.Host,
		dbc.Port,
		dbc.User,
		dbc.Password,
		dbc.DBName,
		dbc.SslMode,
		dbc.Timezone,
	)
}
