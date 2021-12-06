package handlers

import (
	"github.com/stretchr/testify/mock"
	cMoks "github.com/vetal4ik10/lets-go-chat/internal/chat/mocks"
	ottMoks "github.com/vetal4ik10/lets-go-chat/pkdg/onetimetoken/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChatHandlers_ChatStart(t *testing.T) {
	url := "/chat/ws.rtm.start"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	cS := new(cMoks.ChatServer)
	tM := new(ottMoks.TokenManager)
	cH := NewChatHandlers(tM, cS)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cH.ChatStart)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestChatHandlers_ChatActiveUsers(t *testing.T) {
	url := "/user/active"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	cS := new(cMoks.ChatServer)
	cS.On("GetActiveUsers", mock.Anything).Return(1)
	tM := new(ottMoks.TokenManager)
	cH := NewChatHandlers(tM, cS)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cH.ChatActiveUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expect := "{\"count\":1}\n"
	if rr.Body.String() != expect {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expect)
	}
}

func TestChatHandlers_ChatConnect(t *testing.T) {
	url := "/ws"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	cS := new(cMoks.ChatServer)
	cS.On("GetActiveUsers", mock.Anything).Return(1)
	tM := new(ottMoks.TokenManager)
	cH := NewChatHandlers(tM, cS)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cH.ChatConnect)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Connects works without token")
	}
}
