package delivery

import (
	"GunTour/features/admin/domain"
)

type RegisterFormat struct {
	ProductName    string `json:"product_name" form:"product_name" validate:"required"`
	RentPrice      int    `json:"rent_price" form:"rent_price" validate:"required"`
	Detail         string `json:"detail" form:"detail" validate:"required"`
	Note           string `json:"note" form:"note" validate:"required"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type BookingProduct struct {
	ID         uint
	IdBooking  uint
	IdProduct  uint `json:"id_product" form:"id_product"`
	ProductQty int  `json:"product_qty" form:"product_qty"`
}
type UpdateFormat struct {
	ID             uint   `json:"id_product" form:"id_product"`
	ProductName    string `json:"product_name" form:"product_name"`
	RentPrice      int    `json:"rent_price" form:"rent_price"`
	Detail         string `json:"detail" form:"detail"`
	Note           string `json:"note" form:"note"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type UpdateMidtrans struct {
	ID            uint   `json:"id" form:"id"`
	OrderID       string `json:"order_id" form:"order_id"`
	StatusBooking string `json:"status" form:"status"`
}

type GetId struct {
	id uint `param:"id"`
}

type RangerFormat struct {
	ID          uint   `json:"id_ranger" form:"id_ranger"`
	UserID      uint   `json:"id_user" form:"id_user"`
	StatusApply string `json:"status_apply" form:"status_apply"`
	Role        string
}

func ToDomain(i interface{}) domain.ProductCore {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.ProductCore{ProductName: cnv.ProductName,
			RentPrice:      cnv.RentPrice,
			Detail:         cnv.Detail,
			Note:           cnv.Note,
			ProductPicture: cnv.ProductPicture}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.ProductCore{ID: cnv.ID, ProductName: cnv.ProductName,
			RentPrice:      cnv.RentPrice,
			Detail:         cnv.Detail,
			Note:           cnv.Note,
			ProductPicture: cnv.ProductPicture}
	}
	return domain.ProductCore{}
}

func ToDomainRanger(i interface{}) domain.RangerCore {
	switch i.(type) {
	case RangerFormat:
		cnv := i.(RangerFormat)
		return domain.RangerCore{ID: cnv.ID, UserID: cnv.UserID, StatusApply: cnv.StatusApply, User: domain.UserCore{Role: "ranger"}}
	}
	return domain.RangerCore{}
}
