package delivery

import (
	"GunTour/features/gmail/domain"
	"fmt"

	"golang.org/x/oauth2"
)

func SuccessResponse(msg string, url string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    fmt.Sprintf("%v", url),
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

type URL struct {
	Url string `json:"url" form:"url"`
}

func FromDomain(core domain.Code) *oauth2.Token {

	return &oauth2.Token{AccessToken: core.AccessToken, TokenType: core.TokenType, RefreshToken: core.RefreshToken, Expiry: core.Expiry}
}
