package services

import (
	"GunTour/features/booking/domain"
	mocks "GunTour/mocks/booking"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {
	repo := new(mocks.Repository)
	Start, _ := time.Parse("2006-01-02", "2022-11-10")
	End, _ := time.Parse("2006-01-02", "2022-11-12")
	input := domain.Core{IdUser: uint(1), DateStart: Start, DateEnd: End, Entrance: "Cibodas", Ticket: 3,
		OrderId: "Order-101", BookingProductCores: nil, IdRanger: uint(1), GrossAmount: 120000, StatusBooking: "unpaid"}
	returnRespon := domain.Core{ID: 1, IdUser: uint(1), DateStart: Start, DateEnd: End, Entrance: "Cibodas", Ticket: 3,
		OrderId: "Order-101", BookingProductCores: nil, IdRanger: uint(1), GrossAmount: 120000, StatusBooking: "unpaid"}

	t.Run("create success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.InsertData(input)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed create booking", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.InsertData(domain.Core{})
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.Repository)
	Start, _ := time.Parse("2006-01-02", "2022-11-10")
	End, _ := time.Parse("2006-01-02", "2022-11-12")
	data := domain.Core{ID: 1, IdUser: uint(1), DateStart: Start, DateEnd: End, Entrance: "Cibodas", Ticket: 3,
		OrderId: "Order-101", BookingProductCores: nil, IdRanger: uint(1), GrossAmount: 120000, StatusBooking: "unpaid"}
	returnRespon := domain.Core{ID: 1, IdUser: uint(1), DateStart: Start, DateEnd: End, Entrance: "Cibodas", Ticket: 3,
		OrderId: "Order-101", BookingProductCores: nil, IdRanger: uint(1), GrossAmount: 120000, StatusBooking: "unpaid"}

	t.Run("success update booking", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.UpdateData(data)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed update booking", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New("some problem on database")).Once()
		usecase := New(repo)

		res, err := usecase.UpdateData(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.Repository)

	t.Run("success delete data", func(t *testing.T) {

		repo.On("Delete", mock.Anything).Return(nil).Once()

		useCase := New(repo)

		err := useCase.DeleteData(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("Delete", mock.Anything).Return(errors.New("error")).Once()

		useCase := New(repo)

		err := useCase.DeleteData(uint(2))
		assert.Error(t, errors.New("error"))
		assert.Equal(t, err, err)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get All Booking", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{{ID: uint(1), GrossAmount: 120000, OrderId: "Order-101", Link: "www.mid", StatusBooking: "unpaid"}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetAll(uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Booking", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetAll(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetRanger(t *testing.T) {
	repo := mocks.NewRepository(t)
	Start, _ := time.Parse("2006-01-02", "2022-11-10")
	End, _ := time.Parse("2006-01-02", "2022-11-12")
	t.Run("Sukses Get Ranger Booking", func(t *testing.T) {
		repo.On("GetRanger", mock.Anything).Return([]domain.Core{{ID: uint(1), IdUser: uint(1), FullName: "Bambang", Phone: "0822", DateStart: Start, DateEnd: End, Ticket: 3}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetRangerBooking(uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Ranger Booking", func(t *testing.T) {
		repo.On("GetRanger", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetRangerBooking(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDetail(t *testing.T) {
	repo := mocks.NewRepository(t)
	Start, _ := time.Parse("2006-01-02", "2022-11-10")
	End, _ := time.Parse("2006-01-02", "2022-11-12")
	returnRespon := domain.Core{ID: 1, IdUser: uint(1), DateStart: Start, DateEnd: End, Entrance: "Cibodas", Ticket: 3,
		OrderId: "Order-101", BookingProductCores: nil, IdRanger: uint(1), GrossAmount: 120000, StatusBooking: "unpaid"}
	t.Run("Sukses Get Detail Booking", func(t *testing.T) {
		repo.On("GetID", mock.Anything).Return(returnRespon, nil).Once()
		srv := New(repo)
		res, err := srv.GetDetail(uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Detail Booking", func(t *testing.T) {
		repo.On("GetID", mock.Anything).Return(domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetDetail(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
