package delivery

import (
	"GunTour/features/booking/domain"
	"time"
)

type RegisterFormat struct {
	IdUser      uint      `json:"id_user" form:"id_user"`
	DateStart   time.Time `json:"date_start" form:"date_start"`
	DateEnd     time.Time `json:"date_end" form:"date_end"`
	Entrance    string    `json:"entrance" form:"entrance"`
	Ticket      int       `json:"ticket" form:"ticket"`
	IdRanger    uint      `json:"id_ranger" form:"id_ranger"`
	GrossAmount int       `json:"gross_amount" form:"gross_amount"`
}

type UpdateFormat struct {
	ID              uint      `json:"id" form:"id"`
	IdUser          uint      `json:"id_user" form:"id_user"`
	DateStart       time.Time `json:"date_start" form:"date_start"`
	DateEnd         time.Time `json:"date_end" form:"date_end"`
	Entrance        string    `json:"entrance" form:"entrance"`
	Ticket          int       `json:"ticket" form:"ticket"`
	IdRanger        uint      `json:"id_ranger" form:"id_ranger"`
	GrossAmount     int       `json:"gross_amount" form:"gross_amount"`
	StatusBooking   string    `json:"status_booking" form:"status_booking"`
	StatusPendakian string    `json:"status_pendakian" form:"status_pendakian"`
}

type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, IdUser: cnv.IdRanger, DateStart: cnv.DateStart, DateEnd: cnv.DateEnd, Entrance: cnv.Entrance, Ticket: cnv.Ticket,
			IdRanger: cnv.IdRanger, GrossAmount: cnv.GrossAmount, StatusBooking: cnv.StatusBooking, StatusPendakian: cnv.StatusPendakian}
	}
	return domain.Core{}
}
