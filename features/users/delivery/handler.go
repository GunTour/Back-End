package delivery

import (
	"GunTour/features/users/domain"
	"GunTour/utils/helper"
	"GunTour/utils/middlewares"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	validate = validator.New()
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.POST("/users", handler.Insert())
	e.PUT("/users", handler.Update(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.DELETE("/users", handler.Delete(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.POST("/login", handler.Login())
}

func (uh *userHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		er := validate.Struct(input)
		if er != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
		}

		valid := helper.Password(input.Password)
		if valid != "Valid" {
			return c.JSON(http.StatusBadRequest, FailResponse(valid))
		}
		input.UserPicture = "https://guntour.s3.ap-southeast-1.amazonaws.com/posts/iTs1Ve2IJ71i6wSGzMBp-profile.jpg"
		cnv := ToCore(input)
		res, err := uh.srv.Insert(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success register user", ToResponse(res, "register")))

	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role := middlewares.ExtractToken(c)

		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else if role == "pendaki" {
			var input UpdateFormat

			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
			}

			if input.FullName != "" {
				var name FullName
				name.FullName = input.FullName

				err := validate.Struct(name)
				if err != nil {
					return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
				}
			}

			if input.Email != "" {
				var ema Email
				ema.Email = input.Email

				err := validate.Struct(ema)
				if err != nil {
					return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
				}
			}

			if input.Password != "" {
				valid := helper.Password(input.Password)
				if valid != "Valid" {
					return c.JSON(http.StatusBadRequest, FailResponse(valid))
				}
			}

			file, fileheader, _ := c.Request().FormFile("user_picture")

			cnv := ToCore(input)
			res, err := uh.srv.Update(cnv, file, fileheader, userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusAccepted, SuccessResponse("success update user", ToResponse(res, "update")))
		} else {
			return c.JSON(http.StatusUnauthorized, FailResponse("wrong role"))
		}
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role := middlewares.ExtractToken(c)

		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else if role == "pendaki" {
			res, err := uh.srv.Delete(userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusOK, SuccessResponse("success delete data", ToResponse(res, "")))
		} else {
			return c.JSON(http.StatusUnauthorized, FailResponse("wrong role"))
		}
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success login", ToResponse(res, "login")))
	}
}
