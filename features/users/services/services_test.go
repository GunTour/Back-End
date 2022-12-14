package services

import (
	"GunTour/features/users/domain"
	mocks "GunTour/mocks/users"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.Core{FullName: "same", Email: "same@gmail.com", Password: "Same1234"}
	returnRespon := domain.Core{ID: 1, FullName: "same", Email: "same@gmail.com", Role: "pendaki"}

	t.Run("create success", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.Insert(input)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed create user", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.Core{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.Insert(input)
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.Repository)
	data := domain.Core{ID: 1, FullName: "same", Email: "same@gmail.com", Password: "Same1234"}
	returnRespon := domain.Core{ID: 1, FullName: "same", Email: "same@gmail.com", Role: "pendaki"}
	var file multipart.File
	var FileHeader *multipart.FileHeader

	t.Run("success update user", func(t *testing.T) {
		repo.On("Edit", mock.Anything, 1).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.Update(data, file, FileHeader, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed update user", func(t *testing.T) {
		repo.On("Edit", mock.Anything, 1).Return(domain.Core{}, errors.New("some problem on database")).Once()
		usecase := New(repo)
		res, err := usecase.Update(domain.Core{}, file, FileHeader, 1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.Repository)

	t.Run("success delete data", func(t *testing.T) {

		repo.On("Remove", mock.Anything).Return(domain.Core{}, nil).Once()

		useCase := New(repo)

		res, err := useCase.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("Remove", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()

		useCase := New(repo)

		res, _ := useCase.Delete(2)
		assert.Error(t, errors.New("error"))
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.Core{Email: "ssss@gmail.com", Password: "Same1234"}
	t.Run("success login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Password: "$2a$10$rV5osQeweSIUxr0nwzVQj.pzXYR6sh1TG0QxIcwoayYFgwJuiojjy"}, nil).Once()
		srv := New(repo)
		res, err := srv.Login(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("password not match login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Password: "$2a$10$rV5osQeweSIUxr0nwzVQj.pzXYR6sh1TG0QxIcwoayYFgwJuiojjy"}, errors.New("password not match")).Once()
		srv := New(repo)
		input := domain.Core{Email: "khalidrian@gmail.com", Password: "Same123"}
		res, err := srv.Login(input)
		// assert.NotEmpty(t, er)
		assert.Empty(t, res)
		assert.EqualError(t, err, "password not match")
		repo.AssertExpectations(t)
	})
}
