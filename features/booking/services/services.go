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

func (bs *bookingService) GetAll(idUser uint) ([]domain.Core, error) {
	res, err := bs.qry.Get(idUser)
	if err != nil {
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}
func (bs *bookingService) GetDetail(idBooking uint) (domain.Core, error) {
	res, err := bs.qry.GetID(idBooking)
	if err != nil {
		return domain.Core{}, errors.New("no data")
	}

	return res, nil
}

func (bs *bookingService) GetRangerBooking(idRanger uint) ([]domain.Core, error) {
	res, err := bs.qry.GetRanger(idRanger)
	if err != nil {
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}

func (bs *bookingService) InsertData(newBooking domain.Core) (domain.Core, error) {
	midtrans := helper.OrderMidtrans(newBooking.OrderId, int64(newBooking.GrossAmount))
	newBooking.Link = midtrans.RedirectURL
	res, err := bs.qry.Insert(newBooking)
	if err != nil {
		return domain.Core{}, err
	}

	return res, nil
}

func (bs *bookingService) UpdateData(newBooking domain.Core) (domain.Core, error) {
	res, err := bs.qry.Update(newBooking)
	if err != nil {
		return domain.Core{}, err
	}
	return res, nil
}

func (bs *bookingService) DeleteData(idBooking uint) error {
	err := bs.qry.Delete(idBooking)
	if err != nil {
		return err
	}
	return nil
}

func (bs *bookingService) UpdateMidtrans(newBooking domain.Core) error {
	err := bs.qry.UpdateMidtrans(newBooking)
	if err != nil {
		return err
	}
	return nil
}
