// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/climber/domain"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// ShowClimber provides a mock function with given fields:
func (_m *Services) ShowClimber() (domain.Core, error) {
	ret := _m.Called()

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func() domain.Core); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServices(t mockConstructorTestingTNewServices) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}