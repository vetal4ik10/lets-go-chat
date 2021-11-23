package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"log"
	"net/http"
)

func initDatabase() (*sql.DB) {
	// Init database.
	dataSoutceName := configs.GetPostgresUrl()
	db, err := sql.Open("postgres", dataSoutceName)
	if err != nil {
		log.Fatal(err)
	}
	schema(db)

	return db
}

func schema(db *sql.DB) {
	drop := "DROP TABLE IF exists users;"
	create := drop + "CREATE TABLE users (" +
		"uid varchar(60) NOT NULL, " +
		"name varchar(60) NOT NULL, " +
		"pass varchar(255) DEFAULT NULL" +
		")"
	_, err := db.Query(create)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := initDatabase()
	userRepo := reposetories.NewUserRepo(db)
	userH := handlers.NewUserHandlers(userRepo)

	r := mux.NewRouter()
	r.HandleFunc("/user", userH.UserCreate).Methods("POST")
	r.HandleFunc("/user/login", userH.Login).Methods("POST")
	r.HandleFunc("/user/list", userH.UserList).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))
}
