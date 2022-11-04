package delivery

import (
	"GunTour/features/booking/domain"
	"GunTour/utils/middlewares"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type bookingHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := bookingHandler{srv: srv}
	e.POST("/booking", handler.InsertData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))               // TAMBAH BOOKING
	e.GET("/booking", handler.GetAll(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                    // GET BOOKING
	e.GET("/booking/:id_booking", handler.GetDetail(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))     // GET BOOKING DETAIL
	e.DELETE("/booking/:id_booking", handler.DeleteData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // DELETE BOOKING
	e.PUT("/booking/:id_booking", handler.UpdateData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))    // UPDATE BOOKING
	e.GET("/booking/ranger", handler.GetRangerBooking(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))   // GET RANGER BOOKING
}

func (bs *bookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middlewares.ExtractToken(c)
		res, err := bs.srv.GetAll(uint(id))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking history", ToResponseArray(res, "getall")))
	}
}

func (bs *bookingHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		idBooking, err := strconv.Atoi(c.Param("id_booking"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("failed to get id booking"))
		}
		res, err := bs.srv.GetDetail(uint(idBooking))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking detail", ToResponse(res, "update")))
	}
}

func (bs *bookingHandler) GetRangerBooking() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middlewares.ExtractToken(c)
		res, err := bs.srv.GetRangerBooking(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking ranger", ToResponseArray(res, "getranger")))
	}
}

func (bs *bookingHandler) InsertData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}
		temp := uuid.New()
		input.OrderId = "Order-" + temp.String()
		input.StatusBooking = "unpaid"
		IdUser, _ := middlewares.ExtractToken(c)
		input.IdUser = uint(IdUser)
		cnv := ToDomain(input)
		res, err := bs.srv.InsertData(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add booking plan", ToResponse(res, "register")))
	}
}

func (bs *bookingHandler) UpdateData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id, err := strconv.Atoi(c.Param("id_booking"))
		if err != nil {
			return c.JSON(http.StatusBadGateway, FailResponse("failed to get id booking"))
		}
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		idUser, _ := middlewares.ExtractToken(c)
		input.IdUser = uint(idUser)
		cnv := ToDomain(input)
		res, err := bs.srv.UpdateData(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success edit booking plan", ToResponse(res, "update")))
	}
}

func (bs *bookingHandler) DeleteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("failed to get id booking")
		}
		log.Print(id)
		err = bs.srv.DeleteData(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("success delete data."))
	}
}
