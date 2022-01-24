package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/pkdg/middlewares"
	"log"
	"net/http"
	"net/http/pprof"
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

// @title           Fancy Golang chat
// @version         1.0
// @description     Just a simple chat service
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	r := mux.NewRouter()
	AttachProfiler(r)

	// Init user handlers.
	userH := InitializeUserHandlers()
	r.Use(middlewares.RequestLoggingHandler)
	r.Use(middlewares.RequestErrorLoggingHandler)
	r.HandleFunc("/user", userH.UserCreate).Methods(http.MethodPost)
	r.HandleFunc("/user/login", userH.Login).Methods(http.MethodPost)

	// Init chat handlers.
	cH := InitializeChatHandlers()
	r.HandleFunc("/chat/ws.rtm.start", cH.ChatStart).Methods(http.MethodGet)
	r.HandleFunc("/user/active", cH.ChatActiveUsers).Methods(http.MethodGet)
	r.HandleFunc("/ws", cH.ChatConnect).Queries("token", "{token}").Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))
}

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}
