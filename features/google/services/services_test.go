package services

import (
	"GunTour/features/google/domain"
	mocks "GunTour/mocks/google"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPesanCal(t *testing.T) {
	repo := new(mocks.Repository)
	returnRes := domain.BookingCore{ID: uint(1), IdRanger: uint(18), Email: "khalidrianda12@gmail.com"}

	t.Run("success get pesan", func(t *testing.T) {
		repo.On("GetPesanCal", mock.Anything).Return(returnRes).Once()

		usecase := New(repo)
		res := usecase.GetPesanCal()
		assert.Equal(t, returnRes, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed get pesan", func(t *testing.T) {
		repo.On("GetPesanCal", mock.Anything).Return(domain.BookingCore{}).Once()

		usecase := New(repo)
		res := usecase.GetPesanCal()
		assert.Equal(t, domain.BookingCore{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetPesan(t *testing.T) {
	repo := new(mocks.Repository)
	returnRes := domain.PesanCore{ID: uint(1), IdRanger: uint(18), Email: "khalidrianda12@gmail.com", Status: "accepted"}
	returnRespon := domain.RangerCore{ID: 1, UserID: uint(2)}

	t.Run("success get pesan", func(t *testing.T) {
		repo.On("GetPesan", mock.Anything).Return(returnRes, returnRespon).Once()

		usecase := New(repo)
		res, resRanger := usecase.GetPesan()
		assert.Equal(t, returnRes, res)
		assert.Equal(t, returnRespon, resRanger)
		repo.AssertExpectations(t)
	})

	t.Run("failed get pesan", func(t *testing.T) {
		repo.On("GetPesan", mock.Anything).Return(domain.PesanCore{}, domain.RangerCore{}).Once()

		usecase := New(repo)
		res, resRanger := usecase.GetPesan()
		assert.Equal(t, domain.PesanCore{}, res)
		assert.Equal(t, domain.RangerCore{}, resRanger)
		repo.AssertExpectations(t)
	})
}

func TestInsertCode(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.Code{ID: uint(1), Code: "adsasf"}

	t.Run("success get pesan", func(t *testing.T) {
		repo.On("InsertCode", mock.Anything).Return(nil).Once()

		usecase := New(repo)
		err := usecase.InsertCode(input)
		assert.Equal(t, nil, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get pesan", func(t *testing.T) {
		repo.On("InsertCode", mock.Anything).Return(errors.New("some problem on database")).Once()

		usecase := New(repo)
		err := usecase.InsertCode(domain.Code{})
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}
