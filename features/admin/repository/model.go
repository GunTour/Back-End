package repository

import (
	"GunTour/features/admin/domain"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string
	Email       string
	Password    string
	Role        string
	Phone       string
	Address     string
	Dob         string
	Gender      string
	UserPicture string
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
	FullName        string `gorm:"-:migration" gorm:"<-"`
	Phone           string `gorm:"-:migration" gorm:"<-"`
}

type Product struct {
	gorm.Model
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
}

func FromDomainProduct(db domain.ProductCore) Product {
	return Product{
		Model:          gorm.Model{ID: db.ID},
		ProductName:    db.ProductName,
		RentPrice:      db.RentPrice,
		Detail:         db.Detail,
		Note:           db.Note,
		ProductPicture: db.ProductPicture,
	}
}

func FromDomainBooking(db domain.BookingCore) Booking {
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

func ToDomainProduct(db Product) domain.ProductCore {
	return domain.ProductCore{
		ID:             db.ID,
		ProductName:    db.ProductName,
		RentPrice:      db.RentPrice,
		Detail:         db.Detail,
		Note:           db.Note,
		ProductPicture: db.ProductPicture,
	}
}

func ToDomainBooking(db []Booking) []domain.BookingCore {
	var arr []domain.BookingCore
	for _, val := range db {
		arr = append(arr, domain.BookingCore{
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
			FullName:        val.FullName,
			Phone:           val.Phone,
		})
	}
	return arr
}

func ToDomainProductArr(db []Product) []domain.ProductCore {
	var arr []domain.ProductCore
	for _, val := range db {
		arr = append(arr, domain.ProductCore{
			ID:             val.ID,
			ProductName:    val.ProductName,
			RentPrice:      val.RentPrice,
			Detail:         val.Detail,
			Note:           val.Note,
			ProductPicture: val.ProductPicture,
		})
	}
	return arr
}
