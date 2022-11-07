package services

import (
	"GunTour/features/ranger/domain"
	"GunTour/utils/helper"
	"errors"
	"mime/multipart"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

type rangerService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &rangerService{qry: repo}
}

func (rs *rangerService) Apply(data domain.Core, file multipart.File, fileheader *multipart.FileHeader) (domain.Core, error) {

	if fileheader != nil {
		res, err := helper.UploadDocs(file, fileheader)
		if err != nil {
			return domain.Core{}, errors.New("error on upload docs")
		}
		data.Docs = res
	}

	data.Status = "off"
	data.StatusApply = "waiting"

	res, err := rs.qry.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

func (rs *rangerService) ShowAll(start time.Time, end time.Time) ([]domain.Core, error) {

	res, err := rs.qry.GetAll(start, end)
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
