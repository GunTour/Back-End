package delivery

import "GunTour/features/climber/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type ClimberResponse struct {
	IsClimber     int `json:"is_climber" form:"is_climber"`
	MaleClimber   int `json:"male_climber" form:"male_climber"`
	FemaleClimber int `json:"female_climber" form:"female_climber"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "climber":
		cnv := core.(domain.Core)
		res = ClimberResponse{IsClimber: cnv.IsClimber, MaleClimber: cnv.MaleClimber, FemaleClimber: cnv.FemaleClimber}
	}

	return res
}
