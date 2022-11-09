package domain

import "time"

type Core struct {
	ID             uint
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Repository interface {
	GetAll(page uint) ([]Core, int, int, error)
	GetByID(id uint) (Core, error)
}

type Service interface {
	ShowAll(page uint) ([]Core, int, int, error)
	ShowByID(id uint) (Core, error)
}
