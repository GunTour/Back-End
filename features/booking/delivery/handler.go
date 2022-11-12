package delivery

import (
	"GunTour/features/booking/domain"
	"GunTour/utils/helper"
	"GunTour/utils/middlewares"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	validate = validator.New()
)

type bookingHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := bookingHandler{srv: srv}
	e.POST("/booking", handler.InsertData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))               // TAMBAH BOOKING
	e.GET("/booking", handler.GetAll(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                    // SHOW ALL USER'S BOOKING DATA
	e.GET("/booking/:id_booking", handler.GetDetail(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))     // GET BOOKING DETAIL
	e.DELETE("/booking/:id_booking", handler.DeleteData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // DELETE BOOKING
	e.PUT("/booking/:id_booking", handler.UpdateData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))    // UPDATE BOOKING
	e.GET("/booking/ranger", handler.GetRangerBooking(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))   // GET RANGER BOOKING
	e.POST("/midtrans", handler.UpdateMidtrans())                                                           // CALLBACK UPDATE MIDTRANS
}

// SHOW ALL USER'S BOOKING DATA
func (bs *bookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middlewares.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "pendaki" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unaothorized access detected"))
		}
		res, err := bs.srv.GetAll(uint(id))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking history", ToResponseArray(res, "getall")))
	}
}

// SHOW USER'S DETAIL BOOKING
func (bs *bookingHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middlewares.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "pendaki" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unaothorized access detected"))
		}

		idBooking, err := strconv.Atoi(c.Param("id_booking"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id booking must integer"))
		}
		res, err := bs.srv.GetDetail(uint(idBooking))
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("data not found"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking detail", ToResponse(res, "getdetails")))
	}
}

// SHOW LIST RANGER'S BOOKING DATA
func (bs *bookingHandler) GetRangerBooking() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middlewares.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "ranger" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unauthorized access detected"))
		}

		res, err := bs.srv.GetRangerBooking(uint(id))
		if err != nil {
			if strings.Contains(err.Error(), "no data") {
				return c.JSON(http.StatusNotFound, SuccessResponse("no data.", ToResponseArray(res, "getranger")))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking ranger", ToResponseArray(res, "getranger")))
	}
}

// ADD BOOKING
func (bs *bookingHandler) InsertData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat

		IdUser, role := middlewares.ExtractToken(c)
		if IdUser == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "pendaki" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unaothorized access detected"))
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		er := validate.Struct(input)
		if er != nil {
			if strings.Contains(er.Error(), "entrance") {
				return c.JSON(http.StatusBadRequest, FailResponse("must fill required field entrance"))
			} else if strings.Contains(er.Error(), "ticket") {
				return c.JSON(http.StatusBadRequest, FailResponse("must fill required field ticket"))
			} else if strings.Contains(er.Error(), "grossamount") {
				return c.JSON(http.StatusBadRequest, FailResponse("must fill required field gross_amount"))
			} else if strings.Contains(er.Error(), "date") {
				c.JSON(http.StatusBadRequest, FailResponse("must input date time, or wrong input"))
			} else {
				return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
			}
		}

		input.DateStart, er = time.Parse("2006-01-02", input.Start)
		if er != nil {
			c.JSON(http.StatusBadRequest, FailResponse("wrong input date start"))
		}

		input.DateEnd, _ = time.Parse("2006-01-02", input.End)
		if er != nil {
			c.JSON(http.StatusBadRequest, FailResponse("wrong input date end"))
		}

		temp := uuid.New()
		input.OrderId = "Order-" + temp.String()
		input.IdUser = uint(IdUser)
		cnv := ToDomain(input)
		res, err := bs.srv.InsertData(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "datetime") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input date time, or wrong input"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server"))
		}

		if input.Start != "" && input.End != "" && input.Ticket != 0 {
			resCode, err := bs.srv.GetCode()
			if err != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/calendar/send")
			}

			err = helper.EventCalendar(resCode, res)
			if err != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/calendar/send")
			}
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add booking plan", ToResponse(res, "getdetails")))
	}
}

// UPDATE BOOKING DATA
func (bs *bookingHandler) UpdateData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		IdUser, role := middlewares.ExtractToken(c)
		if IdUser == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "pendaki" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unaothorized access detected"))
		}

		id, err := strconv.Atoi(c.Param("id_booking"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id booking must integer"))
		}
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		if input.End == "" && input.Start == "" && input.Entrance == "" && input.Ticket == 0 && input.Product == nil && input.GrossAmount == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
		}

		if input.Start != "" {
			input.DateStart, err = time.Parse("2006-01-02", input.Start)
			if err != nil {
				c.JSON(http.StatusBadRequest, FailResponse("wrong input date start"))
			}
		}
		if input.End != "" {
			input.DateEnd, err = time.Parse("2006-01-02", input.End)
			if err != nil {
				c.JSON(http.StatusBadRequest, FailResponse("wrong input date end"))
			}
		}
		input.IdUser = uint(IdUser)
		cnv := ToDomain(input)

		res, err := bs.srv.UpdateData(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "data") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success edit booking plan", ToResponse(res, "getdetails")))
	}
}

// DELETE BOOKING DATA BY ID
func (bs *bookingHandler) DeleteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id_booking"))
		IdUser, role := middlewares.ExtractToken(c)
		if IdUser == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role != "pendaki" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unaothorized access detected"))
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id booking must integer"))
		}

		err = bs.srv.DeleteData(uint(id))

		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("data not found"))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("success delete data."))
	}
}

// CALLBACK MIDTRANS
func (bs *bookingHandler) UpdateMidtrans() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateMidtrans
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}

		res := ToDomain(input)
		bs.srv.UpdateMidtrans(res)
		return c.JSON(http.StatusAccepted, SuccessResponseNoData("Success update data."))
	}
}
