// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/vetal4ik10/lets-go-chat/internal/models"
)

// Token is an autogenerated mock type for the Token type
type Token struct {
	mock.Mock
}

// GetSecret provides a mock function with given fields:
func (_m *Token) GetSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUser provides a mock function with given fields:
func (_m *Token) GetUser() *models.User {
	ret := _m.Called()

	var r0 *models.User
	if rf, ok := ret.Get(0).(func() *models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	return r0
}

// Remove provides a mock function with given fields:
func (_m *Token) Remove() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
