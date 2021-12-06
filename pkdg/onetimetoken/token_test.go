package onetimetoken

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	urMoks "github.com/vetal4ik10/lets-go-chat/internal/reposetories/mocks"
	"reflect"
	"testing"
)

func getTestToken(t *testing.T) *token {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := new(urMoks.UserRepo)
	tM := NewTokenManager(db, userRepo)
	user := &models.User{"test-uid", "test", "testest"}
	token := NewToken(user, "test", tM)
	return token
}

func TestToken_GetSecret(t *testing.T) {
	expect := "test"
	secret := getTestToken(t).GetSecret()
	if secret != expect {
		t.Errorf("secret is incorrect: got %v want %v",
			secret, expect)
	}
}

func TestToken_GetUser(t *testing.T) {
	expect := &models.User{"test-uid", "test", "testest"}
	tUser := getTestToken(t).GetUser()
	if !reflect.DeepEqual(expect, tUser) {
		t.Errorf("user is incorrect: got %v want %v",
			tUser, expect)
	}
}

func TestToken_Remove(t *testing.T) {
	token := getTestToken(t)
	if err := token.Remove(); err != nil {
		t.Errorf("Remove comes with error: %v", err)
	}
}
