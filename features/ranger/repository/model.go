package repository

import (
	"GunTour/features/ranger/domain"

	"gorm.io/gorm"
)

type User struct {
	ID       uint
	FullName string
	Dob      string
	Address  string
	Phone    string
	Gender   string
	// Rangers  []Ranger `gorm:"foreignKey:UserID"`
}

type Ranger struct {
	gorm.Model
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}

func FromCore(rc domain.Core) Ranger {
	return Ranger{
		Model:       gorm.Model{ID: uint(rc.ID), CreatedAt: rc.CreatedAt, UpdatedAt: rc.UpdatedAt},
		Docs:        rc.Docs,
		Price:       rc.Price,
		Detail:      rc.Detail,
		Status:      rc.Status,
		StatusApply: rc.StatusApply,
		UserID:      uint(rc.UserID),
	}
}

func ToCore(r Ranger) domain.Core {
	return domain.Core{
		ID:          r.ID,
		Docs:        r.Docs,
		Price:       r.Price,
		Detail:      r.Detail,
		Status:      r.Status,
		StatusApply: r.StatusApply,
		UserID:      r.UserID,
		User:        domain.User(r.User),
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func ToCoreArray(ar []Ranger) []domain.Core {
	var arr []domain.Core
	for _, val := range ar {
		arr = append(arr, domain.Core{
			ID:          val.ID,
			Docs:        val.Docs,
			Price:       val.Price,
			Detail:      val.Detail,
			Status:      val.Status,
			StatusApply: val.StatusApply,
			UserID:      val.UserID,
			User:        domain.User(val.User),
		})
	}
	return arr
}
