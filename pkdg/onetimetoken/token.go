package onetimetoken

import (
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type token struct {
	user   *models.User
	secret string
	tM 		 *TokenManager
}

type Token interface {
	GetUser() *models.User
	GetSecret() string
	Remove()
}

func (t *token) GetUser() *models.User {
	return t.user
}

func (t *token) GetSecret() string {
	return t.secret
}

func (t *token) Remove() {
	t.tM.Remove(t)
}