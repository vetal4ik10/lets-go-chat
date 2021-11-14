package env

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/db/postgres"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"log"
)

type Env struct {
	db       *sql.DB
	userRepo models.UserRepo
}

var env Env

func init() {
	// Init database.
	dataSoutceName := configs.GetPostgresUrl()
	db, err := sql.Open("postgres", dataSoutceName)
	if err != nil {
		log.Fatal(err)
	}

	// Create database schema.
	schema(db)

	// Init user repo.
	userRepo := postgres.PostgresUserRepo{db}

	env = Env{db, userRepo}
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

func GetUserRepo() models.UserRepo {
	return env.userRepo
}
