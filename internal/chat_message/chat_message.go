package chat_message

import (
	"encoding/json"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"time"
)

type ChatMessage interface {
	GetUser() *models.User
	GetText() []byte
	GetTime() *time.Time
	Json() []byte
}

type chatMessage struct {
	user *models.User
	text []byte
	time *time.Time
}

func NewChatMessage(user *models.User, message []byte) *chatMessage {
	return &chatMessage{user, message, nil}
}

func (m *chatMessage) GetUser() *models.User {
	return m.user
}

func (m *chatMessage) GetText() []byte {
	return m.text
}

func (m *chatMessage) GetTime() *time.Time {
	return m.time
}

func (m *chatMessage) Json() []byte {

	c, _ := json.Marshal(map[string]string{
		"time": m.GetTime().Format(time.Kitchen),
		"text": string(m.text),
		"user": m.user.UserName,
	})
	return c
}
