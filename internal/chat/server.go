package chat

import "github.com/gorilla/websocket"

type chatServer struct {
	clients         map[ChatClient]bool
	send            chan []byte
	clientConnect   chan ChatClient
	closeConnection chan ChatClient
}

type ChatServer interface {
	Run()
	ClientConnect(c ChatClient)
	SendMessage(m []byte)
	CloseConnection(c ChatClient)
	GetActiveUsers() int
}

func NewChatServer() *chatServer {
	return &chatServer{
		clients:         make(map[ChatClient]bool),
		send:            make(chan []byte),
		clientConnect:   make(chan ChatClient),
		closeConnection: make(chan ChatClient),
	}
}

func (s *chatServer) Run() {
	for {
		select {
		case m := <-s.send:
			for c, a := range s.clients {
				if a {
					c.GetConnection().WriteMessage(websocket.TextMessage, m)
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

func (s *chatServer) ClientConnect(c ChatClient) {
	s.clientConnect <- c
}

func (s *chatServer) SendMessage(m []byte) {
	s.send <- m
}

func (s *chatServer) CloseConnection(c ChatClient) {
	s.closeConnection <- c
}

func (s *chatServer) GetActiveUsers() int {
	return len(s.clients)
}
