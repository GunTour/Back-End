package services

import (
	"GunTour/features/climber/domain"
	mocks "GunTour/mocks/climber"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPendaki(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Show Climber", func(t *testing.T) {
		repo.On("GetClimber", mock.Anything).Return(domain.Core{IsClimber: 100, FemaleClimber: 99, MaleClimber: 1}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowClimber()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Show Climber", func(t *testing.T) {
		repo.On("GetClimber", mock.Anything).Return(domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.ShowClimber()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
