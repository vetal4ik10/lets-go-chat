package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/internal/chat"
	"github.com/vetal4ik10/lets-go-chat/internal/chat_message"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/middlewares"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"log"
	"net/http"
)

func initDatabase() *sql.DB {
	// Init database.
	dataSoutceName := configs.GetPostgresUrl()
	db, err := sql.Open("postgres", dataSoutceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	// Dependencies.
	db := initDatabase()                                    // Database
	userRepo := reposetories.NewUserRepo(db)                // User repository for working with user.
	tM := onetimetoken.NewTokenManager(db, userRepo)        // One time token manager for working with token.
	cMM := chat_message.NewChatMessageManager(db, userRepo) // Chat message manager.
	cS := chat.NewChatServer(cMM)                           // Chat server.

	r := mux.NewRouter()

	// Init user handlers.
	userH := handlers.NewUserHandlers(userRepo)
	r.Use(middlewares.RequestLoggingHandler)
	r.Use(middlewares.RequestErrorLoggingHandler)
	r.HandleFunc("/user", userH.UserCreate).Methods(http.MethodPost)
	r.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		userH.Login(tM, w, r)
	}).Methods(http.MethodPost)

	// Init chat handlers.
	cH := handlers.NewChatHandlers(tM, cS)
	go cS.Run()
	r.HandleFunc("/chat/ws.rtm.start", cH.ChatStart).Methods(http.MethodGet)
	r.HandleFunc("/user/active", cH.ChatActiveUsers).Methods(http.MethodGet)
	r.HandleFunc("/ws", cH.ChatConnect).Queries("token", "{token}").Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))
}
