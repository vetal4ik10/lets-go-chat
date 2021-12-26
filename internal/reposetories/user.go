package reposetories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type userRepo struct {
	db *sql.DB
}

type UserRepo interface {
	Create(user *models.User) error
	GetByUserName(username string) (*models.User, error)
	GetByUid(uid string) (*models.User, error)
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (uR *userRepo) Create(user *models.User) error {
	user.Password, _ = hasher.HashPassword(user.Password)
	uid := uuid.New().String()
	sqlStatement := `INSERT INTO users (uid, name, pass)
		VALUES ($1, $2, $3)`

	_, err := uR.db.Exec(sqlStatement, uid, user.UserName, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (uR *userRepo) GetByUserName(username string) (*models.User, error) {
	sqlStatement := "SELECT uid, name, pass FROM users WHERE name=$1"
	rows, err := uR.db.Query(sqlStatement, username)
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

func (uR *userRepo) GetByUid(uid string) (*models.User, error) {
	sqlStatement := "SELECT uid, name, pass FROM users WHERE uid=$1"
	rows, err := uR.db.Query(sqlStatement, uid)
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
