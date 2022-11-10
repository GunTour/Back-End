// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/climber/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetClimber provides a mock function with given fields:
func (_m *Repository) GetClimber() (domain.Core, error) {
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

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}