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

// func (rq *repoQuery) GetProduct(page int) ([]ProductCore, error){

// }

// func (rq *repoQuery) InsertProduct(newProduct ProductCore) (ProductCore, error){

// }

// func (rq *repoQuery) UpdateProduct(newProduct ProductCore) (ProductCore, error){

// }
// func (rq *repoQuery) DeleteProduct(id int) error{

// }
