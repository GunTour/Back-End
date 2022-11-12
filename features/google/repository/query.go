package repository

import (
	"GunTour/features/google/domain"

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

func (rq *repoQuery) InsertCode(code domain.Code) error {
	var resQry Code = FromDomain(code)
	rq.db.Create(&resQry)

	return nil
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
	var resProductQry []BookingProduct

	// selesai dari DB
	if err := rq.db.Order("created_at desc").First(&resQry).Error; err != nil {
		return domain.BookingCore{}
	}

	if err := rq.db.Select("bookings.date_start", "bookings.date_end", "users.email").Order("bookings.created_at desc").Joins("left join users on users.id=bookings.id_user").Where("bookings.id_user = ?", resQry.IdUser).First(&resQry).Scan(&resQry).Error; err != nil {
		return domain.BookingCore{}
	}

	rq.db.Select("booking_products.id", "booking_products.id_booking", "booking_products.id_product",
		"booking_products.product_qty", "products.product_name", "products.rent_price").
		Order("booking_products.created_at desc").Joins("left join products on products.id = booking_products.id_product").
		Where("id_booking=?", resQry.ID).Find(&resProductQry).Scan(&resProductQry)

	res := ToDomainCore(resQry, resProductQry)

	return res
}
