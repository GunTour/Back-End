package delivery

import (
	"GunTour/features/booking/domain"
)

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type RegisterResponse struct {
	IdUser      uint             `json:"id_user" form:"id_user"`
	DateStart   string           `json:"date_start" form:"date_start"`
	DateEnd     string           `json:"date_end" form:"date_end"`
	Entrance    string           `json:"entrance" form:"entrance"`
	Ticket      int              `json:"ticket" form:"ticket"`
	Product     []BookingProduct `json:"product" form:"product"`
	IdRanger    uint             `json:"id_ranger" form:"id_ranger"`
	GrossAmount int              `json:"gross_amount" form:"gross_amount"`
	OrderId     string           `json:"order_id" form:"order_id"`
	Link        string           `json:"link" form:"link"`
}

type UpdateResponse struct {
	ID            uint             `json:"id" form:"id"`
	IdUser        uint             `json:"id_user" form:"id_user"`
	DateStart     string           `json:"date_start" form:"date_start"`
	DateEnd       string           `json:"date_end" form:"date_end"`
	Entrance      string           `json:"entrance" form:"entrance"`
	Ticket        int              `json:"ticket" form:"ticket"`
	Product       []BookingProduct `json:"product" form:"product"`
	IdRanger      uint             `json:"id_ranger" form:"id_ranger"`
	GrossAmount   int              `json:"gross_amount" form:"gross_amount"`
	OrderId       string           `json:"order_id" form:"order_id"`
	Link          string           `json:"link" form:"link"`
	StatusBooking string           `json:"status" form:"status"`
}

type DetailResponse struct {
	ID            uint             `json:"id_booking" form:"id_booking"`
	Start         string           `json:"date_start" form:"date_start"`
	End           string           `json:"date_end" form:"date_end"`
	Entrance      string           `json:"entrance" form:"entrance"`
	Ticket        int              `json:"ticket" form:"ticket"`
	Product       []BookingProduct `json:"product" form:"product"`
	IdRanger      uint             `json:"id_ranger" form:"id_ranger"`
	GrossAmount   int              `json:"gross_amount" form:"gross_amount"`
	OrderId       string           `json:"order_id" form:"order_id"`
	Link          string           `json:"link" form:"link"`
	StatusBooking string           `json:"status" form:"status"`
	DateStart     string           `json:"-" form:"-"`
	DateEnd       string           `json:"-" form:"-"`
}

type GetResponse struct {
	ID            uint   `json:"id_booking" form:"id_booking"`
	GrossAmount   int    `json:"gross_amount" form:"gross_amount"`
	OrderId       string `json:"order_id" form:"order_id"`
	Link          string `json:"link" form:"link"`
	StatusBooking string `json:"status" form:"status"`
}

type GetRangerResponse struct {
	ID        uint   `json:"id_booking" form:"id_booking"`
	IdUser    uint   `json:"id_pendaki" form:"id_pendaki"`
	FullName  string `json:"fullname" form:"fullname"`
	Phone     string `json:"phone" form:"phone"`
	Start     string `json:"date_start" form:"date_start"`
	End       string `json:"date_end" form:"date_end"`
	Ticket    int    `json:"ticket" form:"ticket"`
	DateStart string `json:"-" form:"-"`
	DateEnd   string `json:"-" form:"-"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		var arr []BookingProduct
		for _, val := range cnv.BookingProductCores {
			arr = append(arr, BookingProduct{ID: val.ID, IdBooking: val.IdBooking, IdProduct: val.IdProduct, ProductQty: val.ProductQty})
		}
		res = RegisterResponse{IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			Product: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link}
	case "update":
		cnv := core.(domain.Core)
		var arr []BookingProduct
		for _, val := range cnv.BookingProductCores {
			arr = append(arr, BookingProduct{ID: val.ID, IdBooking: val.IdBooking, IdProduct: val.IdProduct, ProductQty: val.ProductQty})
		}
		res = UpdateResponse{ID: cnv.ID, IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			Product: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link}
	case "getdetails":
		cnv := core.(domain.Core)
		var arr []BookingProduct
		for _, val := range cnv.BookingProductCores {
			arr = append(arr, BookingProduct{ID: val.ID, IdBooking: val.IdBooking, IdProduct: val.IdProduct, ProductQty: val.ProductQty,
				ProductName: val.ProductName, RentPrice: val.RentPrice})
		}
		res = DetailResponse{ID: cnv.ID, Start: cnv.DateStart, End: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			Product: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link, StatusBooking: cnv.StatusBooking}
	}

	return res
}

func ToResponseArray(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "getall":
		var arr []GetResponse
		val := core.([]domain.Core)
		for _, cnv := range val {
			arr = append(arr, GetResponse{ID: cnv.ID, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link, StatusBooking: cnv.StatusBooking})
		}
		res = arr

	case "getranger":
		var arr []GetRangerResponse
		val := core.([]domain.Core)
		for _, cnv := range val {
			arr = append(arr, GetRangerResponse{ID: cnv.ID, IdUser: cnv.IdUser, FullName: cnv.FullName, Phone: cnv.Phone, Start: cnv.DateStart, End: cnv.DateEnd, Ticket: cnv.Ticket})
		}
		res = arr
	}
	return res
}

func ToDomainCheckMidtrans(i *coreapi.TransactionStatusResponse) domain.Core {
	return domain.Core{OrderId: i.OrderID, StatusBooking: i.TransactionStatus}
}
