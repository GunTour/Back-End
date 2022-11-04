package domain

import "time"

type Core struct {
	ID              uint
	IdUser          uint
	DateStart       time.Time
	DateEnd         time.Time
	Entrance        string
	Ticket          int
	IdRanger        uint
	GrossAmount     int
	Token           string
	Link            string
	StatusBooking   string
	StatusPendakian string
}

type BookingProductCore struct {
	ID        uint
	IdBooking uint
	IdProduct uint
}

type Repository interface {
	Get() ([]Core, error)
	GetID(idBooking uint) (Core, error)
	GetRanger(id_ranger uint) ([]Core, error)
	Insert(newBooking Core) (Core, error)
	Update(newBooking Core) (Core, error)
	Delete(idBooking uint) error
}

type Services interface {
	GetAll() ([]Core, error)
	GetDetail(idBooking uint) (Core, error)
	GetRangerBooking(id_ranger uint) ([]Core, error)
	InsertData(newBooking Core) (Core, error)
	UpdateData(newBooking Core) (Core, error)
	DeleteData(idBooking uint) error
}
