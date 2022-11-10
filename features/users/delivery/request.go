package delivery

import "GunTour/features/users/domain"

type RegisterFormat struct {
	FullName    string `json:"fullname" form:"fullname" validate:"required,min=4,max=30"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password"`
	UserPicture string
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateFormat struct {
	FullName    string `json:"fullname" form:"fullname" validate:"min=4,max=30"`
	Email       string `json:"email" form:"email" validate:"email"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone" validate:"min=11,max=15"`
	Address     string `json:"address" form:"address"`
	Dob         string `json:"dob" form:"dob"`
	Gender      string `json:"gender" form:"gender"`
	UserPicture string `json:"user_picture" form:"user_picture"`
}

type FullName struct {
	FullName string `json:"fullname" form:"fullname" validate:"min=4,max=30"`
}

type Email struct {
	Email string `json:"email" form:"email" validate:"email"`
}

func ToCore(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{FullName: cnv.FullName, Email: cnv.Email, Password: cnv.Password, UserPicture: "https://guntour.s3.ap-southeast-1.amazonaws.com/posts/iTs1Ve2IJ71i6wSGzMBp-profile.jpg"}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{FullName: cnv.FullName, Email: cnv.Email, Password: cnv.Password, Phone: cnv.Phone, Address: cnv.Address, Dob: cnv.Dob, Gender: cnv.Gender, UserPicture: cnv.UserPicture}
	}

	return domain.Core{}
}
