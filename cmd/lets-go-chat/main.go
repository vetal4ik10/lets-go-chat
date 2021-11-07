package main

import (
	"github.com/vetal4ik10/lets-go-chat/internal/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user", endpoints.UserCreate)
	http.HandleFunc("/user/list", endpoints.UserList)
	http.HandleFunc("/user/login", endpoints.Login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}