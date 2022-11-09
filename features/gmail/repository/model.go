package repository

import (
	"GunTour/features/gmail/domain"
	"time"

	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type Pesan struct {
	gorm.Model
	Email  string
	Status string
}

func FromDomain(dc domain.Code) Code {
	return Code{
		Model:        gorm.Model{ID: dc.ID},
		Code:         dc.Code,
		AccessToken:  dc.AccessToken,
		TokenType:    dc.TokenType,
		RefreshToken: dc.RefreshToken,
		Expiry:       dc.Expiry,
	}
}

func ToDomain(dc Code) domain.Code {
	return domain.Code{
		ID:           dc.ID,
		Code:         dc.Code,
		AccessToken:  dc.AccessToken,
		TokenType:    dc.TokenType,
		RefreshToken: dc.RefreshToken,
		Expiry:       dc.Expiry,
	}
}

func ToDomainPesan(dc Pesan) domain.PesanCore {
	return domain.PesanCore{
		ID:     dc.ID,
		Email:  dc.Email,
		Status: dc.Status,
	}
}
