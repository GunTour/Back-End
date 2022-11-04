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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Add(data Core) (Core, error)
	Edit(data Core, id int) (Core, error)
	Remove(id int) (Core, error)
	Login(input Core) (Core, error)
}

type Service interface {
	Insert(data Core) (Core, error)
	Update(data Core, file multipart.File, fileheader *multipart.FileHeader, id int) (Core, error)
	Delete(id int) (Core, error)
	Login(input Core) (Core, error)
}
