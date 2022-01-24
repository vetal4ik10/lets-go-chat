package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/internal/chat"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"log"
	"net/http"
)

type ChatHandlers struct {
	tM *onetimetoken.TokenManager
	cS *chat.ChatServer
}

func NewChatHandlers(tm *onetimetoken.TokenManager, ch *chat.ChatServer) *ChatHandlers {
	return &ChatHandlers{tm, ch}
}

// ChatStart
// @Summary  Endpoint to start real time chat
// @Tags     chat
// @Accept   html
// @Produce  html
// @Param    token  query     string  true  "One time token for a loged user"
// @Success  200    {string}  string  "Html page witch chat"
// @Router   /chat/ws.rtm.start [get]
func (cH *ChatHandlers) ChatStart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "home.html")
}

type ActiveUsersResponse struct {
	Count int `json:"count"`
}

// ChatActiveUsers
// @Summary  Number of active users in a chat
// @Tags     chat
// @Accept   json
// @Produce  json
// @Success  200  {object}  handlers.ActiveUsersResponse  "successful operation, returns number of active users"
// @Router   /user/active [get]
func (cH *ChatHandlers) ChatActiveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ActiveUsersResponse{Count: cH.cS.GetActiveUsers()})
}

// ChatConnect
// @Summary  Endpoint to start real time chat
// @Tags     chat
// @Accept   json
// @Produce  json
// @Param    token  query  string  true  "One time token for a logged user"
// @Success  100    "Upgrade to websocket protocol"
// @Failure  400    {string}  string  "Token is required|Token is not valid"
// @Router   /ws [get]
func (cH *ChatHandlers) ChatConnect(w http.ResponseWriter, r *http.Request) {
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
	ctx := context.WithValue(r.Context(), "token", token)
	cH.cS.ClientConnect(c)
	go c.Reader(ctx, cH.cS)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
