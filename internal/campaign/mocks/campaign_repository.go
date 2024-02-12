// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	campaign "go_notifier/internal/campaign"

	mock "github.com/stretchr/testify/mock"
)

// CampaignRepository is an autogenerated mock type for the CampaignRepository type
type CampaignRepository struct {
	mock.Mock
}

// GetScheduledNotifications provides a mock function with given fields: day, currentTime
func (_m *CampaignRepository) GetScheduledNotifications(day string, currentTime string) ([]*campaign.ScheduledNotification, error) {
	ret := _m.Called(day, currentTime)

	var r0 []*campaign.ScheduledNotification
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*campaign.ScheduledNotification, error)); ok {
		return rf(day, currentTime)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*campaign.ScheduledNotification); ok {
		r0 = rf(day, currentTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*campaign.ScheduledNotification)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(day, currentTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCampaignRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCampaignRepository creates a new instance of CampaignRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCampaignRepository(t mockConstructorTestingTNewCampaignRepository) *CampaignRepository {
	mock := &CampaignRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
