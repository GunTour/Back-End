package repository

import (
	"GunTour/features/ranger/domain"

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

	var cnv Ranger = FromCore(data)

	rq.db.Where("user_id = ?", cnv.UserID).Delete(&Ranger{})

	if err := rq.db.Save(&cnv).Error; err != nil {
		log.Error("error on add ranger", err.Error())
		return domain.Core{}, err
	}

	if err := rq.db.Preload("User").First(&cnv, "id = ?", cnv.ID).Error; err != nil {
		log.Error("error on getting after add ranger", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil

}

func (rq *repoQuery) GetAll() ([]domain.Core, error) {

	var data []Ranger

	if err := rq.db.Preload("User").Find(&data).Error; err != nil {
		log.Error("error on get all ranger", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	log.Print(res)
	return res, nil

}
