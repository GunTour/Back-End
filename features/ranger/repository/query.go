package repository

import (
	"GunTour/features/ranger/domain"
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

// QUERY TO INSERT DATA RANGER
func (rq *repoQuery) Add(data domain.Core, dataUser domain.User) (domain.Core, error) {

	var cnv Ranger = FromCore(data)
	if err := rq.db.Where("id=?", dataUser.ID).Updates(&dataUser).Error; err != nil {
		log.Error("error on updates users", err.Error())
		return domain.Core{}, err
	}

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

// QUERY TO GET RANGER DATA
func (rq *repoQuery) GetAll(start time.Time, end time.Time) ([]domain.Core, error) {

	var data []Ranger
	var idRanger []uint
	log.Print(start, end)
	rq.db.Model(&Booking{}).Where("date_start BETWEEN ? AND ? OR date_end BETWEEN ? AND ?",
		start, end, start, end).Distinct("id_ranger").Select("id_ranger").Find(&idRanger)

	err := rq.db.Not(&idRanger).Not("status = 'off' OR status = 'duty'").Preload("User").Find(&data)
	if err.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	res := ToCoreArray(data)

	return res, nil

}
