// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/google/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetPesan provides a mock function with given fields:
func (_m *Repository) GetPesan() (domain.PesanCore, domain.RangerCore) {
	ret := _m.Called()

	var r0 domain.PesanCore
	if rf, ok := ret.Get(0).(func() domain.PesanCore); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.PesanCore)
	}

	var r1 domain.RangerCore
	if rf, ok := ret.Get(1).(func() domain.RangerCore); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(domain.RangerCore)
	}

	return r0, r1
}

// GetPesanCal provides a mock function with given fields:
func (_m *Repository) GetPesanCal() domain.BookingCore {
	ret := _m.Called()

	var r0 domain.BookingCore
	if rf, ok := ret.Get(0).(func() domain.BookingCore); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.BookingCore)
	}

	return r0
}

// InsertCode provides a mock function with given fields: code
func (_m *Repository) InsertCode(code domain.Code) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Code) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
