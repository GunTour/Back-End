package delivery

import (
	"GunTour/features/gmail/domain"

	"golang.org/x/oauth2"
)

type RegisterFormat struct {
	Code string `json:"code" form:"code"`
}

type Pesan struct {
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
}

type UserInfoFormat struct {
	Email        string `json:"email"`
	Fullname     string `json:"name"`
	Photoprofile string `json:"picture"`
}

func ToDomain(core *oauth2.Token, code string) domain.Code {
	return domain.Code{Code: code, AccessToken: core.AccessToken, TokenType: core.TokenType, RefreshToken: core.RefreshToken, Expiry: core.Expiry}
}
