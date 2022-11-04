package repository

import (
	"GunTour/features/ranger/domain"

	"gorm.io/gorm"
)

type User struct {
	ID       int
	FullName string
	Dob      string
	Address  string
	Phone    string
	Gender   string
}

type Ranger struct {
	gorm.Model
	UserID int
	User   User
	Docs   string
	Price  int
	Detail string
}

func FromCore(rc domain.Core) Ranger {
	return Ranger{
		Model:  gorm.Model{ID: uint(rc.ID), CreatedAt: rc.CreatedAt, UpdatedAt: rc.UpdatedAt},
		UserID: rc.UserID,
		User:   User(rc.User),
		Docs:   rc.Docs,
		Price:  rc.Price,
		Detail: rc.Detail,
	}
}

func ToCore(r Ranger) domain.Core {
	return domain.Core{
		ID:        int(r.ID),
		UserID:    r.UserID,
		User:      domain.User(r.User),
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
			UserID:    val.UserID,
			User:      domain.User(val.User),
			Docs:      val.Docs,
			Price:     val.Price,
			Detail:    val.Detail,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return arr
}
