package domain

import "time"

type User struct {
	ID       int
	FullName string
	Dob      string
	Address  string
	Phone    string
	Gender   string
	RangerID int
}

type Core struct {
	ID        int
	User      User
	Docs      string
	Price     int
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Add(data Core) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	Apply(data Core) (Core, error)
	ShowAll() ([]Core, error)
}
