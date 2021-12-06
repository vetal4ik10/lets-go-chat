package handlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/internal/chat"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"log"
	"net/http"
)

type chatHandlers struct {
	tM onetimetoken.TokenManager
	cS chat.ChatServer
}

func NewChatHandlers(tm onetimetoken.TokenManager, ch chat.ChatServer) *chatHandlers {
	return &chatHandlers{tm, ch}
}

func (cH *chatHandlers) ChatStart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "home.html")
}

func (cH *chatHandlers) ChatActiveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": cH.cS.GetActiveUsers()})
}

func (cH *chatHandlers) ChatConnect(w http.ResponseWriter, r *http.Request) {
	secret := r.URL.Query().Get("token")
	if secret == "" {
		ErrorResponse(w, "Token is required", http.StatusBadRequest)
		return
	}

	token, err := cH.tM.InitToken(secret)
	if err != nil || !cH.tM.Verified(token) {
		ErrorResponse(w, "Token is not valid", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := chat.NewChatClient(conn, token)
	cH.cS.ClientConnect(c)
	go c.Reader(cH.cS)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
