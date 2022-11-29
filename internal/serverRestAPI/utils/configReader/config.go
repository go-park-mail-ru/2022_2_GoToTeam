package configReader

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerAddress             string   `toml:"serverAddress"`
	AllowOriginsAddressesCORS []string `toml:"originsAddressesCORS"`
	LogLevel                  string   `toml:"logLevel"`
	LogFilePath               string   `toml:"logFilePath"`

	DatabaseUser               string `toml:"databaseUser"`
	DatabaseName               string `toml:"databaseName"`
	DatabasePassword           string `toml:"databasePassword"`
	DatabaseHost               string `toml:"databaseHost"`
	DatabasePort               string `toml:"databasePort"`
	DatabaseMaxOpenConnections string `toml:"databaseMaxOpenConnections"`

	AuthSessionServiceAddress string `toml:"authSessionServiceAddress"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
