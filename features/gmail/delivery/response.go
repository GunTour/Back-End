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

func SuccessResponseRanger(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
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

type RangerAccepted struct {
	ID          uint   `json:"id_ranger" form:"id_ranger"`
	Status      string `json:"status" form:"status"`
	StatusApply string `json:"status_apply" form:"status_apply"`
}

type RangerAcceptedNoMail struct {
	ID          uint   `json:"id_ranger" form:"id_ranger"`
	Status      string `json:"status" form:"status"`
	StatusApply string `json:"status_apply" form:"status_apply"`
	Url         string `json:"email_sender" form:"email_sender"`
}

func FromDomain(core domain.Code) *oauth2.Token {

	return &oauth2.Token{AccessToken: core.AccessToken, TokenType: core.TokenType, RefreshToken: core.RefreshToken, Expiry: core.Expiry}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "ranger":
		cnv := core.(domain.RangerCore)
		res = RangerAccepted{ID: cnv.ID, Status: cnv.Status, StatusApply: cnv.StatusApply}
	}

	return res
}

func ToResponseGagal(core interface{}, authURL string, code string) interface{} {
	var res interface{}
	switch code {
	case "ranger":
		cnv := core.(domain.RangerCore)
		res = RangerAcceptedNoMail{ID: cnv.ID, Status: cnv.Status, StatusApply: cnv.StatusApply, Url: authURL}
	}

	return res
}
