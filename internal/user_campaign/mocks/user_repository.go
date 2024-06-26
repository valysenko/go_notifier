// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	user "go_notifier/internal/user"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetUserIDAndTimezoneByUUID provides a mock function with given fields: uuid
func (_m *UserRepository) GetUserIDAndTimezoneByUUID(uuid string) (*user.UserIdTimezone, error) {
	ret := _m.Called(uuid)

	var r0 *user.UserIdTimezone
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.UserIdTimezone, error)); ok {
		return rf(uuid)
	}
	if rf, ok := ret.Get(0).(func(string) *user.UserIdTimezone); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserIdTimezone)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
