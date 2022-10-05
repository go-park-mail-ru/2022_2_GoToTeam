package server

import (
	"2022_2_GoTo_team/server/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(serverAddress string) {
	r := mux.NewRouter()

	api := api.GetApi(serverAddress)
	r.HandleFunc("/", api.RootHandler)
	r.HandleFunc("/api/v1/user/signup", api.SignupUserHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/session/create", api.CreateSessionHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/feed", api.FeedHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/login", api.LoginHandler)
	r.HandleFunc("/logout", api.LogoutHandler)

	log.Println("Starting server on:", serverAddress)
	http.ListenAndServe(serverAddress, r)
}
