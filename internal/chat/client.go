package chat

import (
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"log"
)

type chatClient struct {
	conn  *websocket.Conn
	token onetimetoken.Token
}

func NewChatClient(conn *websocket.Conn, t onetimetoken.Token) *chatClient {
	return &chatClient{conn, t}
}

// Reader to fetch message from connection and send to server.
func (c *chatClient) Reader(s ChatServer) {
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {

			log.Println(err)
			break
		}
		s.SendMessage(m)
	}
}

func (c *chatClient) Close() {
	c.token.Remove()
}