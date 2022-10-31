package main

import (
	"2022_2_GoTo_team/server"
)

const serverAddress = "127.0.0.1:8080"

var allowOriginsAddressesCORS = []string{"http://127.0.0.1:8080"}

//var allowOriginsAddressesCORS = []string{"http://95.163.213.142:8081/"}

func main() {
	server.Run(serverAddress, allowOriginsAddressesCORS)
}
