package repository

import (
	"GunTour/features/google/domain"
	"time"

	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type Pesan struct {
	gorm.Model
	IdRanger uint
	Email    string
	Status   string
}

type Ranger struct {
	gorm.Model
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
}

type Date struct {
	gorm.Model
	DateStart time.Time
	DateEnd   time.Time
	Email     string
}

type Booking struct {
	gorm.Model
	IdUser          uint
	DateStart       time.Time
	DateEnd         time.Time
	Entrance        string
	Ticket          int
	IdRanger        uint
	GrossAmount     int
	Token           string
	OrderId         string
	Link            string
	StatusBooking   string
	StatusPendakian string
	FullName        string           `gorm:"-:migration;<-:false"`
	Phone           string           `gorm:"-:migration;<-:false"`
	Email           string           `gorm:"-:migration;<-:false"`
	BookingProducts []BookingProduct `gorm:"foreignKey:IdBooking"`
}

type BookingProduct struct {
	gorm.Model
	IdBooking   uint
	IdProduct   uint
	ProductQty  int
	ProductName string `gorm:"-:migration;<-:false"`
	RentPrice   int    `gorm:"-:migration;<-:false"`
}

func FromDomain(dc domain.Code) Code {
	return Code{
		Model:        gorm.Model{ID: dc.ID},
		Code:         dc.Code,
		AccessToken:  dc.AccessToken,
		TokenType:    dc.TokenType,
		RefreshToken: dc.RefreshToken,
		Expiry:       dc.Expiry,
	}
}

func ToDomain(dc Code) domain.Code {
	return domain.Code{
		ID:           dc.ID,
		Code:         dc.Code,
		AccessToken:  dc.AccessToken,
		TokenType:    dc.TokenType,
		RefreshToken: dc.RefreshToken,
		Expiry:       dc.Expiry,
	}
}

func ToDomainPesan(dc Pesan) domain.PesanCore {
	return domain.PesanCore{
		ID:       dc.ID,
		IdRanger: dc.IdRanger,
		Email:    dc.Email,
		Status:   dc.Status,
	}
}

func ToDomainRanger(dc Ranger) domain.RangerCore {
	return domain.RangerCore{
		ID:          dc.ID,
		Status:      dc.Status,
		StatusApply: dc.StatusApply,
	}
}

func ToDomainBooking(db Booking) domain.BookingCore {
	return domain.BookingCore{
		ID:                  db.ID,
		IdUser:              db.IdUser,
		DateStart:           db.DateStart,
		DateEnd:             db.DateEnd,
		Entrance:            db.Entrance,
		Ticket:              db.Ticket,
		IdRanger:            db.IdRanger,
		GrossAmount:         db.GrossAmount,
		Token:               db.Token,
		OrderId:             db.OrderId,
		Link:                db.Link,
		StatusBooking:       db.StatusBooking,
		StatusPendakian:     db.StatusPendakian,
		FullName:            db.FullName,
		Phone:               db.Phone,
		Email:               db.Email,
		BookingProductCores: []domain.BookingProductCore{},
	}
}

func ToDomainCore(db Booking, dp []BookingProduct) domain.BookingCore {
	var res []domain.BookingProductCore
	for _, val := range dp {
		res = append(res, domain.BookingProductCore{
			ID:          val.ID,
			IdBooking:   val.IdBooking,
			IdProduct:   val.IdProduct,
			ProductQty:  val.ProductQty,
			ProductName: val.ProductName,
			RentPrice:   val.RentPrice,
		})
	}
	return domain.BookingCore{
		ID:                  db.ID,
		IdUser:              db.IdUser,
		DateStart:           db.DateStart,
		DateEnd:             db.DateEnd,
		Entrance:            db.Entrance,
		Ticket:              db.Ticket,
		IdRanger:            db.IdRanger,
		GrossAmount:         db.GrossAmount,
		Token:               db.Token,
		OrderId:             db.OrderId,
		Link:                db.Link,
		StatusBooking:       db.StatusBooking,
		StatusPendakian:     db.StatusPendakian,
		FullName:            db.FullName,
		Email:               db.Email,
		BookingProductCores: res,
	}
}
