// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	chat "github.com/vetal4ik10/lets-go-chat/internal/chat"
)

// ChatServer is an autogenerated mock type for the ChatServer type
type ChatServer struct {
	mock.Mock
}

// ClientConnect provides a mock function with given fields: c
func (_m *ChatServer) ClientConnect(c chat.ChatClient) {
	_m.Called(c)
}

// CloseConnection provides a mock function with given fields: c
func (_m *ChatServer) CloseConnection(c chat.ChatClient) {
	_m.Called(c)
}

// GetActiveUsers provides a mock function with given fields:
func (_m *ChatServer) GetActiveUsers() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *ChatServer) Run() {
	_m.Called()
}

// SendMessage provides a mock function with given fields: m
func (_m *ChatServer) SendMessage(m []byte) {
	_m.Called(m)
}
