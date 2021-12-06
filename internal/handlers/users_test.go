package handlers

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/vetal4ik10/lets-go-chat/internal/models"
	urMoks "github.com/vetal4ik10/lets-go-chat/internal/reposetories/mocks"
	ottMoks "github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testUserHandler struct {
	request         string
	responseCode    int
	responseMessage string
}

func TestUserHandlers_UserCreate(t *testing.T) {
	var userCreateTestCases = []testUserHandler{
		{"{\"userName\": \"newTest\", \"password\" : \"testtest\"}", http.StatusCreated, "{\"id\":\"test-uid\",\"userName\":\"newTest\"}\n"},
		{"{\"userName\": \"exNewTest\", \"password\" : \"testtest\"}", http.StatusInternalServerError, "user is already exist\n"},
		{"{\"userName\": \"Tes\", \"password\" : \"testtest\"}", http.StatusInternalServerError, "user name should contain at least 4 chars\n"},
		{"{\"userName\": \"newTest\", \"password\" : \"test\"}", http.StatusInternalServerError, "user password should contain at least 8 chars\n"},
		{"{\"userName\": \"broken\", \"password\" : \"testtest\"}", http.StatusBadRequest, "something went wrong\n"},
	}

	for _, testCase := range userCreateTestCases {
		url := "/user"
		body := strings.NewReader(testCase.request)
		req, err := http.NewRequest(http.MethodPost, url, body)
		if err != nil {
			t.Fatal(err)
		}

		userRepo := new(urMoks.UserRepo)
		userRepo.On("Create", &models.User{
			UserName: "broken",
			Password: "testtest",
		}).Return(errors.New("something went wrong"))
		userRepo.On("Create", mock.Anything).Return(nil)

		userRepo.On("GetByUserName", "exNewTest").Return(&models.User{
			Uid:      "test-uid",
			UserName: "exNewTest",
			Password: "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S",
		}, nil)
		userRepo.On("GetByUserName", mock.Anything).Once().Return(nil, nil)
		userRepo.On("GetByUserName", mock.Anything).Return(&models.User{
			Uid:      "test-uid",
			UserName: "newTest",
			Password: "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S",
		}, nil)

		userH := NewUserHandlers(userRepo)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userH.UserCreate)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != testCase.responseCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.responseCode)
		}

		if rr.Body.String() != testCase.responseMessage {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), testCase.responseMessage)
		}
	}
}

func TestUserHandlers_Login(t *testing.T) {
	var userLoginTestCases = []testUserHandler{
		{"{\"userName\": \"newTest\", \"password\" : \"testtest\"}", http.StatusOK, "{\"url\":\"/chat/ws.rtm.start?token=test\"}\n"},
		{"{\"userName\": \"exNewTest\", \"password\" : \"testteste\"}", http.StatusBadRequest, "User name or password is incorrect.\n"},
		{"{\"userName\": \"newTest\", \"password1\" : \"testteste\"}", http.StatusBadRequest, "User name or password is incorrect.\n"},
	}

	for _, testCase := range userLoginTestCases {
		url := "/user/login"
		body := strings.NewReader(testCase.request)
		req, err := http.NewRequest(http.MethodPost, url, body)
		if err != nil {
			t.Fatal(err)
		}

		userRepo := new(urMoks.UserRepo)
		userRepo.On("GetByUserName", mock.Anything).Return(&models.User{
			Uid:      "test-uid",
			UserName: "newTest",
			Password: "$2a$14$XRTdMSXb/vRBbG8dlnihq.RBFphcDDMPS5EU8q8UX/DI0vHTL164S",
		}, nil)

		token := new(ottMoks.Token)
		token.On("GetSecret", mock.Anything).Return("test")
		tM := new(ottMoks.TokenManager)
		tM.On("NewToken", mock.Anything).Return(token, nil)

		userH := NewUserHandlers(userRepo)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userH.Login(tM, w, r)
		})
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != testCase.responseCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.responseCode)
		}

		if rr.Body.String() != testCase.responseMessage {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), testCase.responseMessage)
		}
	}
}
