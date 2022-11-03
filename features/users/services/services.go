package services

import (
	"GunTour/features/users"
	"GunTour/utils/middlewares"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry users.Repository
}

func New(repo users.Repository) users.Service {
	return &userService{qry: repo}
}

func (us *userService) Insert(data users.Core) (users.Core, error) {

	generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt password insert user", err.Error())
		return users.Core{}, errors.New("cannot encrypt password")
	}
	data.Password = string(generate)

	res, err := us.qry.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return users.Core{}, errors.New("rejected from database")
		}
		return users.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

func (us *userService) Update(data users.Core, id int) (users.Core, error) {

	if data.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("error on bcrypt password update user", err.Error())
			return users.Core{}, errors.New("cannot encrypt password")
		}
		data.Password = string(generate)
	}

	res, err := us.qry.Edit(data, id)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return users.Core{}, errors.New("rejected from database")
		}
		return users.Core{}, errors.New("some problem on database")
	}
	return res, nil

}

func (us *userService) Delete(id int) (users.Core, error) {

	res, err := us.qry.Remove(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return users.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return users.Core{}, errors.New("no data")
		}
	}
	return res, nil

}

func (us *userService) ShowAll() ([]users.Core, error) {

	res, err := us.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
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

func (us *userService) ShowByID(id int) (users.Core, error) {

	res, err := us.qry.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return users.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return users.Core{}, errors.New("no data")
		}
	}
	return res, nil

}

func (us *userService) Login(input users.Core) (users.Core, string, error) {

	res, err := us.qry.Login(input)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return users.Core{}, "", errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return users.Core{}, "", errors.New("no data")
		}
	}

	token := middlewares.GenerateToken(res.ID, res.Role)

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
	if err != nil {
		return users.Core{}, "", errors.New("password not match")
	}

	return res, token, nil
}
