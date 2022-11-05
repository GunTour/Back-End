package repository

import (
	"GunTour/features/ranger/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Dob      string
	Address  string
	Phone    string
	Gender   string
	RangerID int
}

type Ranger struct {
	gorm.Model
	User   User `gorm:"Foreignkey:RangerID;association_foreignkey:ID;"`
	Docs   string
	Price  int
	Detail string
}

func FromCore(rc domain.Core) Ranger {
	return Ranger{
		Model:  gorm.Model{ID: uint(rc.ID), CreatedAt: rc.CreatedAt, UpdatedAt: rc.UpdatedAt},
		User:   User{},
		Docs:   rc.Docs,
		Price:  rc.Price,
		Detail: rc.Detail,
	}
}

func ToCore(r Ranger) domain.Core {
	return domain.Core{
		ID:        int(r.ID),
		User:      domain.User{},
		Docs:      r.Docs,
		Price:     r.Price,
		Detail:    r.Detail,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ToCoreArray(ar []Ranger) []domain.Core {
	var arr []domain.Core
	for _, val := range ar {
		arr = append(arr, domain.Core{
			ID:        int(val.ID),
			User:      domain.User{},
			Docs:      val.Docs,
			Price:     val.Price,
			Detail:    val.Detail,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return arr
}
