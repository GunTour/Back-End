package services

import (
	"GunTour/features/climber/domain"
	"errors"
)

type climberService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &climberService{
		qry: repo,
	}
}

func (as *climberService) ShowClimber() (domain.Core, error) {
	res, err := as.qry.GetClimber()
	if err != nil {
		return domain.Core{}, errors.New("no data")
	}

	return res, nil
}
