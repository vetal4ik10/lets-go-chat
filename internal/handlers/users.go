package handlers

import (
	"encoding/json"
	"errors"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/env"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"net/http"
)

type loginRequest struct {
	UserName string
	Password string
}

func validateUser(user *models.User) error {
	if len(user.UserName) < 4 {
		return errors.New("user name should contain at least 4 chars")
	} else if len(user.Password) < 8 {
		return errors.New("user password should contain at least 8 chars")
	} else if u, _ := env.GetUserRepo().GetByUserName(user.UserName); u != nil {
		return errors.New("user is already exist")
	}
	return nil
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	// Validate new user.
	err = validateUser(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userRequest.Password, _ = hasher.HashPassword(userRequest.Password)
	err = env.GetUserRepo().Create(&userRequest)

	// Create new user.
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		u, _ := env.GetUserRepo().GetByUserName(userRequest.UserName)
		response := map[string]string{"id": u.Uid, "userName": u.UserName}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest loginRequest

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	user, err := env.GetUserRepo().GetByUserName(loginRequest.UserName)
	if err != nil || !hasher.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "User name or password is incorrect.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Rate-Limit", "10")
	w.Header().Set("X-Expires-After", "10")
	response := map[string]string{"url": "ws://fancy-chat.io/ws&token=one-time-token"}
	json.NewEncoder(w).Encode(response)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	repo := env.GetUserRepo()
	repo.Create(&models.User{UserName: "test", Password: "test"})

	user, _ := repo.GetByUserName("test")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
