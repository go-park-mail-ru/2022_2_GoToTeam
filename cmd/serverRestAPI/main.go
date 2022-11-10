package main

import (
	"2022_2_GoTo_team/internal/serverRestAPI"
	"flag"
)

const DEFAULT_CONFIG_FILE_PATH = "configs/serverRestAPI/server.toml"

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config_file_path", DEFAULT_CONFIG_FILE_PATH, "config file path")
}

func main() {
	flag.Parse()

	serverRestAPI.Run(configFilePath)
}
