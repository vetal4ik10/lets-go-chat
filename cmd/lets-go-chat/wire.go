//go:build ignore
// +build ignore

package main

import (
	"github.com/google/wire"
	"github.com/vetal4ik10/lets-go-chat/internal/chat"
	"github.com/vetal4ik10/lets-go-chat/internal/chat_message"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
)

func InitializeUserHandlers() *handlers.UserHandlers {
	wire.Build(initDatabase, reposetories.NewUserRepo, onetimetoken.NewTokenManager, handlers.NewUserHandlers)
	return &handlers.UserHandlers{}
}

func InitializeChatHandlers() *handlers.ChatHandlers {
	wire.Build(
		initDatabase,
		reposetories.NewUserRepo,
		chat_message.NewChatMessageManager,
		onetimetoken.NewTokenManager,
		chat.NewChatServer,
		handlers.NewChatHandlers,
	)
	return &handlers.ChatHandlers{}
}
