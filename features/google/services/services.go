package services

import "GunTour/features/google/domain"

type googleService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &googleService{
		qry: repo,
	}
}

func (bs *googleService) InsertCode(Code domain.Code) error {
	err := bs.qry.InsertCode(Code)
	if err != nil {
		return err
	}

	return nil
}

func (bs *googleService) GetPesan() (domain.PesanCore, domain.RangerCore) {
	res, resRanger := bs.qry.GetPesan()
	return res, resRanger
}

func (bs *googleService) GetPesanCal() domain.BookingCore {
	res := bs.qry.GetPesanCal()
	return res
}
