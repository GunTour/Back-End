package delivery

import (
	"GunTour/features/admin/domain"
	"GunTour/utils/middlewares"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type adminHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := adminHandler{srv: srv}
	e.GET("/admin/pendaki", handler.GetPendaki(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                   // GET LIST PENDAKI
	e.GET("/admin/booking", handler.GetBooking(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                   // GET LIST BOOKING
	e.GET("/admin/product", handler.GetProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                   // GET LIST PRODUCT
	e.POST("/admin/product", handler.AddProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))                  // ADD NEW PRODUCT
	e.PUT("/admin/product/:id_product", handler.EditProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))      // UPDATE DATA PRODUCT
	e.DELETE("/admin/product/:id_product", handler.RemoveProduct(), middleware.JWT([]byte(os.Getenv("JWT_SECRET")))) // DELETE PRODUCT
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

func (ah *adminHandler) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := middlewares.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, FailResponse("jangan macam-macam, anda bukan admin"))
		}
		page, _ := strconv.Atoi(c.QueryParam("page"))

		res, pages, totalPage, err := ah.srv.GetProduct(page)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
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

		file, fileheader, _ := c.Request().FormFile("product_picture")

		cnv := ToDomain(input)
		res, err := ah.srv.AddProduct(cnv, file, fileheader)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
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

		id, _ := strconv.Atoi(c.Param("id_product"))
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}

		file, fileheader, _ := c.Request().FormFile("product_picture")

		cnv := ToDomain(input)
		res, err := ah.srv.EditProduct(cnv, file, fileheader)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
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
		id, _ := strconv.Atoi(c.Param("id_product"))
		err := ah.srv.RemoveProduct(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("success delete product"))
	}
}
