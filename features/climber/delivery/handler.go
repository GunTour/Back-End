package delivery

import (
	"GunTour/features/climber/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type climberHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := climberHandler{srv: srv}
	e.GET("/climber", handler.ShowClimber()) // GET LIST PENDAKI
}

// HANDLER TO SHOW CLIMBER DATA
func (ch *climberHandler) ShowClimber() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ch.srv.ShowClimber()

		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("no data."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get climber", ToResponse(res, "climber")))
	}
}
