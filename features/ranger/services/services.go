package services

import (
	"GunTour/features/ranger/domain"
	"GunTour/utils/helper"
	"errors"
	"mime/multipart"
	"time"
)

type rangerService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &rangerService{qry: repo}
}

// SERVICE TO APPLY FORM RANGER
func (rs *rangerService) Apply(data domain.Core, dataUser domain.User, file multipart.File, fileheader *multipart.FileHeader) (domain.Core, error) {

	if fileheader != nil {
		res, _ := helper.UploadDocs(file, fileheader)
		data.Docs = res
	}

	data.Status = "off"
	data.StatusApply = "waiting"
	data.Price = 100000

	res, err := rs.qry.Add(data, dataUser)
	if err != nil {
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

// SERVICE TO SHOW RANGERS AVALAIBLE
func (rs *rangerService) ShowAll(start time.Time, end time.Time) ([]domain.Core, error) {

	res, err := rs.qry.GetAll(start, end)
	if err != nil {
		return nil, err
	}

	return res, nil

}
