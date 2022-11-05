package services

import (
	"GunTour/features/ranger/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type rangerService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &rangerService{qry: repo}
}

func (rs *rangerService) Apply(data domain.Core) (domain.Core, error) {

	res, err := rs.qry.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

func (rs *rangerService) ShowAll() ([]domain.Core, error) {

	res, err := rs.qry.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil

}
