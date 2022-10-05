package main

import (
	"2022_2_GoTo_team/server"
)

const serverAddress = "127.0.0.1:8080"

func main() {
	server.Run(serverAddress)
}
