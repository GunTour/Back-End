package repository

import (
	"GunTour/features/admin/domain"
	"errors"

	"github.com/labstack/gommon/log"
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

func (rq *repoQuery) GetProduct(page int) ([]domain.ProductCore, int, int, error) {
	var resQry []Product
	var sum float64
	var totalPage int64

	if page == 0 || page == 1 {
		page = 1
		if err := rq.db.Limit(8).Find(&resQry).Error; err != nil {
			return nil, 0, 0, err
		}
	} else {
		i := (page - 1) * 8
		if err := rq.db.Offset(i).Limit(8).Find(&resQry).Scan(&resQry).Error; err != nil {
			return nil, 0, 0, err
		}
	}

	rq.db.Model(&Product{}).Count(&totalPage)
	sum = float64(totalPage) / 8
	if sum > float64((int(sum))) {
		totalPage = int64(sum) + 1
	} else if sum == float64((int(sum))) {
		totalPage = 1
	} else if sum > 0 {
		totalPage = 1
	}

	if page > int(totalPage) {
		return nil, 0, 0, errors.New("page not found")
	}

	// selesai dari DB
	res := ToDomainProductArr(resQry)
	return res, page, int(totalPage), nil
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
	err := rq.db.Where("id=?", id).Delete(&Product{})
	if err.RowsAffected == 0 {
		return errors.New("no data")
	}
	return nil
}

func (rq *repoQuery) GetAllRanger() ([]domain.RangerCore, []domain.RangerCore, error) {
	var data []Ranger
	var datas []Ranger

	if err := rq.db.Preload("User").Where("status_apply = ?", "accepted").Find(&data).Error; err != nil {
		log.Error("error on get all ranger", err.Error())
		return nil, nil, err
	}

	if err := rq.db.Preload("User").Where("status_apply = ?", "waiting").Find(&datas).Error; err != nil {
		log.Error("error on get all applied ranger", err.Error())
		return nil, nil, err
	}

	resAccepted := ToDomainRangerArray(data)
	res := ToDomainRangerArray(datas)
	return resAccepted, res, nil
}

func (rq *repoQuery) EditRanger(data domain.RangerCore, id uint) (domain.RangerCore, error) {
	var cnv Ranger = FromDomainRanger(data)

	if err := rq.db.Table("rangers").Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on edit status apply ranger", err.Error())
		return domain.RangerCore{}, err
	}

	if err := rq.db.Preload("User").Table("rangers").Where("id = ?", id).First(&cnv).Error; err != nil {
		log.Error("error on getting after edit", err.Error())
		return domain.RangerCore{}, err
	}

	if err := rq.db.Table("users").Where("id = ?", cnv.UserID).Update("role", "ranger").Error; err != nil {
		log.Error("error on edit role user", err.Error())
		return domain.RangerCore{}, err
	}

	res := ToDomainRanger(cnv)
	return res, nil
}
