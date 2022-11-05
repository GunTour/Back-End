package services

import (
	"GunTour/features/admin/domain"
	"errors"
)

type adminService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &adminService{
		qry: repo,
	}
}

func (as *adminService) GetPendaki() ([]domain.BookingCore, error) {
	res, err := as.qry.GetPendaki()
	if err != nil {
		return []domain.BookingCore{}, errors.New("no data")
	}

	return res, nil

}

// GetRanger(id uint) ([]UserCore, []UserCore, error)
func (as *adminService) GetBooking() ([]domain.BookingCore, error) {
	res, err := as.qry.GetBooking()
	if err != nil {
		return []domain.BookingCore{}, errors.New("no data")
	}

	return res, nil
}
