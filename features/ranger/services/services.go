package services

import (
	"GunTour/features/ranger/domain"
	"GunTour/utils/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"
)

type rangerService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &rangerService{qry: repo}
}

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

func (rs *rangerService) ShowAll(start time.Time, end time.Time) ([]domain.Core, error) {

	res, err := rs.qry.GetAll(start, end)
	if err != nil {
		log.Print(err.Error())
		if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
		return nil, errors.New("there is problem on server")
	}

	return res, nil

}
