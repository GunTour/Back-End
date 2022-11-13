package repository

import (
	"GunTour/features/users/domain"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

// QUERY REGISTER USER
func (rq *repoQuery) Add(data domain.Core) (domain.Core, error) {
	var cnv User = FromCore(data)

	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on add user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

// QUERY UPDATE USER
func (rq *repoQuery) Edit(data domain.Core, id int) (domain.Core, error) {
	var cnv User = FromCore(data)

	err := rq.db.Where("id = ?", id).Updates(&cnv).Error
	if err != nil {
		return domain.Core{}, errors.New("no data")
	}

	if err := rq.db.First(&cnv, "id = ?", id).Error; err != nil {
		log.Error("error on getting after edit", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

// QUERY DELETE USER
func (rq *repoQuery) Remove(id int) (domain.Core, error) {
	var data User

	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on delete user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}

// QUERY LOGIN
func (rq *repoQuery) Login(input domain.Core) (domain.Core, error) {
	var data User

	err := rq.db.First(&data, "email = ?", input.Email)
	if err.RowsAffected == 0 {
		return domain.Core{}, errors.New("an invalid client request")
	}

	res := ToCore(data)
	return res, nil

}
