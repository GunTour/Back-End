package delivery

import (
	"GunTour/features/gmail/domain"
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
		res = BookingMake{ID: cnv.ID, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket}
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
		res = BookingMake{ID: cnv.ID, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket, Url: authURL}
	}

	return res
}
