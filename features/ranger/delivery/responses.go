package delivery

import "GunTour/features/ranger/domain"

type ApplyResponse struct {
	ID          uint        `json:"id_ranger" form:"id_ranger"`
	User        domain.User `json:"user" form:"user"`
	Docs        string      `json:"docs" form:"docs"`
	Price       uint        `json:"price/day" form:"price/day"`
	Detail      string      `json:"detail" form:"detail"`
	StatusApply string      `json:"status_apply" form:"status_apply"`
}

type RangerResponse struct {
	ID     uint        `json:"id_ranger" form:"id_ranger"`
	User   domain.User `json:"user" form:"user"`
	Docs   string      `json:"docs" form:"docs"`
	Price  uint        `json:"price/day" form:"price/day"`
	Detail string      `json:"detail" form:"detail"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}

	switch code {
	case "apply":
		cnv := core.(domain.Core)
		res = ApplyResponse{ID: cnv.ID, User: cnv.User, Docs: cnv.Docs, Price: cnv.Price, Detail: cnv.Detail, StatusApply: cnv.StatusApply}
	case "ranger":
		cnv := core.([]domain.Core)
		var arr []RangerResponse

		for _, val := range cnv {
			arr = append(arr, RangerResponse{ID: val.ID, Docs: val.Docs, Price: val.Price, Detail: val.Detail, User: val.User})
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
