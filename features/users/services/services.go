package services

import (
	"GunTour/features/users/domain"
	"GunTour/utils/helper"
	"GunTour/utils/middlewares"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{qry: repo}
}

func (us *userService) Insert(data domain.Core) (domain.Core, error) {

	generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt password insert user", err.Error())
		return domain.Core{}, errors.New("cannot encrypt password")
	}
	data.Password = string(generate)

	data.Role = "pendaki"

	res, err := us.qry.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

func (us *userService) Update(data domain.Core, file multipart.File, fileheader *multipart.FileHeader, id int) (domain.Core, error) {

	if data.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("error on bcrypt password update user", err.Error())
			return domain.Core{}, errors.New("cannot encrypt password")
		}
		data.Password = string(generate)
	}

	if fileheader != nil {
		res, err := helper.UploadFile(file, fileheader)
		if err != nil {
			return domain.Core{}, err
		}
		data.UserPicture = res
	}

	res, err := us.qry.Edit(data, id)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil

}

func (us *userService) Delete(id int) (domain.Core, error) {

	res, err := us.qry.Remove(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	return res, nil

}

func (us *userService) Login(input domain.Core) (domain.Core, error) {

	res, err := us.qry.Login(input)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "an invalid") {
			return domain.Core{}, errors.New("an invalid client request")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
	if err != nil {
		return domain.Core{}, errors.New("an invalid client request")
	}
	res.Token = middlewares.GenerateToken(res.ID, res.Role)

	return res, nil
}

func (us *userService) ShowClimber() (domain.ClimberCore, error) {
	res, err := us.qry.GetClimber()
	if err != nil {
		return domain.ClimberCore{}, err
	}

	return res, nil
}
