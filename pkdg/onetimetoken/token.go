package onetimetoken

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
)

type token struct {
	User *models.User
	Secret string
}


type TokenManager struct {
	db *sql.DB
	userRepo reposetories.UserRepo
}

func NewTokenManager(db *sql.DB, userRepo reposetories.UserRepo) *TokenManager {
	return &TokenManager{db, userRepo}
}

func (tm *TokenManager) InitToken(s string) (*token, error)  {
	sqlStatement := "SELECT uid, secret FROM token WHERE secret=$1"
	rows, err := tm.db.Query(sqlStatement, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var uid, tS string

	rows.Next()
	err = rows.Scan(&uid, &tS)
	if err != nil {
		return nil, err
	}

	u, _ := tm.userRepo.GetByUid(uid)
	return &token{u, s}, nil
}

func (tm *TokenManager) NewToken(u *models.User) (*token, error) {
	s := uuid.New().String()
	sqlStatement := `INSERT INTO token (uid, secret)
		VALUES ($1, $2)`
	_, err := tm.db.Exec(sqlStatement, u.Uid, s)
	if err != nil {
		return nil, err
	}
	return &token{User: u, Secret: s}, nil
}

func (tm *TokenManager) Verified(t *token) bool {
	sqlStatement := "SELECT uid FROM token WHERE uid=$1 AND secret=$2"
	rows, err := tm.db.Query(sqlStatement, t.User.Uid, t.Secret)
	defer rows.Close()
	if err != nil {
		return false
	}
	var exists string


	rows.Next()
	err = rows.Scan(&exists)
	if err != nil {
		return false
	}


	return exists != ""
}

func (tm *TokenManager) Remove(t *token) error {
	sqlStatement := "DELETE FROM token WHERE secret=$1"
	_, err := tm.db.Exec(sqlStatement, t.Secret)
	if err != nil {
		return err
	}
	return nil
}
