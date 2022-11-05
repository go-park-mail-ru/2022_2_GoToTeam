package configReader

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerAddress             string   `toml:"serverAddress"`
	AllowOriginsAddressesCORS []string `toml:"originsAddressesCORS"`
	LogLevel                  string   `toml:"logLevel"`
	LogFilePath               string   `toml:"logFilePath"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{
		ServerAddress:             domain.DEFAULT_SERVER_ADDRESS,
		AllowOriginsAddressesCORS: []string{domain.DEFAULT_ORIGINS_ADDRESS_CORS},
		LogLevel:                  domain.DEFAULT_LOG_LEVEL,
		LogFilePath:               domain.DEFAULT_LOG_FILE_PATH,
	}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
