package delivery

import "GunTour/features/product/domain"

type ProductResponse struct {
	ID             uint   `json:"id_product" form:"id_product"`
	ProductName    string `json:"product_name" form:"product_name"`
	RentPrice      int    `json:"rent_price" form:"rent_price"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type ProdResponse struct {
	ID             uint   `json:"id_product" form:"id_product"`
	ProductName    string `json:"product_name" form:"product_name"`
	RentPrice      int    `json:"rent_price" form:"rent_price"`
	Detail         string `json:"detail" form:"detail"`
	Note           string `json:"note" form:"note"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type GetProductResponse struct {
	Message   string            `json:"message" form:"message"`
	Page      uint              `json:"page" form:"page"`
	TotalPage uint              `json:"total_page" form:"total_page"`
	Data      []ProductResponse `json:"data" form:"data"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "getproduct":
		var arr []ProductResponse
		val := core.([]domain.Core)
		for _, cnv := range val {
			arr = append(arr, ProductResponse{ID: uint(cnv.ID), ProductName: cnv.ProductName, RentPrice: cnv.RentPrice, ProductPicture: cnv.ProductPicture})
		}
		res = arr
	case "detail":
		cnv := core.(domain.Core)
		res = ProdResponse{ID: cnv.ID, ProductName: cnv.ProductName, RentPrice: cnv.RentPrice, Detail: cnv.Detail, Note: cnv.Note, ProductPicture: cnv.ProductPicture}
	}
	return res
}

func ToResponseProduct(core interface{}, message string, pages uint, totalPage uint, code string) interface{} {
	var arr []ProductResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, ProductResponse{ID: uint(cnv.ID), ProductName: cnv.ProductName, RentPrice: cnv.RentPrice, ProductPicture: cnv.ProductPicture})
	}

	return GetProductResponse{Message: message, Page: pages, TotalPage: totalPage, Data: arr}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func SuccessResponseProduct(data interface{}) interface{} {
	return data
}
