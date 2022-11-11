package repository

import (
	"GunTour/features/gmail/domain"
	"log"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) InsertCode(code string) error {
	var resQry Code
	resQry.Code = code
	rq.db.Create(&resQry)

	return nil
}

func (rq *repoQuery) UpdateCode(code domain.Code) error {
	var resQry Code = FromDomain(code)
	rq.db.Create(&resQry)

	return nil
}

func (rq *repoQuery) GetCode() (domain.Code, error) {
	var resQry Code
	if err := rq.db.Order("created_at desc").First(&resQry).Error; err != nil {
		return domain.Code{}, err
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) GetPesan() (domain.PesanCore, domain.RangerCore) {
	var resQry Pesan
	var resQryRanger Ranger
	if err := rq.db.Order("created_at desc").First(&resQry).Error; err != nil {
		return domain.PesanCore{}, domain.RangerCore{}
	}
	if err := rq.db.Where("id=?", resQry.ID).First(&resQryRanger).Error; err != nil {
		return domain.PesanCore{}, domain.RangerCore{}
	}
	res := ToDomainPesan(resQry)
	resRanger := ToDomainRanger(resQryRanger)
	return res, resRanger
}

func (rq *repoQuery) GetPesanCal() domain.BookingCore {
	var resQry Booking
	if err := rq.db.Order("created_at desc").First(&resQry).Error; err != nil {
		return domain.BookingCore{}
	}

	if err := rq.db.Select("bookings.date_start", "bookings.date_end", "users.email").Order("bookings.created_at desc").Joins("left join users on users.id=bookings.id_user").Where("bookings.id_user = ?", resQry.IdUser).First(&resQry).Scan(&resQry).Error; err != nil {
		log.Print(resQry)
		return domain.BookingCore{}
	}

	res := ToDomainBooking(resQry)
	return res
}
