package delivery

import (
	"GunTour/features/ranger/domain"
	"time"

	"gorm.io/gorm"
)

type ApplyFormat struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Ttl      string `json:"ttl" form:"ttl" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Docs     string `json:"docs" form:"docs"`
	Price    uint   `json:"price" form:"price"`
	Detail   string `json:"detail" form:"detail"`
	UserID   uint   `json:"id_user" form:"id_user"`
}

type GetFormat struct {
	StartDate time.Time `validate:"required"`
	EndDate   time.Time `validate:"required"`
}

func ToCore(i interface{}) (domain.Core, domain.User) {
	switch i.(type) {
	case ApplyFormat:
		cnv := i.(ApplyFormat)
		return domain.Core{Docs: cnv.Docs, Price: 100000, Detail: cnv.Detail, UserID: cnv.UserID},
			domain.User{Model: gorm.Model{ID: cnv.UserID}, FullName: cnv.Fullname, Phone: cnv.Phone, Dob: cnv.Ttl, Gender: cnv.Gender, Address: cnv.Address}
	}

	return domain.Core{}, domain.User{}
}
