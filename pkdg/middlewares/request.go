package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// RequestLoggingHandler logs all requests.
func RequestLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("Method: %s Uri: %s", r.RemoteAddr, r.Method, r.RequestURI)
		log.Println(logText)
		next.ServeHTTP(w, r)
	})
}

// RequestErrorLoggingHandler logs panics.
func RequestErrorLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Error: ", r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
