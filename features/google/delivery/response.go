package delivery

import (
	"GunTour/features/google/domain"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

func SuccessResponse(msg string, url string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    fmt.Sprintf("%v", url),
	}
}

func SuccessResponseRanger(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseBooking(msg string, data interface{}) map[string]interface{} {
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

type URL struct {
	Url string `json:"url" form:"url"`
}

type RangerAccepted struct {
	ID          uint   `json:"id_ranger" form:"id_ranger"`
	Status      string `json:"status" form:"status"`
	StatusApply string `json:"status_apply" form:"status_apply"`
}

type RangerAcceptedNoMail struct {
	ID          uint   `json:"id_ranger" form:"id_ranger"`
	Status      string `json:"status" form:"status"`
	StatusApply string `json:"status_apply" form:"status_apply"`
	Url         string `json:"email_sender" form:"email_sender"`
}

type BookingProduct struct {
	ID          uint
	IdBooking   uint
	IdProduct   uint   `json:"id_product" form:"id_product"`
	ProductQty  int    `json:"product_qty" form:"product_qty"`
	ProductName string `json:"product_name" form:"product_name"`
	RentPrice   int    `json:"rent_price" form:"rent_price"`
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
	URL           string           `json:"url" form:"url"`
	DateStart     time.Time        `json:"-" form:"-"`
	DateEnd       time.Time        `json:"-" form:"-"`
}
type BookingMake struct {
	ID        uint      `json:"id_booking" form:"id_booking"`
	DateStart time.Time `json:"date_start" form:"date_start"`
	DateEnd   time.Time `json:"date_end" form:"date_end"`
	Entrance  string    `json:"entrance" form:"entrance"`
	Ticket    int       `json:"ticket" form:"ticket"`
	Url       string    `json:"url" form:"url"`
}

func FromDomain(core domain.Code) *oauth2.Token {

	return &oauth2.Token{AccessToken: core.AccessToken, TokenType: core.TokenType, RefreshToken: core.RefreshToken, Expiry: core.Expiry}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "ranger":
		cnv := core.(domain.RangerCore)
		res = RangerAccepted{ID: cnv.ID, Status: cnv.Status, StatusApply: cnv.StatusApply}
	case "book":
		cnv := core.(domain.BookingCore)
		var arr []BookingProduct
		for _, val := range cnv.BookingProductCores {
			arr = append(arr, BookingProduct{ID: val.ID, IdBooking: val.IdBooking, IdProduct: val.IdProduct, ProductQty: val.ProductQty,
				ProductName: val.ProductName, RentPrice: val.RentPrice})
		}
		res = DetailResponse{ID: cnv.ID, Start: cnv.DateStart.Format("2006-01-02"), End: cnv.DateEnd.Format("2006-01-02"), Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			Product: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link, StatusBooking: cnv.StatusBooking}
	}

	return res
}

func ToResponseGagal(core interface{}, authURL string, code string) interface{} {
	var res interface{}
	switch code {
	case "ranger":
		cnv := core.(domain.RangerCore)
		res = RangerAcceptedNoMail{ID: cnv.ID, Status: cnv.Status, StatusApply: cnv.StatusApply, Url: authURL}
	case "book":
		cnv := core.(domain.BookingCore)
		var arr []BookingProduct
		for _, val := range cnv.BookingProductCores {
			arr = append(arr, BookingProduct{ID: val.ID, IdBooking: val.IdBooking, IdProduct: val.IdProduct, ProductQty: val.ProductQty,
				ProductName: val.ProductName, RentPrice: val.RentPrice})
		}
		res = DetailResponse{ID: cnv.ID, Start: cnv.DateStart.Format("2006-01-02"), End: cnv.DateEnd.Format("2006-01-02"), Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			Product: arr, IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, OrderId: cnv.OrderId, Link: cnv.Link, StatusBooking: cnv.StatusBooking, URL: authURL}
	}

	return res
}
