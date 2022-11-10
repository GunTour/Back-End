package repository

import (
	"GunTour/features/users/domain"
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
	Token       string    `gorm:"-:migration;<-:false"`
	Bookings    []Booking `gorm:"foreignKey:IdUser"`
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
	FullName        string           `gorm:"-:migration;<-:false"`
	Phone           string           `gorm:"-:migration;<-:false"`
	BookingProducts []BookingProduct `gorm:"foreignKey:IdBooking"`
}

type Product struct {
	gorm.Model
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
	Products       []BookingProduct `gorm:"foreignKey:IdProduct"`
}

type BookingProduct struct {
	gorm.Model
	IdBooking   uint
	IdProduct   uint
	ProductQty  int
	ProductName string `gorm:"-:migration;<-:false"`
	RentPrice   int    `gorm:"-:migration;<-:false"`
}

type UserInfo struct {
	Email        string
	Fullname     string
	Photoprofile string
}

func FromCore(uc domain.Core) User {
	return User{
		Model:       gorm.Model{ID: uint(uc.ID), CreatedAt: uc.CreatedAt, UpdatedAt: uc.UpdatedAt},
		FullName:    uc.FullName,
		Email:       uc.Email,
		Password:    uc.Password,
		Role:        uc.Role,
		Phone:       uc.Phone,
		Address:     uc.Address,
		Dob:         uc.Dob,
		Gender:      uc.Gender,
		UserPicture: uc.UserPicture,
		Token:       uc.Token,
	}
}

func ToCore(u User) domain.Core {
	return domain.Core{
		ID:          int(u.ID),
		FullName:    u.FullName,
		Email:       u.Email,
		Password:    u.Password,
		Role:        u.Role,
		Phone:       u.Phone,
		Address:     u.Address,
		Dob:         u.Dob,
		Gender:      u.Gender,
		UserPicture: u.UserPicture,
		Token:       u.Token,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func ToCoreArray(ua []User) []domain.Core {
	var arr []domain.Core
	for _, val := range ua {
		arr = append(arr, domain.Core{
			ID:          int(val.ID),
			FullName:    val.FullName,
			Email:       val.Email,
			Password:    val.Password,
			Role:        val.Role,
			Phone:       val.Phone,
			Address:     val.Address,
			Dob:         val.Dob,
			Gender:      val.Gender,
			UserPicture: val.UserPicture,
			Token:       val.Token,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		})
	}
	return arr
}

func (u *UserInfo) ToModelUserInfo() domain.UserInfo {
	return domain.UserInfo{
		Email:        u.Email,
		Fullname:     u.Fullname,
		Photoprofile: u.Photoprofile,
	}
}
