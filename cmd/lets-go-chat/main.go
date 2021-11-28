package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/middlewares"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
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
		");"
	create += "DROP TABLE IF exists token;" +
		"CREATE TABLE token (" +
		"uid varchar(60) NOT NULL, " +
		"secret varchar(60) NOT NULL" +
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
	r.Use(middlewares.RequestLoggingHandler)
	r.Use(middlewares.RequestErrorLoggingHandler)
	r.HandleFunc("/user", userH.UserCreate).Methods("POST")
	r.HandleFunc("/user/list", userH.UserList).Methods("GET")


	r.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		tm := onetimetoken.NewTokenManager(db, userRepo)
		userH.Login(tm, w, r)
	}).Methods("POST")




	//r.HandleFunc("/chat/ws.rtm.start", func(w http.ResponseWriter, r *http.Request) {
	//	tM := onetimetoken.NewTokenManager(db, userRepo)
	//
	//	s := r.FormValue("token")
	//	if s == "" {
	//		http.Error(w, "Token is required", http.StatusBadRequest)
	//		return
	//	}
	//
	//	t := tM.InitToken(s)
	//	if
	//
	//	fmt.Println(s)
	//}).Methods("GET")
	//
	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))
}

