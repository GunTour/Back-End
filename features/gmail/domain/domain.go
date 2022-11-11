package domain

import "time"

type Code struct {
	ID           uint
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type PesanCore struct {
	ID       uint
	IdRanger uint
	Email    string
	Status   string
}

type RangerCore struct {
	ID          uint
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
}

type DateCore struct {
	ID        uint
	DateStart time.Time
	DateEnd   time.Time
	Email     string
}

type BookingCore struct {
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

type Repository interface {
	InsertCode(code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() (PesanCore, RangerCore)
	GetPesanCal() BookingCore
}

type Services interface {
	AddCode(Code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() (PesanCore, RangerCore)
	GetPesanCal() BookingCore
}
