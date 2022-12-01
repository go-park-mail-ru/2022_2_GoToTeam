package configReader

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerAddress           string `toml:"serverAddress"`
	PrometheusServerAddress string `toml:"prometheusServerAddress"`
	LogLevel                string `toml:"logLevel"`
	LogFilePath             string `toml:"logFilePath"`

	DatabaseUser               string `toml:"databaseUser"`
	DatabaseName               string `toml:"databaseName"`
	DatabasePassword           string `toml:"databasePassword"`
	DatabaseHost               string `toml:"databaseHost"`
	DatabasePort               string `toml:"databasePort"`
	DatabaseMaxOpenConnections string `toml:"databaseMaxOpenConnections"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
