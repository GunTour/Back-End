package domain

import (
	"mime/multipart"
	"time"
)

type User struct {
	ID       uint   `json:"id_user" form:"id_user"`
	FullName string `json:"fullname" form:"fullname"`
	Dob      string `json:"dob" form:"dob"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
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
	Add(data Core) (Core, error)
	GetAll(start time.Time, end time.Time) ([]Core, error)
}

type Service interface {
	Apply(data Core, file multipart.File, fileheader *multipart.FileHeader) (Core, error)
	ShowAll(start time.Time, end time.Time) ([]Core, error)
}
