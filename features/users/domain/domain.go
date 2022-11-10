package domain

import (
	"mime/multipart"
	"time"
)

type Core struct {
	ID          int
	FullName    string
	Email       string
	Password    string
	Role        string
	Phone       string
	Address     string
	Dob         string
	Gender      string
	UserPicture string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserInfo struct {
	Email        string
	Fullname     string
	Photoprofile string
}

type ClimberCore struct {
	ID            uint
	IsClimber     int
	MaleClimber   int
	FemaleClimber int
}

type Repository interface {
	GetClimber() (ClimberCore, error)
	Add(data Core) (Core, error)
	Edit(data Core, id int) (Core, error)
	Remove(id int) (Core, error)
	Login(input Core) (Core, error)
}

type Service interface {
	ShowClimber() (ClimberCore, error)
	Insert(data Core) (Core, error)
	Update(data Core, file multipart.File, fileheader *multipart.FileHeader, id int) (Core, error)
	Delete(id int) (Core, error)
	Login(input Core) (Core, error)
}
