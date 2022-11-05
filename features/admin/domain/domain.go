package domain

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type UserCore struct {
	ID          int
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

type BookingCore struct {
	ID              uint
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
	FullName        string
	Phone           string
}

type ProductCore struct {
	ID             uint
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
}

type Repository interface {
	GetPendaki() ([]BookingCore, error)
	// GetRanger(id uint) ([]UserCore, []UserCore, error)
	GetBooking() ([]BookingCore, error)
	GetProduct(page int) ([]ProductCore, int, int, error)
	InsertProduct(newProduct ProductCore) (ProductCore, error)
	UpdateProduct(newProduct ProductCore) (ProductCore, error)
	DeleteProduct(id int) error
}

type Services interface {
	GetPendaki() ([]BookingCore, error)
	// GetRanger(id uint) ([]UserCore, []UserCore, error)
	GetBooking() ([]BookingCore, error)
	GetProduct(page int) ([]ProductCore, int, int, error)
	AddProduct(newProduct ProductCore, file multipart.File, fileheader *multipart.FileHeader) (ProductCore, error)
	EditProduct(newProduct ProductCore, file multipart.File, fileheader *multipart.FileHeader) (ProductCore, error)
	RemoveProduct(id int) error
}

type Handler interface {
	GetPendaki() echo.HandlerFunc
	// GetRanger() echo.HandlerFunc
	GetBooking() echo.HandlerFunc
	GetProduct() echo.HandlerFunc
	AddProduct() echo.HandlerFunc
	EditProduct() echo.HandlerFunc
	RemoveProduct() echo.HandlerFunc
}
