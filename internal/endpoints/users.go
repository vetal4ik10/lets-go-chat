package endpoints

import (
	"encoding/json"
	"github.com/vetal4ik10/lets-go-chat/pkg/userobj"
	"net/http"
)



func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user userobj.User

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,"Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	// Validate new user.
	err = userobj.ValidateUser(&user)
	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	// Create new user.
	err = userobj.AddUser(&user)
	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}
}

func UserList(w http.ResponseWriter, r *http.Request) {
	users := userobj.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func Login(w http.ResponseWriter, r *http.Request) {
	var user userobj.User

	// Parse post body.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	saveUser, err := userobj.GetUserByName(user.UserName)
	if err != nil || saveUser.Password != user.Password {
		http.Error(w, "User name or password is incorrect.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Rate-Limit", "10")
	w.Header().Set("X-Expires-After", "10")
	response := map[string]string{"url": "ws://fancy-chat.io/ws&token=one-time-token"}
	json.NewEncoder(w).Encode(response)
}