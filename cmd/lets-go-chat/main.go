package main

import (
	"github.com/gorilla/mux"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	//.Headers("Content-Type", "application/json")

	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/user/login", handlers.Login).Methods("POST")
	r.HandleFunc("/user/list", handlers.UserList).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))
}
