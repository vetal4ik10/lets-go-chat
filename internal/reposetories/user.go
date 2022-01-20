package reposetories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/hasher"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepoInterface interface {
	Create(user *models.User) error
	GetByUserName(username string) (*models.User, error)
	GetByUid(uid string) (*models.User, error)
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (uR *UserRepo) Create(user *models.User) error {
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

func (uR *UserRepo) GetByUserName(username string) (*models.User, error) {
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

func (uR *UserRepo) GetByUid(uid string) (*models.User, error) {
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
