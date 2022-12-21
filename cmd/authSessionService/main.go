package main

import (
	"2022_2_GoTo_team/internal/authSessionService"
	"flag"
)

const _DEFAULT_CONFIG_FILE_PATH = "configs/authSessionService/server.toml"

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config_file_path__", _DEFAULT_CONFIG_FILE_PATH, "config file path")
}

func main() {
	flag.Parse()

	authSessionService.Run(configFilePath)
}
