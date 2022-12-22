package main

import (
	"2022_2_GoTo_team/internal/userProfileService"
	"flag"
)

const _DEFAULT_CONFIG_FILE_PATH = "configs/userProfileService/server.toml"

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "userProfileService_config_file_path", _DEFAULT_CONFIG_FILE_PATH, "config file path")
}

func main() {
	flag.Parse()

	userProfileService.Run(configFilePath)
}
