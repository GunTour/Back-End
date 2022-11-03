package user

import "time"

type Core struct {
	ID          int
	FullName    string
	Email       string
	Password    string
	Role        string
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
	GetAll() ([]Core, error)
	GetByID(id int) (Core, error)
}

type Service interface {
	Insert(data Core) (Core, error)
	Update(data Core, id int) (Core, error)
	Delete(id int) (Core, error)
	ShowAll() ([]Core, error)
	ShowByID(id int) (Core, error)
}
