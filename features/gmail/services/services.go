package services

import "GunTour/features/gmail/domain"

type gmailService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &gmailService{
		qry: repo,
	}
}

func (bs *gmailService) AddCode(Code string) error {
	err := bs.qry.InsertCode(Code)
	if err != nil {
		return err
	}

	return nil
}

func (bs *gmailService) UpdateCode(Code domain.Code) error {
	err := bs.qry.UpdateCode(Code)
	if err != nil {
		return err
	}

	return nil
}

func (bs *gmailService) GetCode() (domain.Code, error) {
	res, err := bs.qry.GetCode()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (bs *gmailService) GetPesan() (domain.PesanCore, domain.RangerCore) {
	res, resRanger := bs.qry.GetPesan()
	return res, resRanger
}

func (bs *gmailService) GetPesanCal() domain.BookingCore {
	res := bs.qry.GetPesanCal()
	return res
}
