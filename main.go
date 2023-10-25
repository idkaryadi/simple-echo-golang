package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	connectDB()
	e := echo.New()
	e.GET("/v1/products", GetProduct)
	e.GET("/v1/products/:id", GetProductById)
	e.POST("/v1/products", CreateProduct)
	e.PUT("/v1/products/:id", UpdateProduct)
	e.DELETE("/v1/products/:id", DeleteProduct)

	e.Logger.Fatal(e.Start(":4000"))
}

func GetProduct(c echo.Context) error {
	var products []Products
	result := DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", result.Error.Error()))
	}
	return c.JSON(http.StatusOK, NewGetProductListResponse(products))
}

func GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("error", "Invalid Id"))
	}
	product := Products{ID: id}
	result := DB.Find(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", result.Error.Error()))
	}
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, NewErrorResponse("error", "data not found"))
	}

	return c.JSON(http.StatusOK, NewGetProductResponse(product))
}

func CreateProduct(c echo.Context) error {
	req := new(ProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("error", err.Error()))
	}
	newProduct := Products{
		Name:        req.Name,
		Description: req.Description,
	}
	result := DB.Create(&newProduct)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", result.Error.Error()))
	}
	return c.JSON(http.StatusCreated, NewSuccessPayload())
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("error", err.Error()))
	}
	req := new(ProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("error", err.Error()))
	}
	// get product
	product := Products{ID: id}
	resProduct := DB.Find(&product)
	if resProduct.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", err.Error()))
	}
	if resProduct.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, NewErrorResponse("error", "data not found"))
	}
	updateProduct := Products{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   time.Now(),
		// DeletedAt:   product.DeletedAt,
	}
	result := DB.Save(&updateProduct)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", result.Error.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("error", err.Error()))
	}
	// Note: masih hard delete
	result := DB.Delete(&Products{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", result.Error.Error()))
	}
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse("error", "data not found"))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}
