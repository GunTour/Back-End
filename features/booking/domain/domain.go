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
	OrderId             string
	Link                string
	StatusBooking       string
	StatusPendakian     string
	FullName            string
	Phone               string
	Email               string
	BookingProductCores []BookingProductCore
}

type BookingProductCore struct {
	ID          uint
	IdBooking   uint
	IdProduct   uint
	ProductQty  int
	ProductName string
	RentPrice   int
}

type Pendaki struct {
	Email   string
	Address string
}

type Ranger struct {
	UserID   int
	FullName string
	Email    string
}

type Code struct {
	ID           uint
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type Repository interface {
	Get(idUser uint) ([]Core, error)
	GetID(idBooking uint) (Core, error)
	GetRanger(idRanger uint) ([]Core, error)
	Insert(newBooking Core) (Core, error)
	Update(newBooking Core) (Core, error)
	Delete(idBooking uint) error
	UpdateMidtrans(newBooking Core) error
	GetCode() (Code, error)
}

type Services interface {
	GetAll(idUser uint) ([]Core, error)
	GetDetail(idBooking uint) (Core, error)
	GetRangerBooking(idRanger uint) ([]Core, error)
	InsertData(newBooking Core) (Core, error)
	UpdateData(newBooking Core) (Core, error)
	DeleteData(idBooking uint) error
	UpdateMidtrans(newBooking Core) error
	GetCode() (Code, error)
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	GetRangerBooking() echo.HandlerFunc
	InsertData() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
	DeleteData() echo.HandlerFunc
	UpdateMidtrans() echo.HandlerFunc
}
