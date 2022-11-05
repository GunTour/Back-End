// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/booking/domain"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// DeleteData provides a mock function with given fields: idBooking
func (_m *Services) DeleteData(idBooking uint) error {
	ret := _m.Called(idBooking)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(idBooking)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: idUser
func (_m *Services) GetAll(idUser uint) ([]domain.Core, error) {
	ret := _m.Called(idUser)

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func(uint) []domain.Core); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: idBooking
func (_m *Services) GetDetail(idBooking uint) (domain.Core, error) {
	ret := _m.Called(idBooking)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(idBooking)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idBooking)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRangerBooking provides a mock function with given fields: idRanger
func (_m *Services) GetRangerBooking(idRanger uint) ([]domain.Core, error) {
	ret := _m.Called(idRanger)

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func(uint) []domain.Core); ok {
		r0 = rf(idRanger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idRanger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: newBooking
func (_m *Services) InsertData(newBooking domain.Core) (domain.Core, error) {
	ret := _m.Called(newBooking)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newBooking)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newBooking)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateData provides a mock function with given fields: newBooking
func (_m *Services) UpdateData(newBooking domain.Core) (domain.Core, error) {
	ret := _m.Called(newBooking)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newBooking)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newBooking)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMidtrans provides a mock function with given fields: newBooking
func (_m *Services) UpdateMidtrans(newBooking domain.Core) error {
	ret := _m.Called(newBooking)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Core) error); ok {
		r0 = rf(newBooking)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
