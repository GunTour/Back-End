package repository

import (
	"GunTour/features/admin/domain"

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

func (rq *repoQuery) GetPendaki() ([]domain.BookingCore, error) {
	var resQry []Booking
	if err := rq.db.Select("bookings.id_user", "users.full_name", "users.phone", "bookings.date_start", "bookings.date_end").
		Order("bookings.created_at desc").Joins("left join users on users.id = bookings.id_user").
		Find(&resQry).Scan(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainBooking(resQry)

	return res, nil
}

// func (rq *repoQuery) GetRanger(id uint) ([]UserCore, []UserCore, error){

// }

func (rq *repoQuery) GetBooking() ([]domain.BookingCore, error) {
	var resQry []Booking
	if err := rq.db.Select("bookings.id", "bookings.id_user", "users.full_name", "bookings.date_start", "bookings.date_end",
		"bookings.entrance", "bookings.ticket").
		Order("bookings.created_at desc").Joins("left join users on users.id = bookings.id_user").
		Find(&resQry).Scan(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainBooking(resQry)
	return res, nil
}

func (rq *repoQuery) GetProduct(page int) ([]domain.ProductCore, error) {
	var resQry []Product

	if page == 0 {
		if err := rq.db.Limit(20).Find(&resQry).Error; err != nil {
			return nil, err
		}
	} else {
		i := page * 20
		if err := rq.db.Offset(i).Limit(20).Find(&resQry).Scan(&resQry).Error; err != nil {
			return nil, err
		}
	}

	// selesai dari DB
	res := ToDomainProductArr(resQry)
	return res, nil
}

func (rq *repoQuery) InsertProduct(newProduct domain.ProductCore) (domain.ProductCore, error) {
	var res Product = FromDomainProduct(newProduct)
	if err := rq.db.Create(&res).Error; err != nil {
		return domain.ProductCore{}, err
	}

	newProduct = ToDomainProduct(res)
	return newProduct, nil
}

func (rq *repoQuery) UpdateProduct(newProduct domain.ProductCore) (domain.ProductCore, error) {
	var res Product = FromDomainProduct(newProduct)
	if err := rq.db.Where("id=?", newProduct.ID).Updates(&res).Error; err != nil {
		return domain.ProductCore{}, err
	}

	newProduct = ToDomainProduct(res)
	return newProduct, nil
}
func (rq *repoQuery) DeleteProduct(id int) error {
	if err := rq.db.Where("id=?", id).Delete(&Product{}).Error; err != nil {
		return err
	}
	return nil
}
