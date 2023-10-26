package products

import (
	"net/http"
	"simple-echo-backend/common"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
	products, err := GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewGetProductListResponse(products))
}

func GetProductByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid Id"))
	}

	// if result.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	product, err := GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewGetProductResponse(product))
}

func CreateProductController(c echo.Context) error {
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
	err := CreateProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewSuccessPayload())
}

func UpdateProductController(c echo.Context) error {
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
	product, err := GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	// if resProduct.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	updateProduct := Products{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   time.Now(),
		// DeletedAt:   product.DeletedAt,
	}
	err = UpdateProduct(updateProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}

	err = DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}
