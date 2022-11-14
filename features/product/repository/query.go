package repository

import (
	"GunTour/features/product/domain"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

// QUERY TO GET ALL PRODUCT DATA
func (rq *repoQuery) GetAll(page uint) ([]domain.Core, int, int, error) {
	var resQry []Product
	var sum float64
	var totalPage int64

	if page == 0 || page == 1 {
		page = 1
		if err := rq.db.Order("created_at desc").Limit(8).Find(&resQry).Error; err != nil {
			return nil, 0, 0, errors.New("no data")
		}
	} else {
		i := (page - 1) * 8
		if err := rq.db.Order("created_at desc").Offset(int(i)).Limit(8).Find(&resQry).Scan(&resQry).Error; err != nil {
			return nil, 0, 0, errors.New("no data")
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

	if page > uint(totalPage) {
		return nil, 0, 0, errors.New("page not found")
	}

	// selesai dari DB
	res := ToCoreArray(resQry)
	return res, int(page), int(totalPage), nil
}

// QUERY TO GET PRODUCT DETAIL
func (rq *repoQuery) GetByID(id uint) (domain.Core, error) {

	var data Product

	if err := rq.db.First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on login user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}
