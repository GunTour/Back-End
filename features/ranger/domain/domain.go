package domain

import (
	"mime/multipart"
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

type Core struct {
	ID          uint
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Add(data Core, dataUser User) (Core, error)
	GetAll(start time.Time, end time.Time) ([]Core, error)
}

type Service interface {
	Apply(data Core, dataUser User, file multipart.File, fileheader *multipart.FileHeader) (Core, error)
	ShowAll(start time.Time, end time.Time) ([]Core, error)
}
