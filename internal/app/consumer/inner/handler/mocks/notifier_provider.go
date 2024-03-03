// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	notifier "go_notifier/internal/notifier"

	mock "github.com/stretchr/testify/mock"
)

// NotifierProvider is an autogenerated mock type for the NotifierProvider type
type NotifierProvider struct {
	mock.Mock
}

// Provide provides a mock function with given fields: appType
func (_m *NotifierProvider) Provide(appType string) notifier.Notifier {
	ret := _m.Called(appType)

	var r0 notifier.Notifier
	if rf, ok := ret.Get(0).(func(string) notifier.Notifier); ok {
		r0 = rf(appType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(notifier.Notifier)
		}
	}

	return r0
}

type mockConstructorTestingTNewNotifierProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotifierProvider creates a new instance of NotifierProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotifierProvider(t mockConstructorTestingTNewNotifierProvider) *NotifierProvider {
	mock := &NotifierProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
