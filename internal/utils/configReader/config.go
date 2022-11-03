package configReader

import (
	"github.com/BurntSushi/toml"
)

const (
	DEFAULT_SERVER_ADDRESS       = "127.0.0.1:8080"
	DEFAULT_ORIGINS_ADDRESS_CORS = "http://127.0.0.1:8080"
	DEFAULT_LOG_LEVEL            = "debug"
	DEFAULT_LOG_FILE_PATH        = "logs/serverRestApi/logs.log"
)

type Config struct {
	ServerAddress             string   `toml:"serverAddress"`
	AllowOriginsAddressesCORS []string `toml:"originsAddressesCORS"`
	LogLevel                  string   `toml:"logLevel"`
	LogFilePath               string   `toml:"logFilePath"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{
		ServerAddress:             DEFAULT_SERVER_ADDRESS,
		AllowOriginsAddressesCORS: []string{DEFAULT_ORIGINS_ADDRESS_CORS},
		LogLevel:                  DEFAULT_LOG_LEVEL,
		LogFilePath:               DEFAULT_LOG_FILE_PATH,
	}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
