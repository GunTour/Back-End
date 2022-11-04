package domain

import "time"

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
	DateStart   string
	DateEnd     string
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
	Update(data Core, id int) (Core, error)
	Delete(id int) (Core, error)
	Login(input Core) (Core, error)
}
