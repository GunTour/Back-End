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
	var resProductQry []BookingProduct
	if err := rq.db.Select("bookings.id", "bookings.id_ranger", "users.full_name", "bookings.entrance", "bookings.date_start", "bookings.date_end",
		"bookings.ticket", "bookings.gross_amount", "bookings.order_id", "bookings.link", "bookings.status_booking").
		Order("bookings.created_at desc").Joins("left join users on users.id = bookings.id_ranger").
		Where("bookings.id = ?", int(idBooking)).Find(&resQry).Scan(&resQry).Error; err != nil {
		return domain.Core{}, err
	}

	if err := rq.db.Select("booking_products.id", "booking_products.id_booking", "booking_products.id_product",
		"booking_products.product_qty", "products.product_name", "products.rent_price").
		Order("booking_products.created_at desc").Joins("left join products on products.id = booking_products.id_product").
		Where("id_booking=?", idBooking).Find(&resProductQry).Scan(&resProductQry).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	res := ToDomainCore(resQry, resProductQry)
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
	var productCnv []BookingProduct
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	if newBooking.BookingProductCores != nil {
		productCnv = FromDomainProduct(newBooking.BookingProductCores, cnv.ID)
		err := rq.db.Create(&productCnv).Error
		if err != nil {
			return domain.Core{}, err
		}
		if err := rq.db.Select("booking_products.id", "booking_products.id_booking", "booking_products.id_product",
			"booking_products.product_qty", "products.product_name", "products.rent_price").
			Order("booking_products.created_at desc").Joins("left join products on products.id = booking_products.id_product").
			Where("id_booking=?", cnv.ID).Find(&productCnv).Scan(&productCnv).Error; err != nil {
			return domain.Core{}, err
		}
	}
	// selesai dari DB
	newBooking = ToDomainCore(cnv, productCnv)
	return newBooking, nil
}

func (rq *repoQuery) Update(newBooking domain.Core) (domain.Core, error) {
	var cnv Booking = FromDomain(newBooking)
	var productCnv []BookingProduct
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	rq.db.Where("id=?", cnv.ID).Find(&cnv)

	if newBooking.BookingProductCores != nil {
		productCnv = FromDomainProduct(newBooking.BookingProductCores, cnv.ID)
		rq.db.Where("id_booking=?", cnv.ID).Delete(&BookingProduct{})
		err := rq.db.Create(&productCnv).Error
		if err != nil {
			return domain.Core{}, err
		}
		if err := rq.db.Select("booking_products.id", "booking_products.id_booking", "booking_products.id_product",
			"booking_products.product_qty", "products.product_name", "products.rent_price").
			Order("booking_products.created_at desc").Joins("left join products on products.id = booking_products.id_product").
			Where("id_booking=?", cnv.ID).Find(&productCnv).Scan(&productCnv).Error; err != nil {
			return domain.Core{}, err
		}
	}
	// selesai dari DB
	newBooking = ToDomainCore(cnv, productCnv)
	return newBooking, nil
}

func (rq *repoQuery) Delete(idBooking uint) error {
	if err := rq.db.Where("id_booking = ?", idBooking).Delete(&BookingProduct{}); err != nil {
		return err.Error
	}
	if err := rq.db.Where("id = ?", idBooking).Delete(&Booking{}); err != nil {
		return err.Error
	}
	return nil
}

func (rq *repoQuery) UpdateMidtrans(newBooking domain.Core) error {
	var cnv Booking = FromDomain(newBooking)
	if err := rq.db.Where("order_id = ?", cnv.OrderId).Updates(&cnv).Error; err != nil {
		return err
	}

	return nil
}
