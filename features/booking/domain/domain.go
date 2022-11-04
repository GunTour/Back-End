package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                  uint
	IdUser              uint
	DateStart           time.Time
	DateEnd             time.Time
	Entrance            string
	Ticket              int
	IdRanger            uint
	GrossAmount         int
	Token               string
	Link                string
	StatusBooking       string
	StatusPendakian     string
	BookingProductCores []BookingProductCore
}

type BookingProductCore struct {
	ID        uint
	IdBooking uint
	IdProduct uint
}

type Repository interface {
	Get(idUser uint) ([]Core, error)
	GetID(idBooking uint) (Core, error)
	GetRanger(idRanger uint) ([]Core, error)
	Insert(newBooking Core) (Core, error)
	Update(newBooking Core) (Core, error)
	Delete(idBooking uint) error
}

type Services interface {
	GetAll(idUser uint) ([]Core, error)
	GetDetail(idBooking uint) (Core, error)
	GetRangerBooking(idRanger uint) ([]Core, error)
	InsertData(newBooking Core) (Core, error)
	UpdateData(newBooking Core) (Core, error)
	DeleteData(idBooking uint) error
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	GetRangerBooking() echo.HandlerFunc
	InsertData() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
	DeleteData() echo.HandlerFunc
}
