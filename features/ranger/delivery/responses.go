package delivery

import "GunTour/features/ranger/domain"

type ApplyResponse struct {
	ID          uint        `json:"id_ranger" form:"id_ranger"`
	User        domain.User `json:"-" form:"-"`
	IdUser      uint        `json:"id_user" form:"id_user"`
	Fullname    string      `json:"fullname" form:"fullname"`
	Ttl         string      `json:"ttl" form:"ttl"`
	Address     string      `json:"address" form:"address"`
	Phone       string      `json:"phone" form:"phone"`
	Gender      string      `json:"gender" form:"gender"`
	Docs        string      `json:"docs" form:"docs"`
	Price       uint        `json:"price/day" form:"price/day"`
	Detail      string      `json:"detail" form:"detail"`
	StatusApply string      `json:"status_apply" form:"status_apply"`
}

type RangerResponse struct {
	ID          uint        `json:"id_ranger" form:"id_ranger"`
	User        domain.User `json:"-" form:"-"`
	Fullname    string      `json:"fullname" form:"fullname"`
	Price       uint        `json:"price/day" form:"price/day"`
	Detail      string      `json:"detail" form:"detail"`
	UserPicture string      `json:"user_picture" form:"user_picture"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}

	switch code {
	case "apply":
		cnv := core.(domain.Core)
		res = ApplyResponse{ID: cnv.ID, IdUser: cnv.User.ID, Fullname: cnv.User.FullName, Ttl: cnv.User.Dob, Address: cnv.User.Address,
			Phone: cnv.User.Phone, Gender: cnv.User.Gender, Docs: cnv.Docs, Price: cnv.Price, Detail: cnv.Detail, StatusApply: cnv.StatusApply}
	case "ranger":
		cnv := core.([]domain.Core)
		var arr []RangerResponse

		for _, val := range cnv {
			arr = append(arr, RangerResponse{ID: val.ID, Price: val.Price, Detail: val.Detail, Fullname: val.User.FullName, UserPicture: val.User.UserPicture})
		}
		res = arr
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
