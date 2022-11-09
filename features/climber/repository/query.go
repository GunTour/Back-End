package repository

import (
	"GunTour/features/climber/domain"
	"errors"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) GetClimber() (domain.Core, error) {
	var resQryClimber Climber

	result := rq.db.Order("created_at desc").First(&resQryClimber)
	if result.RowsAffected == 0 {
		return domain.Core{}, errors.New("not found")
	}

	// selesai dari DB
	resClimber := ToDomain(resQryClimber)

	return resClimber, nil
}
