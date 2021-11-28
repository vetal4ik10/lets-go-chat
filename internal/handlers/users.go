package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"net/http"
)

type UserHandlers struct {
	repo reposetories.UserRepo
}

func NewUserHandlers(repo reposetories.UserRepo) *UserHandlers {
	return &UserHandlers{repo}
}

func (uH *UserHandlers) validateUser(user *createUserRequest) error {
	if len(user.UserName) < 4 {
		return errors.New("user name should contain at least 4 chars")
	} else if len(user.Password) < 8 {
		return errors.New("user password should contain at least 8 chars")
	} else if u, _ := uH.repo.GetByUserName(user.UserName); u != nil {
		return errors.New("user is already exist")
	}
	return nil
}

func (uH *UserHandlers) UserCreate(w http.ResponseWriter, r *http.Request) {
	var userRequest createUserRequest

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	// Validate new user.
	err = uH.validateUser(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create new user object.
	user := &models.User{
		UserName: userRequest.UserName,
		Password: userRequest.Password,
	}

	err = uH.repo.Create(user);
	// Create new user.
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		u, _ := uH.repo.GetByUserName(userRequest.UserName)
		response := map[string]string{"id": u.Uid, "userName": u.UserName}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}


type loginRequest struct {
	UserName string `json:"userName"`
	Password string	`json:"password"`
}

func (uH *UserHandlers) Login(tM *onetimetoken.TokenManager, w http.ResponseWriter, r *http.Request) {
	var loginRequest loginRequest

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	user, err := uH.repo.GetByUserName(loginRequest.UserName)
	if err != nil || !hasher.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "User name or password is incorrect.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Rate-Limit", "10")
	w.Header().Set("X-Expires-After", "10")

	t, _ := tM.NewToken(user)


	t, _ = tM.InitToken(t.Secret)

	v := tM.Verified(t)
	fmt.Println(v)

	tM.Remove(t)

	response := map[string]string{"token": t.Secret}
	json.NewEncoder(w).Encode(response)
}

func (uH *UserHandlers) UserList(w http.ResponseWriter, r *http.Request) {
	uH.repo.Create(&models.User{UserName: "test", Password: "test"})
	user, _ := uH.repo.GetByUserName("test")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (uH *UserHandlers) User(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}
