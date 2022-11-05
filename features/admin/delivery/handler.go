package delivery

import (
	"GunTour/features/admin/domain"
	"GunTour/utils/middlewares"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type adminHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := adminHandler{srv: srv}
	e.GET("/admin/pendaki", handler.GetPendaki(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // GET BOOKING
	e.GET("/admin/booking", handler.GetBooking(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // GET BOOKING DETAIL
}

func (ah *adminHandler) GetPendaki() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		res, err := ah.srv.GetPendaki()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get list pendaki", ToResponseArray(res, "getpendaki")))
	}
}

func (ah *adminHandler) GetBooking() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		res, err := ah.srv.GetBooking()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success show all booking data", ToResponseArray(res, "getbooking")))
	}
}
