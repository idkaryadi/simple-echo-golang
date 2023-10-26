package handler

import (
	"net/http"
	"simple-echo-backend/common"
	"simple-echo-backend/products/repository"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	products, err := repository.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewGetProductListResponse(products))
}

func GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid Id"))
	}

	// if result.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	product, err := repository.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewGetProductResponse(product))
}

func CreateProduct(c echo.Context) error {
	req := new(ProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	if err := common.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	newProduct := repository.Products{
		Name:        req.Name,
		Description: req.Description,
	}
	err := repository.CreateProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewSuccessPayload())
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	req := new(ProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	if err := common.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	// get product
	product, err := repository.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	// if resProduct.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	updateProduct := repository.Products{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   time.Now(),
		// DeletedAt:   product.DeletedAt,
	}
	err = repository.UpdateProduct(updateProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}

	err = repository.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}
