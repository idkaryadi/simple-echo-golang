package handler

import (
	"net/http"
	"simple-echo-backend/common"
	"simple-echo-backend/products/core"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo core.Repository
}

func NewHandler(repo core.Repository) Handler {
	return Handler{repo: repo}
}

func (h Handler) GetProduct(c echo.Context) error {
	products, err := h.repo.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewGetProductListResponse(products))
}

func (h Handler) GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid Id"))
	}

	// if result.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	product, err := h.repo.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewGetProductResponse(product))
}

func (h Handler) CreateProduct(c echo.Context) error {
	req := new(ProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	if err := common.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}
	newProduct := core.Products{
		Name:        req.Name,
		Description: req.Description,
	}
	err := h.repo.CreateProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, NewSuccessPayload())
}

func (h Handler) UpdateProduct(c echo.Context) error {
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
	product, err := h.repo.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	// if resProduct.RowsAffected < 1 {
	// 	return c.JSON(http.StatusNotFound, NewErrorResponse("data not found"))
	// }
	updateProduct := core.Products{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   time.Now(),
		// DeletedAt:   product.DeletedAt,
	}
	err = h.repo.UpdateProduct(updateProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}

func (h Handler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}

	err = h.repo.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewSuccessPayload())
}
