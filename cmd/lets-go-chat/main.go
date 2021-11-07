package main

import (
	"github.com/vetal4ik10/lets-go-chat/internal/endpoints"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}


	http.HandleFunc("/user", endpoints.UserCreate)
	http.HandleFunc("/user/list", endpoints.UserList)
	http.HandleFunc("/user/login", endpoints.Login)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}