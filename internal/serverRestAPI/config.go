package serverRestAPI

//var allowOriginsAddressesCORS = []string{"http://95.163.213.142:8081/"}

type Config struct {
	BindServerAddress             string   `toml:"serverAddress"`
	BindAllowOriginsAddressesCORS []string `toml:"originsAddressesCORS"`
}

func NewConfig() *Config {
	return &Config{
		BindServerAddress:             "127.0.0.1:8080",
		BindAllowOriginsAddressesCORS: []string{"http://127.0.0.1:8080"},
	}
}
