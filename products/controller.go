package products

import (
	"net/http"
	"simple-echo-backend/common"
	"simple-echo-backend/config"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	var products []Products
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(result.Error.Error()))
	}
	return c.JSON(http.StatusOK, NewGetProductListResponse(products))
}

func GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid Id"))
	}
	product := Products{ID: id}
	result := config.DB.Find(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(result.Error.Error()))
	}
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
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
	newProduct := Products{
		Name:        req.Name,
		Description: req.Description,
	}
	result := config.DB.Create(&newProduct)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(result.Error.Error()))
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
	product := Products{ID: id}
	resProduct := config.DB.Find(&product)
	if resProduct.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	if resProduct.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	}
	updateProduct := Products{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   time.Now(),
		// DeletedAt:   product.DeletedAt,
	}
	result := config.DB.Save(&updateProduct)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(result.Error.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	// Note: masih hard delete
	result := config.DB.Delete(&Products{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(result.Error.Error()))
	}
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("data not found"))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}
