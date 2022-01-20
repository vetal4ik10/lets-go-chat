package onetimetoken

import (
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type Token interface {
	GetUser() *models.User
	GetSecret() string
	Remove() error
}

type token struct {
	user   *models.User
	secret string
	tM     TokenManagerInterface
}

func NewToken(user *models.User, secret string, tM TokenManagerInterface) *token {
	return &token{user, secret, tM}
}

func (t *token) GetUser() *models.User {
	return t.user
}

func (t *token) GetSecret() string {
	return t.secret
}

func (t *token) Remove() error {
	err := t.tM.Remove(t)
	return err
}
