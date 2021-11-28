package middlewares

import (
	"fmt"
	"log"
	"net/http"
)


func RequestLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("Method: %s Uri: %s", r.RemoteAddr, r.Method, r.RequestURI)
		log.Println(logText)
		next.ServeHTTP(w, r)


	})
}

func RequestErrorLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

