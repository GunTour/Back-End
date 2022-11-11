package delivery

import (
	"GunTour/features/booking/domain"
	"GunTour/utils/middlewares"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

var (
	validate = validator.New()
)

type bookingHandler struct {
	srv   domain.Services
	oauth *oauth2.Config
}

func New(e *echo.Echo, srv domain.Services) {
	handler := bookingHandler{srv: srv}
	e.POST("/booking", handler.InsertData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))               // TAMBAH BOOKING
	e.GET("/booking", handler.GetAll(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                    // GET BOOKING
	e.GET("/booking/:id_booking", handler.GetDetail(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))     // GET BOOKING DETAIL
	e.DELETE("/booking/:id_booking", handler.DeleteData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // DELETE BOOKING
	e.PUT("/booking/:id_booking", handler.UpdateData(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))    // UPDATE BOOKING
	e.GET("/booking/ranger", handler.GetRangerBooking(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))   // GET RANGER BOOKING
	e.POST("/midtrans", handler.UpdateMidtrans())                                                           // CALLBACK UPDATE MIDTRANS
}

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

func (bs *bookingHandler) GetRangerBooking() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middlewares.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role == "ranger" {
			return c.JSON(http.StatusUnauthorized, FailResponse("unauthorized access detected"))
		}

		res, err := bs.srv.GetRangerBooking(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get booking ranger", ToResponseArray(res, "getranger")))
	}
}

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

		temp := uuid.New()
		input.OrderId = "Order-" + temp.String()
		input.IdUser = uint(IdUser)
		cnv := ToDomain(input)
		res, err := bs.srv.InsertData(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "datetime") {
				c.JSON(http.StatusBadRequest, FailResponse("must input date time, or wrong input"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server"))
		}

		if input.DateStart != "" && input.DateEnd != "" && input.Ticket != 0 {
			c.Redirect(http.StatusTemporaryRedirect, "/calendar/send")
			// helper.Openbrowser("localhost:8000/gmail")
		}

		// if input.Token != "" {
		// 	pendaki, ranger := bs.srv.GetEmail(IdUser, int(input.IdRanger))
		// 	if pendaki.Email == "" {
		// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 			"code":    500,
		// 			"message": "Internal Server Error",
		// 		})
		// 	}

		// 	Start := fmt.Sprintf("%sT07:20:50.52Z", input.DateStart)
		// 	End := fmt.Sprintf("%sT07:20:50.52Z", input.DateEnd)

		// 	events := &calendar.Event{
		// 		Summary:     "Day of Your Climbing",
		// 		Description: "Climbing Day",
		// 		Start: &calendar.EventDateTime{
		// 			DateTime: Start,
		// 			TimeZone: "Asia/Jakarta",
		// 		},
		// 		End: &calendar.EventDateTime{
		// 			DateTime: End,
		// 			TimeZone: "Asia/Jakarta",
		// 		},
		// 		Attendees: []*calendar.EventAttendee{
		// 			{Email: pendaki.Email},
		// 			{Email: ranger.Email},
		// 		},
		// 	}

		// 	tokenOauth := &oauth2.Token{AccessToken: input.Token}

		// 	client := bs.oauth.Client(c.Request().Context(), tokenOauth)

		// 	srv, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
		// 	if err != nil {
		// 		log.Printf("Unable to retrieve Calendar client: %v", err)
		// 	}

		// 	_, err = srv.Events.Insert("primary", events).SendUpdates("all").Do()
		// 	if err != nil {
		// 		log.Printf("Unable to create event. %v\n", err)
		// 	}
		// }

		return c.JSON(http.StatusCreated, SuccessResponse("success add booking plan", ToResponse(res, "getdetails")))
	}
}

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
			log.Print(input)
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		if input.DateStart == "" && input.DateEnd == "" && input.Entrance == "" && input.Ticket == 0 && input.Product == nil && input.GrossAmount == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
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

		return c.JSON(http.StatusCreated, SuccessResponse("success edit booking plan", ToResponse(res, "getdetails")))
	}
}

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

func (bs *bookingHandler) UpdateMidtrans() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateMidtrans
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}

		res := ToDomain(input)
		bs.srv.UpdateMidtrans(res)
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success update data."))
	}
}
