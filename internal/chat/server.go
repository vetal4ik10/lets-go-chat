package chat

import (
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/internal/chat_message"
)

type chatServer struct {
	clients         map[ChatClient]bool
	send            chan chat_message.ChatMessage
	clientConnect   chan ChatClient
	closeConnection chan ChatClient
	messageManager  chat_message.ChatMessageManager
}

type ChatServer interface {
	Run()
	ClientConnect(c ChatClient)
	SendMessage(m chat_message.ChatMessage)
	CloseConnection(c ChatClient)
	GetActiveUsers() int
}

func NewChatServer(mM chat_message.ChatMessageManager) *chatServer {
	return &chatServer{
		clients:         make(map[ChatClient]bool),
		send:            make(chan chat_message.ChatMessage),
		clientConnect:   make(chan ChatClient),
		closeConnection: make(chan ChatClient),
		messageManager:  mM,
	}
}

func (s *chatServer) Run() {
	for {
		select {
		case m := <-s.send:
			cm, _ := s.messageManager.SaveMessage(m)
			for c, a := range s.clients {
				if a {
					c.GetConnection().WriteMessage(websocket.TextMessage, cm.Json())
				}
			}
		case c := <-s.clientConnect:
			messages := s.messageManager.LoadAllMessages()
			for _, m := range messages {
				c.GetConnection().WriteMessage(websocket.TextMessage, m.Json())
			}
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

func (s *chatServer) SendMessage(m chat_message.ChatMessage) {
	s.send <- m
}

func (s *chatServer) CloseConnection(c ChatClient) {
	s.closeConnection <- c
}

func (s *chatServer) GetActiveUsers() int {
	return len(s.clients)
}
