package delivery

import (
	"GunTour/features/booking/domain"
	"time"
)

type RegisterFormat struct {
	IdUser        uint             `json:"id_user" form:"id_user"`
	DateStart     time.Time        `json:"date_start" form:"date_start"`
	DateEnd       time.Time        `json:"date_end" form:"date_end"`
	Entrance      string           `json:"entrance" form:"entrance"`
	Ticket        int              `json:"ticket" form:"ticket"`
	OrderId       string           `json:"order_id" form:"order_id"`
	Product       []BookingProduct `json:"product" form:"product"`
	IdRanger      uint             `json:"id_ranger" form:"id_ranger"`
	StatusBooking string           `json:"status_booking" form:"status_booking"`
	GrossAmount   int              `json:"gross_amount" form:"gross_amount"`
}

type BookingProduct struct {
	ID         uint
	IdBooking  uint
	IdProduct  uint `json:"id_product" form:"id_product"`
	ProductQty int  `json:"product_qty" form:"product_qty"`
}
type UpdateFormat struct {
	ID              uint             `json:"id" form:"id"`
	IdUser          uint             `json:"id_user" form:"id_user"`
	DateStart       time.Time        `json:"date_start" form:"date_start"`
	DateEnd         time.Time        `json:"date_end" form:"date_end"`
	Entrance        string           `json:"entrance" form:"entrance"`
	Ticket          int              `json:"ticket" form:"ticket"`
	Product         []BookingProduct `json:"product" form:"product"`
	IdRanger        uint             `json:"id_ranger" form:"id_ranger"`
	GrossAmount     int              `json:"gross_amount" form:"gross_amount"`
	StatusBooking   string           `json:"status" form:"status"`
	StatusPendakian string           `json:"status_pendakian" form:"status_pendakian"`
}

type UpdateMidtrans struct {
	ID            uint   `json:"id" form:"id"`
	OrderID       string `json:"order_id" form:"order_id"`
	StatusBooking string `json:"status" form:"status"`
}

type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		var arr []domain.BookingProductCore
		cnv := i.(RegisterFormat)
		for _, val := range cnv.Product {
			arr = append(arr, domain.BookingProductCore{IdProduct: val.IdProduct, ProductQty: val.ProductQty})
		}
		return domain.Core{IdUser: cnv.IdUser, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			OrderId: cnv.OrderId, BookingProductCores: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, StatusBooking: cnv.StatusBooking}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		var arr []domain.BookingProductCore
		cnv := i.(UpdateFormat)
		for _, val := range cnv.Product {
			arr = append(arr, domain.BookingProductCore{IdProduct: val.IdProduct, ProductQty: val.ProductQty})
		}
		return domain.Core{ID: cnv.ID, IdUser: cnv.IdUser, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			BookingProductCores: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, StatusBooking: cnv.StatusBooking}
	case UpdateMidtrans:
		cnv := i.(UpdateMidtrans)
		return domain.Core{OrderId: cnv.OrderID, StatusBooking: cnv.StatusBooking}
	}
	return domain.Core{}
}
