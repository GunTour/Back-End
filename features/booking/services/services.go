package services

import (
	"GunTour/features/booking/domain"
	"GunTour/utils/helper"
	"errors"
)

type bookingService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &bookingService{
		qry: repo,
	}
}

// SERVICE TO GET ALL USER'S BOOKING DATA
func (bs *bookingService) GetAll(idUser uint) ([]domain.Core, error) {
	res, err := bs.qry.Get(idUser)
	if err != nil {
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}

// SERVICE TO GET BOOKING DETAIL'S DATA
func (bs *bookingService) GetDetail(idBooking uint) (domain.Core, error) {
	res, err := bs.qry.GetID(idBooking)
	if err != nil {
		return domain.Core{}, errors.New("no data")
	}

	return res, nil
}

// SERVICE TO GET ALL RANGER'S BOOKING DATA
func (bs *bookingService) GetRangerBooking(idRanger uint) ([]domain.Core, error) {
	res, err := bs.qry.GetRanger(idRanger)
	if err != nil {
		return []domain.Core{}, err
	}

	return res, nil
}

// SERVICE TO ADD BOOKING
func (bs *bookingService) InsertData(newBooking domain.Core) (domain.Core, error) {
	midtrans := helper.OrderMidtrans(newBooking.OrderId, int64(newBooking.GrossAmount))
	newBooking.Link = midtrans.RedirectURL
	res, err := bs.qry.Insert(newBooking)
	if err != nil {
		return domain.Core{}, err
	}

	return res, nil
}

// SERVICE TO UPDATE BOOKING DATA
func (bs *bookingService) UpdateData(newBooking domain.Core) (domain.Core, error) {
	res, err := bs.qry.Update(newBooking)
	if err != nil {
		return domain.Core{}, err
	}
	return res, nil
}

// SERVICE TO DELETE BOOKING DATA
func (bs *bookingService) DeleteData(idBooking uint) error {
	err := bs.qry.Delete(idBooking)
	if err != nil {
		return err
	}
	return nil
}

// SERVICE TO UPDATE BOOKING DATA AFTER PAYMENT MIDTRANS
func (bs *bookingService) UpdateMidtrans(newBooking domain.Core) error {
	inputMidtrans := helper.CheckMidtrans(newBooking.OrderId)
	newBooking.StatusBooking = inputMidtrans.TransactionStatus
	err := bs.qry.UpdateMidtrans(newBooking)
	if err != nil {
		return err
	}
	return nil
}

// SERVICE TO GET OAUTH TOKEN
func (bs *bookingService) GetCode() (domain.Code, error) {
	res, err := bs.qry.GetCode()
	if err != nil {
		return res, err
	}

	return res, nil
}
