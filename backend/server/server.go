package server

import (
	"2022_2_GoTo_team/server/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(serverAddress string) {
	r := mux.NewRouter()

	api := api.GetApi()

	r.HandleFunc("/api/v1/user/signup", api.SignupUserHandler).Methods("POST")
	r.HandleFunc("/api/v1/session/create", api.CreateSessionHandler).Methods("POST")
	r.HandleFunc("/api/v1/feed", api.FeedHandler).Methods("GET")

	log.Println("Starting server on:", serverAddress)
	http.ListenAndServe(serverAddress, r)
}
