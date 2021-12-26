package reposetories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO users (uid, name, pass)").WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := NewUserRepo(db)
	userRepo.Create(&models.User{
		UserName: "newTest",
		Password: "testtest",
	})
}

func TestUserRepo_GetByUid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := mock.NewRows([]string{"uid", "name", "pass"}).AddRow("test-uid", "newTest", "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S")
	mock.ExpectQuery("SELECT uid, name, pass FROM users").WithArgs("test-uid").WillReturnRows(rows)

	userRepo := NewUserRepo(db)
	_, err = userRepo.GetByUid("test-uid")
	if err != nil {
		t.Errorf("GetByUid comes with error: %v", err)
	}
}

func TestUserRepo_GetByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := mock.NewRows([]string{"uid", "name", "pass"}).AddRow("test-uid", "newTest", "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S")
	mock.ExpectQuery("SELECT uid, name, pass FROM users").WithArgs("newTest").WillReturnRows(rows)

	userRepo := NewUserRepo(db)
	_, err = userRepo.GetByUserName("newTest")
	if err != nil {
		t.Errorf("GetByUserName comes with error: %v", err)
	}
}
