package chat

import (
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/internal/chat_message"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"log"
)

type chatClient struct {
	conn  *websocket.Conn
	token onetimetoken.Token
}

type ChatClient interface {
	GetConnection() *websocket.Conn
	getToken() onetimetoken.Token
	Reader(s ChatServer)
	Close()
}

func NewChatClient(conn *websocket.Conn, t onetimetoken.Token) *chatClient {
	return &chatClient{conn, t}
}

func (c *chatClient) GetConnection() *websocket.Conn {
	return c.conn
}

func (c *chatClient) getToken() onetimetoken.Token {
	return c.token
}

// Reader to fetch message from connection and send to server.
func (c *chatClient) Reader(s ChatServer) {
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {

			log.Println(err)
			break
		}
		u := c.getToken().GetUser()
		cm := chat_message.NewChatMessage(u, m)
		s.SendMessage(cm)
	}
}

func (c *chatClient) Close() {
	c.token.Remove()
}
