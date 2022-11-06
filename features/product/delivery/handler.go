package delivery

import (
	"GunTour/features/product/domain"
	"net/http"
	"strconv"

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

		page, _ := strconv.Atoi(c.QueryParam("page"))

		res, pages, totalPage, err := ph.srv.ShowAll(uint(page))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponseProduct(ToResponseProduct(res, "success get all product", int(pages), int(totalPage), "getproduct")))

	}
}

func (ph *productHandler) ShowByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		productID, _ := strconv.Atoi(c.Param("id_product"))

		res, err := ph.srv.ShowByID(uint(productID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get product detail", ToResponse(res, "detail")))

	}
}
