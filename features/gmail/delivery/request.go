package delivery

import (
	"GunTour/features/gmail/domain"

	"golang.org/x/oauth2"
)

type RegisterFormat struct {
	Code string `json:"code" form:"code"`
}

func ToDomain(core *oauth2.Token, code string) domain.Code {
	return domain.Code{Code: code, AccessToken: core.AccessToken, TokenType: core.TokenType, RefreshToken: core.RefreshToken, Expiry: core.Expiry}
}
