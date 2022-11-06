package delivery

import (
	"GunTour/features/admin/domain"
	"time"
)

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseProduct(data interface{}) interface{} {
	return data
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

type GetPendakiResponse struct {
	IdUser    uint      `json:"id_pendaki" form:"id_pendaki"`
	FullName  string    `json:"fullname" form:"fullname"`
	Phone     string    `json:"phone" form:"phone"`
	Start     string    `json:"date_start" form:"date_start"`
	End       string    `json:"date_end" form:"date_end"`
	DateStart time.Time `json:"-" form:"-"`
	DateEnd   time.Time `json:"-" form:"-"`
}

type GetBookingResponse struct {
	ID        uint      `json:"id_booking" form:"id_booking"`
	IdUser    uint      `json:"id_pendaki" form:"id_pendaki"`
	FullName  string    `json:"fullname" form:"fullname"`
	Start     string    `json:"date_start" form:"date_start"`
	End       string    `json:"date_end" form:"date_end"`
	Entrance  string    `json:"entrance" form:"entrance"`
	Ticket    int       `json:"ticket" form:"ticket"`
	DateStart time.Time `json:"-" form:"-"`
	DateEnd   time.Time `json:"-" form:"-"`
}

type GetRangerResponse struct {
	ID            uint      `json:"id" form:"id"`
	IdUser        uint      `json:"id_user" form:"id_user"`
	DateStart     time.Time `json:"date_start" form:"date_start"`
	DateEnd       time.Time `json:"date_end" form:"date_end"`
	Entrance      string    `json:"entrance" form:"entrance"`
	Ticket        int       `json:"ticket" form:"ticket"`
	IdRanger      uint      `json:"id_ranger" form:"id_ranger"`
	GrossAmount   int       `json:"gross_amount" form:"gross_amount"`
	OrderId       string    `json:"order_id" form:"order_id"`
	Link          string    `json:"link" form:"link"`
	StatusBooking string    `json:"status" form:"status"`
}

type GetRangerApplyResponse struct {
	ID        uint      `json:"id_booking" form:"id_booking"`
	IdUser    uint      `json:"id_pendaki" form:"id_pendaki"`
	FullName  string    `json:"fullname" form:"fullname"`
	Phone     string    `json:"phone" form:"phone"`
	Start     string    `json:"date_start" form:"date_start"`
	End       string    `json:"date_end" form:"date_end"`
	Ticket    int       `json:"ticket" form:"ticket"`
	DateStart time.Time `json:"-" form:"-"`
	DateEnd   time.Time `json:"-" form:"-"`
}

type GetProductResponse struct {
	Message   string            `json:"message" form:"message"`
	Page      int               `json:"page" form:"page"`
	TotalPage int               `json:"total_page" form:"total_page"`
	Data      []ProductResponse `json:"data" form:"data"`
}
type ProductResponse struct {
	ID             uint   `json:"id_product" form:"id_product"`
	ProductName    string `json:"product_name" form:"product_name"`
	RentPrice      int    `json:"rent_price" form:"rent_price"`
	Detail         string `json:"detail" form:"detail"`
	Note           string `json:"note" form:"note"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "addproduct":
		cnv := core.(domain.ProductCore)
		res = ProductResponse{ID: cnv.ID, ProductName: cnv.ProductName, RentPrice: cnv.RentPrice, Detail: cnv.Detail, Note: cnv.Note, ProductPicture: cnv.ProductPicture}
	case "update":
		cnv := core.(domain.ProductCore)
		res = ProductResponse{ID: cnv.ID, ProductName: cnv.ProductName, RentPrice: cnv.RentPrice, Detail: cnv.Detail, Note: cnv.Note, ProductPicture: cnv.ProductPicture}
	}

	return res
}

func ToResponseArray(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "getpendaki":
		var arr []GetPendakiResponse
		val := core.([]domain.BookingCore)
		for _, cnv := range val {
			arr = append(arr, GetPendakiResponse{IdUser: cnv.IdUser, FullName: cnv.FullName, Phone: cnv.Phone,
				Start: cnv.DateStart.Format("2006-01-02"), End: cnv.DateEnd.Format("2006-01-02")})
		}
		res = arr

	case "getbooking":
		var arr []GetBookingResponse
		val := core.([]domain.BookingCore)
		for _, cnv := range val {
			arr = append(arr, GetBookingResponse{ID: uint(cnv.ID), IdUser: cnv.IdUser, FullName: cnv.FullName,
				Start: cnv.DateStart.Format("2006-01-02"), End: cnv.DateEnd.Format("2006-01-02"), Entrance: cnv.Entrance, Ticket: cnv.Ticket})
		}
		res = arr
	case "getproduct":
		var arr []ProductResponse
		val := core.([]domain.ProductCore)
		for _, cnv := range val {
			arr = append(arr, ProductResponse{ID: uint(cnv.ID), ProductName: cnv.ProductName, RentPrice: cnv.RentPrice,
				Detail: cnv.Detail, Note: cnv.Note, ProductPicture: cnv.ProductPicture})
		}
		res = arr
	}
	return res
}

func ToResponseProduct(core interface{}, message string, pages int, totalPage int, code string) interface{} {
	var arr []ProductResponse
	val := core.([]domain.ProductCore)
	for _, cnv := range val {
		arr = append(arr, ProductResponse{ID: uint(cnv.ID), ProductName: cnv.ProductName, RentPrice: cnv.RentPrice,
			Detail: cnv.Detail, Note: cnv.Note, ProductPicture: cnv.ProductPicture})
	}

	return GetProductResponse{Message: message, Page: pages, TotalPage: totalPage, Data: arr}
}
