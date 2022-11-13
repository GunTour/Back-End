package services

import (
	"GunTour/features/product/domain"
)

type productService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &productService{qry: repo}
}

// SERVICE TO SHOW ALL PRODUCT
func (ps *productService) ShowAll(page uint) ([]domain.Core, int, int, error) {

	res, pages, totalPage, err := ps.qry.GetAll(page)
	if err != nil {
		return []domain.Core{}, 0, 0, err
	}

	return res, pages, totalPage, nil

}

// SERVICE TO SHOW PRODUCT DETAIL
func (ps *productService) ShowByID(id uint) (domain.Core, error) {

	res, err := ps.qry.GetByID(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, nil

}
