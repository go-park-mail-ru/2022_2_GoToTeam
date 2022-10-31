package main

import (
	"2022_2_GoTo_team/internal/serverRestAPI"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := serverRestAPI.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	fmt.Println(config)
	if err != nil {
		log.Fatal(err)
	}
	serverRestAPI.Run(config)
}
