package repository

import (
	"GunTour/features/product/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) GetAll(page uint) ([]domain.Core, uint, uint, error) {
	var resQry []Product
	var sum float64
	var totalPage int64

	if page == 0 || page == 1 {
		page = 1
		if err := rq.db.Limit(10).Find(&resQry).Error; err != nil {
			return nil, 0, 0, err
		}
	} else {
		i := (page - 1) * 10
		if err := rq.db.Offset(int(i)).Limit(10).Find(&resQry).Scan(&resQry).Error; err != nil {
			return nil, 0, 0, err
		}
		page += 1
	}

	rq.db.Count(&totalPage)
	sum = float64(totalPage) / 10
	if sum > float64((int(sum))) {
		totalPage = int64(sum) + 1
	} else if sum == float64((int(sum))) {
		totalPage = 1
	} else if sum > 0 {
		totalPage = 1
	}

	// selesai dari DB
	res := ToCoreArray(resQry)
	return res, page, uint(totalPage), nil
}

func (rq *repoQuery) GetByID(id uint) (domain.Core, error) {

	var data Product

	if err := rq.db.First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on login user", err.Error())
		return domain.Core{}, err
	}

	res := ToCore(data)
	return res, nil

}
