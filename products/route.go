package products

import (
	"simple-echo-backend/products/handler"

	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo, h handler.Handler) {
	e.GET("/v1/products", h.GetProduct)
	e.GET("/v1/products/:id", h.GetProductById)
	e.POST("/v1/products", h.CreateProduct)
	e.PUT("/v1/products/:id", h.UpdateProduct)
	e.DELETE("/v1/products/:id", h.DeleteProduct)
}
