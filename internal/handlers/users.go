package handlers

import (
	"encoding/json"
	"errors"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
	"github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken"
	"net/http"
)

type UserHandlers struct {
	repo *reposetories.UserRepo
	tM   *onetimetoken.TokenManager
}

func NewUserHandlers(repo *reposetories.UserRepo, tM *onetimetoken.TokenManager) *UserHandlers {
	return &UserHandlers{repo, tM}
}

func (uH *UserHandlers) validateUser(user *CreateUserRequest) error {
	if len(user.UserName) < 4 {
		return errors.New("user name should contain at least 4 chars")
	} else if len(user.Password) < 8 {
		return errors.New("user password should contain at least 8 chars")
	} else if u, _ := uH.repo.GetByUserName(user.UserName); u != nil {
		return errors.New("user is already exist")
	}
	return nil
}

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

// UserCreate
// @Summary  Register (create) user
// @Tags     users
// @Accept   json
// @Produce  json
// @Param    message  body      handlers.CreateUserRequest   true  "Created user object"
// @Success  201      {object}  handlers.CreateUserResponse  true  "user created"
// @Failure  400      {string}  string                       "Bad request, empty username or password"
// @Failure  500      {string}  string                       "user name should contain at least 4 chars|user password should contain at least 8 chars"
// @Router   /user [post]
func (uH *UserHandlers) UserCreate(w http.ResponseWriter, r *http.Request) {
	var userRequest CreateUserRequest
	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		ErrorResponse(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	// Validate new user.
	err = uH.validateUser(&userRequest)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create new user object.
	user := &models.User{
		UserName: userRequest.UserName,
		Password: userRequest.Password,
	}

	err = uH.repo.Create(user)
	// Create new user.
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		u, _ := uH.repo.GetByUserName(userRequest.UserName)
		response := CreateUserResponse{Id: u.Uid, UserName: u.UserName}
		json.NewEncoder(w).Encode(response)
	} else {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}
}

type LoginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUserResonse struct {
	Url string `json:"url"`
}

// Login
// @Summary  Logs user into the system
// @Tags     users
// @Accept   json
// @Produce  json
// @Param    message  body      handlers.LoginUserRequest  true  "User credendials"
// @Success  201      {object}  handlers.LoginUserResonse  true  "successful operation, returns link to join chat"
// @Failure  400      {string}  string                     "Bad request, empty username or password|User name or password is incorrect."
// @Router   /user/login [post]
func (uH *UserHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginUserRequest

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		ErrorResponse(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	user, err := uH.repo.GetByUserName(loginRequest.UserName)
	if err != nil || !hasher.CheckPasswordHash(loginRequest.Password, user.Password) {
		ErrorResponse(w, "User name or password is incorrect.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Rate-Limit", "10")
	w.Header().Set("X-Expires-After", "10")

	t, _ := uH.tM.NewToken(user)
	response := LoginUserResonse{Url: "/chat/ws.rtm.start?token=" + t.GetSecret()}
	json.NewEncoder(w).Encode(response)
}
