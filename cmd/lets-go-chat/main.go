package main

import (
	"fmt"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/pkg/role"
)

func main() {
	fmt.Println("Hello")
	password := "secret"
	hash, _ := hasher.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := hasher.CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
	role.AdminMessage()
}