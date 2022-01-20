package configs

import (
	"os"
)

// GetPostgresUrl get database connection url.
func GetPostgresUrl() string {
	url := os.Getenv("HEROKU_POSTGRESQL_RED_URL")
	if url == "" {
		url = "user=postgres password=secret dbname=lets_go_chat sslmode=disable"
	}
	return url
}

// GetServerPort get server port for listening.
func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
