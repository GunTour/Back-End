package repository

import (
	"GunTour/features/users"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Add(data users.Core) (users.Core, error) {

	var cnv User = FromCore(data)

	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on add user", err.Error())
		return users.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

func (rq *repoQuery) Edit(data users.Core, id int) (users.Core, error) {

	var cnv User = FromCore(data)

	if err := rq.db.Table("users").Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on edit user", err.Error())
		return users.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

func (rq *repoQuery) Remove(id int) (users.Core, error) {

	var data User

	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on delete user", err.Error())
		return users.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}

func (rq *repoQuery) GetAll() ([]users.Core, error) {

	var data []User

	if err := rq.db.Find(&data).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	return res, nil

}

func (rq *repoQuery) GetByID(id int) (users.Core, error) {

	var data User

	if err := rq.db.First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on get by id user", err.Error())
		return users.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}

func (rq *repoQuery) Login(input users.Core) (users.Core, error) {

	var data User

	if err := rq.db.First(&data, "email = ?", input.Email).Error; err != nil {
		log.Error("error on login user", err.Error())
		return users.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}
