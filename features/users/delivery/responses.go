package delivery

import (
	"GunTour/features/users/domain"
)

type RegisterResponse struct {
	ID       int    `json:"id_user" form:"id_user"`
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
}

type UpdateResponse struct {
	ID          int    `json:"id_user" form:"id_user"`
	FullName    string `json:"fullname" form:"fullname"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
	UserPicture string `json:"user_picture" form:"user_picture"`
}

type LoginResponse struct {
	ID          int    `json:"id_user" form:"id_user"`
	FullName    string `json:"fullname" form:"fullname"`
	Email       string `json:"email" form:"email"`
	Role        string `json:"role" form:"role"`
	UserPicture string `json:"user_picture" form:"user_picture"`
	Token       string `json:"token" form:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}

	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, FullName: cnv.FullName, Email: cnv.Email, Role: cnv.Role}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, FullName: cnv.FullName, Email: cnv.Email, Phone: cnv.Phone, UserPicture: cnv.UserPicture}
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{ID: cnv.ID, FullName: cnv.FullName, Email: cnv.Email, Role: cnv.Role, UserPicture: cnv.UserPicture, Token: cnv.Token}
	}

	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
