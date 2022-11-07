package delivery

import (
	"GunTour/features/ranger/domain"

	"gorm.io/gorm"
)

type ApplyFormat struct {
	Fullname string `json:"fullname" form:"fullname"`
	Phone    string `json:"phone" form:"phone"`
	Ttl      string `json:"ttl" form:"ttl"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
	Docs     string `json:"docs" form:"docs"`
	Price    uint   `json:"price/day" form:"price/day"`
	Detail   string `json:"detail" form:"detail"`
	UserID   uint   `json:"id_user" form:"id_user"`
}

func ToCore(i interface{}) (domain.Core, domain.User) {
	switch i.(type) {
	case ApplyFormat:
		cnv := i.(ApplyFormat)
		return domain.Core{Docs: cnv.Docs, Price: cnv.Price, Detail: cnv.Detail, UserID: cnv.UserID},
			domain.User{Model: gorm.Model{ID: cnv.UserID}, FullName: cnv.Fullname, Phone: cnv.Phone, Dob: cnv.Ttl, Gender: cnv.Gender, Address: cnv.Address}
	}

	return domain.Core{}, domain.User{}
}
