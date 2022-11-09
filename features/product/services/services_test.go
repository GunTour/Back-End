package services

import (
	"GunTour/features/product/domain"
	mocks "GunTour/mocks/product"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShowAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	page := 1
	t.Run("success get all product", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{{ID: uint(1), ProductName: "ayam", RentPrice: 200, Detail: "ss", Note: "ss", ProductPicture: "ss"}},
			1, 1, nil).Once()
		srv := New(repo)
		res, pages, totalPage, err := srv.ShowAll(uint(page))
		assert.Nil(t, err)
		assert.NotEmpty(t, pages)
		assert.NotEmpty(t, totalPage)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("failed get all product", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{}, 0, 0, errors.New("no data")).Once()
		srv := New(repo)
		res, pages, totalPage, err := srv.ShowAll(0)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.Equal(t, 0, pages)
		assert.Equal(t, 0, totalPage)
		repo.AssertExpectations(t)
	})
	// repo := mocks.NewRepository(t)
	// page := 1
	// ret := []domain.Core{{ID: 1, ProductName: "tenda", RentPrice: 10000, ProductPicture: "https://guntour.s3.ap-southeast-1.amazonaws.com/posts/78Yawh7DexFdxltGofhy-1000_F_346839683_6nAPzbhpSkIpb8pmAwufkC7c5eD7wYws.jpg"}}

	// t.Run("success get all product", func(t *testing.T) {
	// 	repo.On("GetAll", mock.Anything).Return(ret, 1, 1, nil).Once()
	// 	srv := New(repo)
	// 	res, pages, totalPage, err := srv.ShowAll(uint(page))
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, pages)
	// 	assert.NotEmpty(t, totalPage)
	// 	assert.Equal(t, ret, res)
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("failed get all product", func(t *testing.T) {
	// 	repo.On("GetAll", mock.Anything).Return([]domain.Core{}, 0, 0, errors.New("no data")).Once()
	// 	srv := New(repo)
	// 	res, pages, totalPage, err := srv.ShowAll(uint(page))
	// 	assert.NotNil(t, err)
	// 	assert.Empty(t, res)
	// 	assert.Equal(t, 0, pages)
	// 	assert.Equal(t, 0, totalPage)
	// 	repo.AssertExpectations(t)
	// })
}

func TestShowByID(t *testing.T) {
	repo := mocks.NewRepository(t)
	id := 1
	ret := domain.Core{ID: 1, ProductName: "tenda", RentPrice: 10000, ProductPicture: "https://guntour.s3.ap-southeast-1.amazonaws.com/posts/78Yawh7DexFdxltGofhy-1000_F_346839683_6nAPzbhpSkIpb8pmAwufkC7c5eD7wYws.jpg"}

	t.Run("success get product", func(t *testing.T) {
		repo.On("GetByID", mock.Anything).Return(ret, nil).Once()
		srv := New(repo)
		res, err := srv.ShowByID(uint(id))
		assert.NoError(t, err)
		assert.Equal(t, ret, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed get product", func(t *testing.T) {
		repo.On("GetByID", mock.Anything).Return(domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.ShowByID(uint(id))
		assert.Nil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
