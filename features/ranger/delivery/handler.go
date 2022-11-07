package delivery

import (
	"GunTour/features/ranger/domain"
	"GunTour/utils/middlewares"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type rangerHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := rangerHandler{srv: srv}
	e.POST("/ranger", handler.Apply(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("/ranger", handler.ShowAll(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (rh *rangerHandler) Apply() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, role := middlewares.ExtractToken(c)

		if userID == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))
		} else if role == "pendaki" {
			var input ApplyFormat
			input.UserID = uint(userID)

			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}

			file, fileheader, _ := c.Request().FormFile("docs")

			// if input.Docs == "" {
			// 	return c.JSON(http.StatusBadRequest, FailResponse("must submit document"))
			// }

			if input.Price == 0 {
				return c.JSON(http.StatusBadRequest, FailResponse("must submit rate per day"))
			}

			cnv := ToCore(input)
			res, err := rh.srv.Apply(cnv, file, fileheader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusCreated, SuccessResponse("success to apply as ranger", ToResponse(res, "apply")))

		} else {
			return c.JSON(http.StatusUnauthorized, FailResponse("unauthorized or wrong role"))
		}
	}
}

func (rh *rangerHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, role := middlewares.ExtractToken(c)

		if userID == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot validate token"))

		} else if role == "pendaki" {
			Start, _ := time.Parse("2021-02-04", c.QueryParam("date_start"))
			End, _ := time.Parse("2021-02-04", c.QueryParam("date_start"))
			res, err := rh.srv.ShowAll(Start, End)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse("something went wrong"))
			}
			return c.JSON(http.StatusOK, SuccessResponse("success get ranger", ToResponse(res, "ranger")))

		} else {
			return c.JSON(http.StatusUnauthorized, FailResponse("unauthorized or wrong role"))
		}
	}
}
