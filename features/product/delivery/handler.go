package delivery

import (
	"GunTour/features/product/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := productHandler{srv: srv}
	e.GET("/product", handler.ShowAll())
	e.GET("/product/:id_product", handler.ShowByID())
}

func (ph *productHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("page must integer"))
		}

		res, pages, totalPage, err := ph.srv.ShowAll(uint(page))
		if err != nil {
			if strings.Contains(err.Error(), "page") {
				return c.JSON(http.StatusNotFound, FailResponse("page not found."))
			} else if strings.Contains(err.Error(), "no data") {
				return c.JSON(http.StatusOK, SuccessResponseProduct(ToResponseProduct(res, "success get all product", uint(pages), uint(totalPage), "getproduct")))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server."))
		}

		return c.JSON(http.StatusOK, SuccessResponseProduct(ToResponseProduct(res, "success get all product", uint(pages), uint(totalPage), "getproduct")))

	}
}

func (ph *productHandler) ShowByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		productID, err := strconv.Atoi(c.Param("id_product"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("id product must integer"))
		}

		res, err := ph.srv.ShowByID(uint(productID))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server."))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get product detail", ToResponse(res, "detail")))

	}
}
