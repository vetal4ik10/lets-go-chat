package onetimetoken

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	urMoks "github.com/vetal4ik10/lets-go-chat/internal/reposetories/mocks"
	"testing"
)

func getTestTokenManager(db *sql.DB) TokenManagerInterface {
	userRepo := new(urMoks.UserRepo)
	userRepo.On("GetByUid", "test").Return(&models.User{
		Uid:      "test",
		UserName: "test",
		Password: "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S",
	}, nil)
	tM := NewTokenManager(db, userRepo)

	return tM
}

func TestTokenManager_InitToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := mock.NewRows([]string{"uid", "secret"}).AddRow("test", "test")
	mock.ExpectQuery("SELECT uid, secret FROM token").WillReturnRows(rows)

	tM := getTestTokenManager(db)
	_, err = tM.InitToken("test")

	if err != nil {
		t.Errorf("Token was generated with eror: %v", err)
	}
}

func TestNewToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO token").WillReturnResult(sqlmock.NewResult(1, 1))

	tM := getTestTokenManager(db)
	user := &models.User{"test-uid", "test", "testest"}
	_, err = tM.NewToken(user)

	if err != nil {
		t.Errorf("Token was generated with eror: %v", err)
	}
}

func TestTokenManager_Verified(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := mock.NewRows([]string{"uid"}).AddRow("test")
	mock.ExpectQuery("SELECT uid FROM token").WillReturnRows(rows)

	userRepo := new(urMoks.UserRepo)
	tM := NewTokenManager(db, userRepo)
	user := &models.User{"test-uid", "test", "testest"}
	token := NewToken(user, "test", tM)

	ok := tM.Verified(token)
	if !ok {
		t.Errorf("Verification doesn't work correctly")
	}
}

func TestTokenManager_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM token").WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := new(urMoks.UserRepo)
	tM := NewTokenManager(db, userRepo)
	user := &models.User{"test-uid", "test", "testest"}
	token := NewToken(user, "test", tM)

	err = tM.Remove(token)
	if err != nil {
		t.Errorf("Token was removed with eror: %v", err)
	}
}
