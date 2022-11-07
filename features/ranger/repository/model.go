package repository

import (
	"GunTour/features/ranger/domain"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string
	Email       string `gorm:"unique"`
	Password    string
	Role        string
	Phone       string
	Address     string
	Dob         string
	Gender      string
	UserPicture string
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

type Booking struct {
	gorm.Model
	IdUser          uint
	DateStart       time.Time
	DateEnd         time.Time
	Entrance        string
	Ticket          int
	IdRanger        uint
	GrossAmount     int
	Token           string
	OrderId         string
	Link            string
	StatusBooking   string
	StatusPendakian string
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
