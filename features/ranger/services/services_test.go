package services

import (
	"GunTour/features/ranger/domain"
	mocks "GunTour/mocks/ranger"
	"errors"
	"mime/multipart"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.Core{UserID: (uint(1)), Docs: "", Detail: "baik", Price: 120000}
	inputData := domain.User{Model: gorm.Model{ID: uint(1)}, FullName: "Ayam", Phone: "08226795", Dob: "BANDA", Gender: "PRIA", Address: "BANDA ACEH"}
	returnRespon := domain.Core{ID: 1, UserID: (uint(1)), Docs: "", Detail: "baik", Price: 120000, User: inputData}
	var file multipart.File
	var FileHeader *multipart.FileHeader

	t.Run("success apply ranger", func(t *testing.T) {
		repo.On("Add", mock.Anything, mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.Apply(input, inputData, file, FileHeader)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed apply ranger", func(t *testing.T) {
		repo.On("Add", mock.Anything, mock.Anything).Return(domain.Core{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.Apply(domain.Core{}, domain.User{}, file, FileHeader)
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetRanger(t *testing.T) {
	repo := mocks.NewRepository(t)
	Start, _ := time.Parse("2006-01-02", "2023-02-19T04:54:42.123+07:00")
	End, _ := time.Parse("2006-01-02", "2023-02-19T04:54:42.123+07:00")
	t.Run("Sukses Get Ranger Booking", func(t *testing.T) {
		repo.On("GetAll", mock.Anything, mock.Anything).Return([]domain.Core{{ID: uint(1), UserID: (uint(1)), Docs: "", Detail: "baik", Price: 120000}}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAll(Start, End)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Ranger Booking", func(t *testing.T) {
		repo.On("GetAll", mock.Anything, mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.ShowAll(Start, End)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Ranger Booking", func(t *testing.T) {
		repo.On("GetAll", mock.Anything, mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.ShowAll(Start, End)
		if strings.Contains(err.Error(), "found") {
			err = errors.New("no data")
		}
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
