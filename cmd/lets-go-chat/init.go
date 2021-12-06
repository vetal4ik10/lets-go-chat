package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"log"
)

func main() {
	// Init database.
	dataSoutceName := configs.GetPostgresUrl()
	db, err := sql.Open("postgres", dataSoutceName)
	if err != nil {
		log.Fatal(err)
	}

	drop := "DROP TABLE IF exists users;"
	create := drop + "CREATE TABLE users (" +
		"uid varchar(60) NOT NULL, " +
		"name varchar(60) NOT NULL, " +
		"pass varchar(255) DEFAULT NULL" +
		");"
	create += "DROP TABLE IF exists token;" +
		"CREATE TABLE token (" +
		"uid varchar(60) NOT NULL, " +
		"secret varchar(60) NOT NULL" +
		")"
	_, err = db.Query(create)

	if err != nil {
		log.Fatal(err)
	}
}
