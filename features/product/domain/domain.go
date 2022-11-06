package domain

import "time"

type Core struct {
	ID             uint
	ProductName    string
	Price          uint
	Detail         string
	Note           string
	ProductPicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Paginate struct {
	Page  uint
	Limit uint
}

type Repository interface {
	GetAll(page uint) ([]Core, uint, uint, error)
	GetByID(id uint) (Core, error)
}

type Service interface {
	ShowAll(page uint) ([]Core, uint, uint, error)
	ShowByID(id uint) (Core, error)
}
