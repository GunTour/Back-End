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
	FullName        string `gorm:"-:migration;<-:false"`
	Phone           string `gorm:"-:migration;<-:false"`
}

type Pesan struct {
	gorm.Model
	IdRanger uint
	Email    string
	Status   string
}

type Product struct {
	gorm.Model
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
}

type Ranger struct {
	gorm.Model
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
	User        User
}

type Climber struct {
	gorm.Model
	IsClimber     int
	MaleClimber   int
	FemaleClimber int
}

type Code struct {
	gorm.Model
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

func FromDomainClimber(db domain.ClimberCore) Climber {
	return Climber{
		Model:         gorm.Model{ID: db.ID},
		IsClimber:     db.IsClimber,
		MaleClimber:   db.MaleClimber,
		FemaleClimber: db.FemaleClimber,
	}
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

func ToDomainCode(dc Code) domain.Code {
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

func ToDomainClimber(db Climber) domain.ClimberCore {
	return domain.ClimberCore{
		ID:            db.ID,
		IsClimber:     db.IsClimber,
		MaleClimber:   db.MaleClimber,
		FemaleClimber: db.FemaleClimber,
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

func FromDomainPesan(mail string, dr Ranger) Pesan {
	return Pesan{
		Model:    gorm.Model{ID: dr.ID},
		IdRanger: dr.ID,
		Email:    mail,
		Status:   dr.StatusApply,
	}
}

func FromDomainRanger(dr domain.RangerCore) Ranger {
	return Ranger{
		Model:       gorm.Model{ID: dr.ID},
		UserID:      dr.UserID,
		Docs:        dr.Docs,
		Price:       dr.Price,
		Detail:      dr.Detail,
		Status:      dr.Status,
		StatusApply: dr.StatusApply,
		User:        User{Phone: dr.User.Phone},
	}
}

func ToDomainRanger(r Ranger) domain.RangerCore {
	return domain.RangerCore{
		ID:          r.ID,
		UserID:      r.UserID,
		Docs:        r.Docs,
		Price:       r.Price,
		Detail:      r.Detail,
		Status:      r.Status,
		StatusApply: r.StatusApply,
		User:        domain.UserCore{Phone: r.User.Phone},
	}
}

func ToDomainRangerArray(ar []Ranger) []domain.RangerCore {
	var arr []domain.RangerCore
	for _, val := range ar {
		arr = append(arr, domain.RangerCore{
			ID:          val.ID,
			UserID:      val.UserID,
			Docs:        val.Docs,
			Price:       val.Price,
			Detail:      val.Detail,
			Status:      val.Status,
			StatusApply: val.StatusApply,
			User:        domain.UserCore{FullName: val.User.FullName, Address: val.User.Address, Phone: val.User.Phone, Gender: val.User.Gender},
		})
	}
	return arr
}

func FromDomainUser(du domain.UserCore) User {
	return User{
		Model: gorm.Model{ID: (du.ID)},
		Phone: du.Phone,
	}
}

func ToDomainUser(u User) domain.UserCore {
	return domain.UserCore{
		ID:    (u.ID),
		Phone: u.Phone,
	}
}
