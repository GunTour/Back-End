package repository

import (
	"GunTour/features/booking/domain"
	"time"

	"gorm.io/gorm"
)

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
	BookingProducts []BookingProduct `gorm:"foreignKey:IdBooking"`
}

type BookingProduct struct {
	gorm.Model
	IdBooking uint
	IdProduct uint
}

func FromDomain(db domain.Core) Booking {
	return Booking{
		Model:           gorm.Model{ID: db.ID},
		IdUser:          db.IdUser,
		DateStart:       db.DateStart,
		DateEnd:         db.DateEnd,
		Entrance:        db.Entrance,
		Ticket:          db.Ticket,
		IdRanger:        db.IdRanger,
		GrossAmount:     db.GrossAmount,
		Token:           db.Token,
		OrderId:         db.OrderId,
		Link:            db.Link,
		StatusBooking:   db.StatusBooking,
		StatusPendakian: db.StatusPendakian,
	}
}

func FromDomainProduct(dp []domain.BookingProductCore, id uint) []BookingProduct {
	var res []BookingProduct
	for _, val := range dp {
		res = append(res, BookingProduct{Model: gorm.Model{ID: val.ID}, IdBooking: id,
			IdProduct: val.IdProduct})
	}
	return res
}

func ToDomain(db Booking) domain.Core {
	return domain.Core{
		ID:              db.ID,
		IdUser:          db.IdUser,
		DateStart:       db.DateStart,
		DateEnd:         db.DateEnd,
		Entrance:        db.Entrance,
		Ticket:          db.Ticket,
		IdRanger:        db.IdRanger,
		GrossAmount:     db.GrossAmount,
		Token:           db.Token,
		OrderId:         db.OrderId,
		Link:            db.Link,
		StatusBooking:   db.StatusBooking,
		StatusPendakian: db.StatusPendakian,
	}
}

func ToDomainArray(dp []Booking) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{
			ID:              val.ID,
			IdUser:          val.IdUser,
			DateStart:       val.DateStart,
			DateEnd:         val.DateEnd,
			Entrance:        val.Entrance,
			Ticket:          val.Ticket,
			IdRanger:        val.IdRanger,
			GrossAmount:     val.GrossAmount,
			Token:           val.Token,
			OrderId:         val.OrderId,
			Link:            val.Link,
			StatusBooking:   val.StatusBooking,
			StatusPendakian: val.StatusPendakian,
		})
	}
	return res
}