package repository

import (
	"GunTour/features/booking/domain"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	IdUser          uint
	DateStart       string
	DateEnd         string
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

type Pendaki struct {
	Email   string
	Address string
}

type Ranger struct {
	UserID   int
	FullName string
	Email    string
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
			IdProduct: val.IdProduct, ProductQty: val.ProductQty})
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

func ToDomainCore(db Booking, dp []BookingProduct) domain.Core {
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
	return domain.Core{
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
		BookingProductCores: res,
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

func ToDomainArrayRanger(dp []Booking) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{
			ID:        val.ID,
			IdUser:    val.IdUser,
			FullName:  val.FullName,
			Phone:     val.Phone,
			DateStart: val.DateStart,
			DateEnd:   val.DateEnd,
			Ticket:    val.Ticket,
		})
	}
	return res
}

func (m *Ranger) ToModelRanger() domain.Ranger {
	return domain.Ranger{
		UserID:   m.UserID,
		FullName: m.FullName,
		Email:    m.Email,
	}
}

func (m *Pendaki) ToModelPendaki() domain.Pendaki {
	return domain.Pendaki{
		Email:   m.Email,
		Address: m.Address,
	}
}
