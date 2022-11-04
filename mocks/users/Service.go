// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/users/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) (domain.Core, error) {
	ret := _m.Called(id)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(int) domain.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *Service) Insert(data domain.Core) (domain.Core, error) {
	ret := _m.Called(data)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: input
func (_m *Service) Login(input domain.Core) (domain.Core, error) {
	ret := _m.Called(input)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data, id
func (_m *Service) Update(data domain.Core, id int) (domain.Core, error) {
	ret := _m.Called(data, id)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core, int) domain.Core); ok {
		r0 = rf(data, id)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core, int) error); ok {
		r1 = rf(data, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
