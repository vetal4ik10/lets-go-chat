package reposetories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type user struct {
	db *sql.DB
}

func (u *user) Create(user *models.User) error {
	user.Password, _ = hasher.HashPassword(user.Password)
	uid := uuid.New().String()
	sqlStatement := `INSERT INTO users (uid, name, pass)
		VALUES ($1, $2, $3)`

	_, err := u.db.Exec(sqlStatement, uid, user.UserName, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) GetByUserName(username string) (*models.User, error) {
	sqlStatement := "SELECT uid, name, pass FROM users WHERE name=$1"
	rows, err := u.db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	rows.Next()
	err = rows.Scan(&user.Uid, &user.UserName, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, err
}

type UserRepo interface {
	Create(user *models.User) error
	GetByUserName(username string) (*models.User, error)
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &user{db: db}
}

