package repository

import (
	"GunTour/features/users/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Add(data domain.Core) (domain.Core, error) {

	var cnv User = FromCore(data)

	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on add user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

func (rq *repoQuery) Edit(data domain.Core, id int) (domain.Core, error) {

	var cnv User = FromCore(data)

	if err := rq.db.Table("users").Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on edit user", err.Error())
		return domain.Core{}, err
	}

	if err := rq.db.First(&cnv, "id = ?", id).Error; err != nil {
		log.Error("error on getting after edit", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

func (rq *repoQuery) Remove(id int) (domain.Core, error) {

	var data User

	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on delete user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}

func (rq *repoQuery) GetAll() ([]domain.Core, error) {

	var data []User

	if err := rq.db.Find(&data).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	return res, nil

}

func (rq *repoQuery) Login(input domain.Core) (domain.Core, error) {

	var data User

	if err := rq.db.First(&data, "email = ?", input.Email).Error; err != nil {
		log.Error("error on login user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}
