package products

import (
	"simple-echo-backend/products/handler"

	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo) {
	e.GET("/v1/products", handler.GetProduct)
	e.GET("/v1/products/:id", handler.GetProductById)
	e.POST("/v1/products", handler.CreateProduct)
	e.PUT("/v1/products/:id", handler.UpdateProduct)
	e.DELETE("/v1/products/:id", handler.DeleteProduct)
}
