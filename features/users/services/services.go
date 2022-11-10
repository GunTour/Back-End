package services

import (
	"GunTour/features/users/domain"
	"GunTour/utils/helper"
	"GunTour/utils/middlewares"
	"errors"
	"mime/multipart"

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
		return domain.Core{}, errors.New("cannot encrypt password")
	}
	data.Password = string(generate)

	data.Role = "pendaki"

	res, err := us.qry.Add(data)
	if err != nil {
		return domain.Core{}, err
	}
	return res, nil

}

func (us *userService) Update(data domain.Core, file multipart.File, fileheader *multipart.FileHeader, id int) (domain.Core, error) {

	if data.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return domain.Core{}, errors.New("cannot encrypt password")
		}
		data.Password = string(generate)
	}

	if fileheader != nil {
		res, _ := helper.UploadFile(file, fileheader)
		data.UserPicture = res
	}

	res, err := us.qry.Edit(data, id)
	if err != nil {
		return domain.Core{}, err
	}

	return res, nil

}

func (us *userService) Delete(id int) (domain.Core, error) {

	res, err := us.qry.Remove(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, nil

}

func (us *userService) Login(input domain.Core) (domain.Core, error) {

	res, err := us.qry.Login(input)
	if err != nil {
		return domain.Core{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
	if err != nil {
		return domain.Core{}, errors.New("password not match")
	}
	res.Token = middlewares.GenerateToken(res.ID, res.Role)

	return res, nil
}
