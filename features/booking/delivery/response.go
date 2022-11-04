package delivery

import (
	"GunTour/features/booking/domain"
	"time"
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
	IdUser      uint      `json:"id_user" form:"id_user"`
	DateStart   time.Time `json:"date_start" form:"date_start"`
	DateEnd     time.Time `json:"date_end" form:"date_end"`
	Entrance    string    `json:"entrance" form:"entrance"`
	Ticket      int       `json:"ticket" form:"ticket"`
	IdRanger    uint      `json:"id_ranger" form:"id_ranger"`
	GrossAmount int       `json:"gross_amount" form:"gross_amount"`
	Link        string    `json:"link" form:"link"`
}

type UpdateResponse struct {
	ID              uint      `json:"id" form:"id"`
	IdUser          uint      `json:"id_user" form:"id_user"`
	DateStart       time.Time `json:"date_start" form:"date_start"`
	DateEnd         time.Time `json:"date_end" form:"date_end"`
	Entrance        string    `json:"entrance" form:"entrance"`
	Ticket          int       `json:"ticket" form:"ticket"`
	IdRanger        uint      `json:"id_ranger" form:"id_ranger"`
	GrossAmount     int       `json:"gross_amount" form:"gross_amount"`
	Link            string    `json:"link" form:"link"`
	StatusBooking   string    `json:"status_booking" form:"status_booking"`
	StatusPendakian string    `json:"status_pendakian" form:"status_pendakian"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, Link: cnv.Link}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, Link: cnv.Link}
	}

	return res
}

func ToResponseArray(core interface{}, code string) interface{} {
	var res interface{}
	var arr []UpdateResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, UpdateResponse{ID: cnv.ID, IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, Link: cnv.Link})
	}
	res = arr
	return res
}
