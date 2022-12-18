package configReader

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerAddress             string   `toml:"serverAddress"`
	AllowOriginsAddressesCORS []string `toml:"originsAddressesCORS"`
	LogLevel                  string   `toml:"logLevel"`
	LogFilePath               string   `toml:"logFilePath"`

	StaticDirAbsolutePath string `toml:"staticDirAbsolutePath"`
	ProfilePhotosDirPath  string `toml:"profilePhotosDirPath"`

	EnableEchoCsrfToken bool `toml:"enableEchoCsrfToken"`
	EnableEchoSecurity  bool `toml:"enableEchoSecurity"`

	EnableHttpsWithTLS        bool   `toml:"enableHttpsWithTLS"`
	TLSCertificateFilePath    string `toml:"TLSCertificateFilePath"`
	TLSCertificateKeyFilePath string `toml:"TLSCertificateKeyFilePath"`

	DatabaseUser               string `toml:"databaseUser"`
	DatabaseName               string `toml:"databaseName"`
	DatabasePassword           string `toml:"databasePassword"`
	DatabaseHost               string `toml:"databaseHost"`
	DatabasePort               string `toml:"databasePort"`
	DatabaseMaxOpenConnections string `toml:"databaseMaxOpenConnections"`

	AuthSessionServiceAddress string `toml:"authSessionServiceAddress"`
	UserProfileServiceAddress string `toml:"userProfileServiceAddress"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
