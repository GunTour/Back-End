package repository

import (
	"GunTour/features/booking/domain"
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
	if err := rq.db.Select("bookings.id", "bookings.id_user", "users.full_name", "users.phone", "bookings.date_start", "bookings.date_end", "bookings.ticket").
		Order("bookings.created_at desc").Joins("left join users on users.id = bookings.id_user").
		Where("bookings.id_ranger = ?", int(idRanger)).Find(&resQry).Scan(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArrayRanger(resQry)
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

	if newBooking.BookingProductCores != nil {
		var productCnv = FromDomainProduct(newBooking.BookingProductCores, cnv.ID)
		for i := 0; i < len(productCnv); i++ {
			err := rq.db.Where("id_booking=? AND id_product=?", productCnv[i].ID, productCnv[i].IdProduct).Updates(&productCnv[i]).Error
			if err != nil {
				rq.db.Create(&productCnv[i])
			}
		}

	}
	// selesai dari DB
	newBooking = ToDomain(cnv)
	return newBooking, nil
}

func (rq *repoQuery) Delete(idBooking uint) error {
	log.Println(idBooking)
	if err := rq.db.Where("id_booking = ?", idBooking).Delete(&BookingProduct{}); err != nil {
		log.Println("ini pr: ", err.Error)
		return err.Error
	}
	if err := rq.db.Where("id = ?", idBooking).Delete(&Booking{}); err != nil {
		log.Println("ini booking: ", err.Error)
		return err.Error
	}
	return nil
}
