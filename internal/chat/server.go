package chat

import "github.com/gorilla/websocket"

type chatServer struct {
	clients       map[*chatClient]bool
	send          chan []byte
	clientConnect chan *chatClient
	closeConnection chan *chatClient
}

type ChatServer interface {
	Run()
	ClientConnect(c *chatClient)
	SendMessage(m []byte)
	CloseConnection(c *chatClient)
	GetActiveUsers() int
}

func NewChatServer() *chatServer {
	return &chatServer{
		clients:       make(map[*chatClient]bool),
		send:          make(chan []byte),
		clientConnect: make(chan *chatClient),
		closeConnection: make(chan *chatClient),
	}
}

func (s *chatServer) Run() {
	for {
		select {
		case m := <-s.send:
			for c, a := range s.clients {
				if a {
					c.conn.WriteMessage(websocket.TextMessage, m)
				}
			}
		case c := <-s.clientConnect:
			s.clients[c] = true
		case c := <-s.closeConnection:
			s.clients[c] = false
			c.Close()
		}
	}
}

func (s *chatServer) ClientConnect(c *chatClient) {
	s.clientConnect <- c
}

func (s *chatServer) SendMessage(m []byte) {
	s.send <- m
}

func (s *chatServer) CloseConnection(c *chatClient) {
	s.closeConnection <- c
}

func (s *chatServer) GetActiveUsers() int {
	return len(s.clients)
}
