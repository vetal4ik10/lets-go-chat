package onetimetoken

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"github.com/vetal4ik10/lets-go-chat/internal/reposetories"
)

type TokenManagerInterface interface {
	InitToken(s string) (Token, error)
	NewToken(u *models.User) (Token, error)
	Verified(t Token) bool
	Remove(t Token) error
}

type TokenManager struct {
	db       *sql.DB
	userRepo *reposetories.UserRepo
}

func NewTokenManager(db *sql.DB, userRepo *reposetories.UserRepo) *TokenManager {
	return &TokenManager{db, userRepo}
}

func (tM *TokenManager) InitToken(s string) (Token, error) {
	sqlStatement := "SELECT uid, secret FROM token WHERE secret=$1"
	rows, err := tM.db.Query(sqlStatement, s)
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

	u, _ := tM.userRepo.GetByUid(uid)
	return &token{u, s, tM}, nil
}

func (tM *TokenManager) NewToken(u *models.User) (Token, error) {
	s := uuid.New().String()
	sqlStatement := `INSERT INTO token (uid, secret) VALUES ($1, $2)`
	_, err := tM.db.Exec(sqlStatement, u.Uid, s)
	if err != nil {
		return nil, err
	}
	return &token{user: u, secret: s, tM: tM}, nil
}

func (tM *TokenManager) Verified(t Token) bool {
	sqlStatement := "SELECT uid FROM token WHERE uid=$1 AND secret=$2"
	rows, err := tM.db.Query(sqlStatement, t.GetUser().Uid, t.GetSecret())
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

func (tM *TokenManager) Remove(t Token) error {
	sqlStatement := "DELETE FROM token WHERE secret=$1"
	_, err := tM.db.Exec(sqlStatement, t.GetSecret())
	if err != nil {
		return err
	}
	return nil
}
