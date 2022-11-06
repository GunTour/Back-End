package services

import (
	"GunTour/features/product/domain"
	"errors"
	"strings"
)

type productService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &productService{qry: repo}
}

func (ps *productService) ShowAll(page uint) ([]domain.Core, uint, uint, error) {

	res, pages, totalPage, err := ps.qry.GetAll(page)
	if err != nil {
		return []domain.Core{}, 0, 0, errors.New("no data")
	}

	return res, pages, totalPage, nil

}

func (ps *productService) ShowByID(id uint) (domain.Core, error) {

	res, err := ps.qry.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil

}
