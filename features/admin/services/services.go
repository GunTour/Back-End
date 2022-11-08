package services

import (
	"GunTour/features/admin/domain"
	"GunTour/utils/helper"
	"errors"
	"mime/multipart"
	"strings"
)

type adminService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &adminService{
		qry: repo,
	}
}

func (as *adminService) GetPendaki() ([]domain.BookingCore, domain.ClimberCore, error) {
	res, resClimb, err := as.qry.GetPendaki()
	if err != nil {
		return []domain.BookingCore{}, domain.ClimberCore{}, errors.New("no data")
	}

	return res, resClimb, nil

}

func (as *adminService) AddClimber(data domain.ClimberCore) (domain.ClimberCore, error) {
	res, err := as.qry.InsertClimber(data)
	if err != nil {
		return domain.ClimberCore{}, err
	}

	return res, nil
}

func (as *adminService) GetProduct(page int) ([]domain.ProductCore, int, int, error) {
	res, pages, totalPage, err := as.qry.GetProduct(page)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			return []domain.ProductCore{}, 0, 0, errors.New("page not found")
		}
		return []domain.ProductCore{}, 0, 0, errors.New("no data")
	}

	return res, pages, totalPage, nil
}

func (as *adminService) AddProduct(newProduct domain.ProductCore, file multipart.File, fileheader *multipart.FileHeader) (domain.ProductCore, error) {
	if fileheader != nil {
		res, _ := helper.UploadFile(file, fileheader)
		newProduct.ProductPicture = res
	}

	res, err := as.qry.InsertProduct(newProduct)
	if err != nil {
		return domain.ProductCore{}, errors.New("some problem on database")
	}

	return res, nil
}

func (as *adminService) EditProduct(newProduct domain.ProductCore, file multipart.File, fileheader *multipart.FileHeader) (domain.ProductCore, error) {
	if fileheader != nil {
		res, _ := helper.UploadFile(file, fileheader)
		newProduct.ProductPicture = res
	}

	res, err := as.qry.UpdateProduct(newProduct)
	if err != nil {
		return domain.ProductCore{}, errors.New("no data")
	}

	return res, nil
}

func (as *adminService) RemoveProduct(id int) error {
	err := as.qry.DeleteProduct(id)
	if err != nil {
		return errors.New("no data")
	}

	return nil
}

func (as *adminService) ShowAllRanger() ([]domain.RangerCore, []domain.RangerCore, error) {
	resAccepted, res, err := as.qry.GetAllRanger()
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return nil, nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, nil, errors.New("no data")
		}
		return nil, nil, errors.New("no data")
	}

	return resAccepted, res, nil
}

func (as *adminService) UpdateRanger(data domain.RangerCore, id uint) (domain.RangerCore, error) {
	res, err := as.qry.EditRanger(data, id)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.RangerCore{}, errors.New("rejected from database")
		}
		return domain.RangerCore{}, errors.New("some problem on database")
	}

	return res, nil
}
