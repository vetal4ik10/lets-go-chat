package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
)

type PostgresUserRepo struct {
	Db *sql.DB
}

func (r PostgresUserRepo) Create(u *models.User) error {
	uid := uuid.New().String()
	sqlStatement := `INSERT INTO users (uid, name, pass)
		VALUES ($1, $2, $3)`

	_, err := r.Db.Exec(sqlStatement, uid, u.UserName, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r PostgresUserRepo) GetByUserName(username string) (*models.User, error) {
	sqlStatement := "SELECT uid, name, pass FROM users WHERE name=$1"
	rows, err := r.Db.Query(sqlStatement, username)
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
