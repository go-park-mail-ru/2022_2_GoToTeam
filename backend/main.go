package main

import (
	"2022_2_GoTo_team/server"
)

//const serverAddress = "127.0.0.1:8080"

const serverAddress = "95.163.213.142:3004"

func main() {
	server.Run(serverAddress)
}
