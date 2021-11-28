package handlers

type createUserRequest struct {
	UserName string `json:"userName"`
	Password string	`json:"password"`
}
