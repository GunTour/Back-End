package delivery

import (
	"GunTour/features/ranger/domain"
	"GunTour/utils/middlewares"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	validate = validator.New()
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

			_, err := c.FormFile("docs")
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("must insert document"))
			}
			file, fileheader, _ := c.Request().FormFile("docs")

			er := validate.Struct(input)
			if er != nil {
				temp := strings.Split(er.Error(), "Error:")
				return c.JSON(http.StatusBadRequest, FailResponse(temp[1]))
			}
			cnv, cnvUser := ToCore(input)
			res, err := rh.srv.Apply(cnv, cnvUser, file, fileheader)
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
			Start, err := time.Parse("2006-01-02", c.QueryParam("date_start"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
			}
			End, err := time.Parse("2006-01-02", c.QueryParam("date_end"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
			}

			Start.AddDate(0, 0, -1)
			End.AddDate(0, 0, 1)

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
