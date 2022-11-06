package delivery

import (
	"GunTour/features/ranger/domain"
)

type ApplyFormat struct {
	Docs   string `json:"docs" form:"docs"`
	Price  uint   `json:"price/day" form:"price/day"`
	Detail string `json:"detail" form:"detail"`
	UserID uint   `json:"id_user" form:"id_user"`
}

func ToCore(i interface{}) domain.Core {
	switch i.(type) {
	case ApplyFormat:
		cnv := i.(ApplyFormat)
		return domain.Core{Docs: cnv.Docs, Price: cnv.Price, Detail: cnv.Detail, UserID: cnv.UserID}
	}

	return domain.Core{}
}
