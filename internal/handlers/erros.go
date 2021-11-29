package handlers

import (
	"log"
	"net/http"
)

// ErrorResponse return http error and log error in system.
func ErrorResponse(w http.ResponseWriter, error string, code int) {
	log.Println(error)
	http.Error(w, error, code)
}
