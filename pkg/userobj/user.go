package userobj

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	UserName string
	Password string
}

var storage = map[string]*User{}

// AddUser add user to the storage.
func AddUser(user *User)  error {
	_, err := GetUserByName(user.UserName)
	if err != nil {
		for  {
			uuid := uuid.New().String()
			if _, exist := storage[uuid]; !exist {
				storage[uuid] = user
				break
			}
		}
		return nil
	}
	return errors.New("userobj is already exist")
}

// ValidateUser validate user object.
func ValidateUser(user *User)  error {
	if len(user.UserName) < 4 {
		return errors.New("userobj name should contain at least 4 chars")
	} else if len(user.Password) < 8 {
		return errors.New("userobj password should contain at least 8 chars")
	}
	return nil
}

// GetUserByName load user by name
func GetUserByName(name string) (*User, error) {
	for _, user := range storage {
		if user.UserName == name {
			return user, nil
		}
	}
	return nil, errors.New("userobj wasn't found")
}

func GetAllUsers() map[string]*User {
	return storage
}