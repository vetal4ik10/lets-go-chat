// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/vetal4ik10/lets-go-chat/internal/chat"
	"github.com/vetal4ik10/lets-go-chat/internal/chat_message"
	"github.com/vetal4ik10/lets-go-chat/internal/handlers"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitializeUserHandlers() *handlers.UserHandlers {
	db := initDatabase()
	userRepo := reposetories.NewUserRepo(db)
	tokenManager := onetimetoken.NewTokenManager(db, userRepo)
	userHandlers := handlers.NewUserHandlers(userRepo, tokenManager)
	return userHandlers
}

func InitializeChatHandlers() *handlers.ChatHandlers {
	db := initDatabase()
	userRepo := reposetories.NewUserRepo(db)
	tokenManager := onetimetoken.NewTokenManager(db, userRepo)
	chatMessageManager := chat_message.NewChatMessageManager(db, userRepo)
	chatServer := chat.NewChatServer(chatMessageManager)
	chatHandlers := handlers.NewChatHandlers(tokenManager, chatServer)
	return chatHandlers
}