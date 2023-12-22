// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	repository "go_notifier/internal/db/repository"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetUserIDAndTimezoneByUUID provides a mock function with given fields: uuid
func (_m *UserRepository) GetUserIDAndTimezoneByUUID(uuid string) (*repository.UserIdTimezone, error) {
	ret := _m.Called(uuid)

	var r0 *repository.UserIdTimezone
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.UserIdTimezone, error)); ok {
		return rf(uuid)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.UserIdTimezone); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.UserIdTimezone)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIDByUUID provides a mock function with given fields: uuid
func (_m *UserRepository) GetUserIDByUUID(uuid string) (int64, error) {
	ret := _m.Called(uuid)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(uuid)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(int64)
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
