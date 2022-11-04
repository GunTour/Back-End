package repository

import (
	"GunTour/features/booking/domain"

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

func (rq *repoQuery) Get(idUser uint) ([]domain.Core, error) {
	var resQry []Booking
	if err := rq.db.Where("id_user=?", idUser).Find(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) GetID(idBooking uint) (domain.Core, error) {
	var resQry Booking
	if err := rq.db.Where("id=?", idBooking).Find(&resQry).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) GetRanger(idRanger uint) ([]domain.Core, error) {
	var resQry []Booking
	if err := rq.db.Where("id_ranger=?", idRanger).Find(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newBooking domain.Core) (domain.Core, error) {
	var cnv Booking = FromDomain(newBooking)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	if newBooking.BookingProductCores != nil {
		var productCnv = FromDomainProduct(newBooking.BookingProductCores, cnv.ID)
		err := rq.db.Create(&productCnv).Error
		if err != nil {
			return domain.Core{}, err
		}
	}
	// selesai dari DB
	newBooking = ToDomain(cnv)
	return newBooking, nil
}

func (rq *repoQuery) Update(newBooking domain.Core) (domain.Core, error) {
	var cnv Booking = FromDomain(newBooking)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newBooking = ToDomain(cnv)
	return newBooking, nil
}

func (rq *repoQuery) Delete(idBooking uint) error {
	if err := rq.db.Where("id = ?", idBooking).Delete(&Booking{}); err != nil {
		return err.Error
	}
	return nil
}
