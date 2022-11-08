package services

import (
	"GunTour/features/admin/domain"
	mocks "GunTour/mocks/admin"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertProduct(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.ProductCore{ProductName: "Tenda", RentPrice: 2000, Detail: "ss", Note: "ss", ProductPicture: "sss"}
	returnRespon := domain.ProductCore{ID: 1, ProductName: "Tenda", RentPrice: 2000, Detail: "ss", Note: "ss", ProductPicture: "sss"}
	var file multipart.File
	var FileHeader *multipart.FileHeader

	t.Run("create success", func(t *testing.T) {
		repo.On("InsertProduct", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.AddProduct(input, file, FileHeader)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed create product", func(t *testing.T) {
		repo.On("InsertProduct", mock.Anything).Return(domain.ProductCore{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.AddProduct(domain.ProductCore{}, file, FileHeader)
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.ProductCore{}, res)
		repo.AssertExpectations(t)
	})

}

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.Repository)
	var file multipart.File
	var FileHeader *multipart.FileHeader
	input := domain.ProductCore{ID: uint(1), ProductName: "Tenda", RentPrice: 2000, Detail: "ss", Note: "ss", ProductPicture: "sss"}
	returnRespon := domain.ProductCore{ID: 1, ProductName: "Tenda", RentPrice: 2000, Detail: "ss", Note: "ss", ProductPicture: "sss"}

	t.Run("success update booking", func(t *testing.T) {
		repo.On("UpdateProduct", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.EditProduct(input, file, FileHeader)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed update booking", func(t *testing.T) {
		repo.On("UpdateProduct", mock.Anything).Return(domain.ProductCore{}, errors.New("some problem on database")).Once()
		usecase := New(repo)

		res, err := usecase.EditProduct(domain.ProductCore{}, file, FileHeader)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.Repository)

	t.Run("success delete data", func(t *testing.T) {

		repo.On("DeleteProduct", mock.Anything).Return(nil).Once()

		useCase := New(repo)

		err := useCase.RemoveProduct(1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("DeleteProduct", mock.Anything).Return(errors.New("error")).Once()

		useCase := New(repo)

		err := useCase.RemoveProduct(0)
		assert.Error(t, errors.New("error"))
		assert.Equal(t, err, err)
		repo.AssertExpectations(t)
	})
}

func TestGetPendaki(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get All Pendaki", func(t *testing.T) {
		repo.On("GetPendaki", mock.Anything).Return([]domain.BookingCore{{ID: uint(1), GrossAmount: 120000, OrderId: "Order-101", Link: "www.mid", StatusBooking: "unpaid"}},
			domain.ClimberCore{IsClimber: 100, FemaleClimber: 99, MaleClimber: 1}, nil).Once()
		srv := New(repo)
		res, resClimb, err := srv.GetPendaki()
		assert.Nil(t, err)
		assert.NotEmpty(t, resClimb)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Booking", func(t *testing.T) {
		repo.On("GetPendaki", mock.Anything).Return([]domain.BookingCore{}, domain.ClimberCore{}, errors.New("no data")).Once()
		srv := New(repo)
		res, resClimb, err := srv.GetPendaki()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.Empty(t, resClimb)
		repo.AssertExpectations(t)
	})
}

func TestAddClimber(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.ClimberCore{IsClimber: 100, FemaleClimber: 99, MaleClimber: 1}
	returnRespon := domain.ClimberCore{ID: uint(1), IsClimber: 100, FemaleClimber: 99, MaleClimber: 1}

	t.Run("add climber success", func(t *testing.T) {
		repo.On("InsertClimber", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.AddClimber(input)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed add climber", func(t *testing.T) {
		repo.On("InsertClimber", mock.Anything).Return(domain.ClimberCore{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.AddClimber(domain.ClimberCore{})
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.ClimberCore{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetRanger(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnRespon := []domain.RangerCore{{ID: uint(1), Docs: "sss", Price: 200}}
	returnRespon2 := []domain.RangerCore{{ID: uint(1), Docs: "sss", Price: 200}}
	t.Run("Sukses Get Ranger", func(t *testing.T) {
		repo.On("GetAllRanger", mock.Anything).Return(returnRespon, returnRespon2, nil).Once()
		srv := New(repo)
		res, resApply, err := srv.ShowAllRanger()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.NotEmpty(t, resApply)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Ranger Booking", func(t *testing.T) {
		repo.On("GetAllRanger", mock.Anything).Return([]domain.RangerCore{}, []domain.RangerCore{}, errors.New("no data")).Once()
		srv := New(repo)
		res, resApply, err := srv.ShowAllRanger()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.Empty(t, resApply)
		repo.AssertExpectations(t)
	})
}

func TestGetProduct(t *testing.T) {
	repo := mocks.NewRepository(t)
	page := 1
	t.Run("Sukses Get Product", func(t *testing.T) {
		repo.On("GetProduct", mock.Anything).Return([]domain.ProductCore{{ID: uint(1), ProductName: "ayam", RentPrice: 200, Detail: "ss", Note: "ss", ProductPicture: "ss"}},
			1, 1, nil).Once()
		srv := New(repo)
		res, pages, totalPage, err := srv.GetProduct(page)
		assert.Nil(t, err)
		assert.NotEmpty(t, pages)
		assert.NotEmpty(t, totalPage)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Booking", func(t *testing.T) {
		repo.On("GetProduct", mock.Anything).Return([]domain.ProductCore{}, 0, 0, errors.New("no data")).Once()
		srv := New(repo)
		res, pages, totalPage, err := srv.GetProduct(0)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.Equal(t, 0, pages)
		assert.Equal(t, 0, totalPage)
		repo.AssertExpectations(t)
	})
}

func TestUpdateRanger(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnRespon := domain.RangerCore{ID: uint(1), UserID: uint(1), StatusApply: "accepted"}
	returnRespon2 := domain.RangerCore{ID: uint(1), UserID: uint(1), StatusApply: "accepted"}
	t.Run("Sukses Update Ranger Status", func(t *testing.T) {
		repo.On("EditRanger", mock.Anything, mock.Anything).Return(returnRespon2, nil).Once()
		srv := New(repo)
		res, err := srv.UpdateRanger(returnRespon, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update Ranger Status", func(t *testing.T) {
		repo.On("EditRanger", mock.Anything, mock.Anything).Return(domain.RangerCore{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.UpdateRanger(domain.RangerCore{}, 1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
