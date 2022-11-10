package delivery

import (
	"GunTour/features/admin/domain"
	"GunTour/utils/middlewares"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	validate = validator.New()
)

type adminHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := adminHandler{srv: srv}
	e.GET("/admin/pendaki", handler.GetPendaki(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                   // GET LIST PENDAKI
	e.POST("/admin/pendaki", handler.AddClimber(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                  // GET LIST PENDAKI
	e.GET("/admin/product", handler.GetProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                   // GET LIST PRODUCT
	e.POST("/admin/product", handler.AddProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                  // ADD NEW PRODUCT
	e.PUT("/admin/product/:id_product", handler.EditProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))      // UPDATE DATA PRODUCT
	e.DELETE("/admin/product/:id_product", handler.RemoveProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // DELETE PRODUCT
	e.GET("/admin/ranger", handler.ShowAllRanger(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                 // GET LIST RANGER
	e.PUT("/admin/ranger/:id_ranger", handler.UpdateRanger(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))       // EDIT RANGER
	e.DELETE("/admin/ranger/:id_ranger", handler.RemoveRanger(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))    // DELETE RANGER
}

func (ah *adminHandler) GetPendaki() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		res, resClimber, err := ah.srv.GetPendaki()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponseProduct(ToResponsePendaki(res, resClimber, "success get list pendaki", "getpendaki")))
	}
}

func (ah *adminHandler) AddClimber() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}

		var input ClimberFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		er := validate.Struct(input)
		if er != nil {
			temp := strings.Split(er.Error(), "Error:")
			return c.JSON(http.StatusBadRequest, FailResponse(temp[1]))
		}

		cnv := ToDomainClimber(input)
		res, err := ah.srv.AddClimber(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server."))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success post climber", ToResponseArray(res, "climber")))
	}
}

func (ah *adminHandler) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("page must integer"))
		}

		res, pages, totalPage, err := ah.srv.GetProduct(page)

		if err != nil {
			if strings.Contains(err.Error(), "page") {
				return c.JSON(http.StatusNotFound, FailResponse("page not found."))
			}
			return c.JSON(http.StatusNotFound, FailResponse("no data."))
		}
		return c.JSON(http.StatusOK, SuccessResponseProduct(ToResponseProduct(res, "success get all product", pages, totalPage, "getproduct")))
	}
}

func (ah *adminHandler) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		er := validate.Struct(input)
		if er != nil {
			temp := strings.Split(er.Error(), "Error:")
			return c.JSON(http.StatusBadRequest, FailResponse(temp[1]))
		}

		file, fileheader, _ := c.Request().FormFile("product_picture")
		_, err := c.FormFile("product_picture")
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("must insert product picture"))
		}

		cnv := ToDomain(input)
		res, err := ah.srv.AddProduct(cnv, file, fileheader)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success add product", ToResponse(res, "addproduct")))
	}
}

func (ah *adminHandler) EditProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat

		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}

		id, err := strconv.Atoi(c.Param("id_product"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("invalid id"))
		}
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		file, fileheader, _ := c.Request().FormFile("product_picture")
		if input.ProductName == "" && input.Detail == "" && input.RentPrice == 0 && file == nil {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
		}

		cnv := ToDomain(input)
		res, err := ah.srv.EditProduct(cnv, file, fileheader)

		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("data not found."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success update product", ToResponse(res, "update")))
	}
}

func (ah *adminHandler) RemoveProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		id, er := strconv.Atoi(c.Param("id_product"))
		if er != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("invalid id"))
		}
		err := ah.srv.RemoveProduct(id)

		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("data not found."))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("success delete product"))
	}
}

func (ah *adminHandler) ShowAllRanger() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}

		resAccepted, res, err := ah.srv.ShowAllRanger()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponseRanger("success get ranger", ToResponseArray(resAccepted, "ranger"), "success get applied ranger", ToResponseArray(res, "rangerapply")))
	}
}

func (ah *adminHandler) UpdateRanger() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}

		rangerId, err := strconv.Atoi(c.Param("id_ranger"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("page must integer"))
		}

		var input RangerFormat

		input.ID = uint(rangerId)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
		}

		validApply := input.StatusApply == "accepted" || input.StatusApply == "rejected"
		validStatus := input.Status == "off" || input.Status == "duty" || input.Status == "available"

		if input.Phone == "" {
			if !validApply && !validStatus {
				return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
			}
		}

		cnv := ToDomainRanger(input)
		conv := ToDomainUser(input)
		res, resU, err := ah.srv.UpdateRanger(cnv, conv, uint(rangerId))

		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server"))
		}
		// if input.StatusApply != "" {
		// c.Redirect(http.StatusTemporaryRedirect, "/gmail/send")
		// helper.Openbrowser("localhost:8000/gmail")
		// }

		return c.JSON(http.StatusAccepted, SuccessResponse("success update status ranger", ToResponseUser(res, resU, "ranger")))
	}
}

func (ah *adminHandler) RemoveRanger() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		id, er := strconv.Atoi(c.Param("id_ranger"))
		if er != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("invalid id"))
		}
		err := ah.srv.RemoveRanger(id)

		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("data not found."))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("success delete ranger"))
	}
}
