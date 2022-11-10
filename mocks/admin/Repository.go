// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "GunTour/features/admin/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteProduct provides a mock function with given fields: id
func (_m *Repository) DeleteProduct(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRanger provides a mock function with given fields: id
func (_m *Repository) DeleteRanger(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditRanger provides a mock function with given fields: data, datas, id
func (_m *Repository) EditRanger(data domain.RangerCore, datas domain.UserCore, id uint) (domain.RangerCore, domain.UserCore, error) {
	ret := _m.Called(data, datas, id)

	var r0 domain.RangerCore
	if rf, ok := ret.Get(0).(func(domain.RangerCore, domain.UserCore, uint) domain.RangerCore); ok {
		r0 = rf(data, datas, id)
	} else {
		r0 = ret.Get(0).(domain.RangerCore)
	}

	var r1 domain.UserCore
	if rf, ok := ret.Get(1).(func(domain.RangerCore, domain.UserCore, uint) domain.UserCore); ok {
		r1 = rf(data, datas, id)
	} else {
		r1 = ret.Get(1).(domain.UserCore)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(domain.RangerCore, domain.UserCore, uint) error); ok {
		r2 = rf(data, datas, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllRanger provides a mock function with given fields:
func (_m *Repository) GetAllRanger() ([]domain.RangerCore, []domain.RangerCore, error) {
	ret := _m.Called()

	var r0 []domain.RangerCore
	if rf, ok := ret.Get(0).(func() []domain.RangerCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RangerCore)
		}
	}

	var r1 []domain.RangerCore
	if rf, ok := ret.Get(1).(func() []domain.RangerCore); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]domain.RangerCore)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPendaki provides a mock function with given fields:
func (_m *Repository) GetPendaki() ([]domain.BookingCore, domain.ClimberCore, error) {
	ret := _m.Called()

	var r0 []domain.BookingCore
	if rf, ok := ret.Get(0).(func() []domain.BookingCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BookingCore)
		}
	}

	var r1 domain.ClimberCore
	if rf, ok := ret.Get(1).(func() domain.ClimberCore); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(domain.ClimberCore)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProduct provides a mock function with given fields: page
func (_m *Repository) GetProduct(page int) ([]domain.ProductCore, int, int, error) {
	ret := _m.Called(page)

	var r0 []domain.ProductCore
	if rf, ok := ret.Get(0).(func(int) []domain.ProductCore); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProductCore)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(int) int); ok {
		r2 = rf(page)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(int) error); ok {
		r3 = rf(page)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// InsertClimber provides a mock function with given fields: data
func (_m *Repository) InsertClimber(data domain.ClimberCore) (domain.ClimberCore, error) {
	ret := _m.Called(data)

	var r0 domain.ClimberCore
	if rf, ok := ret.Get(0).(func(domain.ClimberCore) domain.ClimberCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.ClimberCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ClimberCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertProduct provides a mock function with given fields: newProduct
func (_m *Repository) InsertProduct(newProduct domain.ProductCore) (domain.ProductCore, error) {
	ret := _m.Called(newProduct)

	var r0 domain.ProductCore
	if rf, ok := ret.Get(0).(func(domain.ProductCore) domain.ProductCore); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(domain.ProductCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ProductCore) error); ok {
		r1 = rf(newProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: newProduct
func (_m *Repository) UpdateProduct(newProduct domain.ProductCore) (domain.ProductCore, error) {
	ret := _m.Called(newProduct)

	var r0 domain.ProductCore
	if rf, ok := ret.Get(0).(func(domain.ProductCore) domain.ProductCore); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(domain.ProductCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ProductCore) error); ok {
		r1 = rf(newProduct)
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
