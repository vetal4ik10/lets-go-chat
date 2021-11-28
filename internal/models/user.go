package models

type User struct {
	Uid      string
	UserName string
	Password string
}

type UserRepo interface {
	Create(u *User) error
	GetByUserName(username string) (*User, error)
}
