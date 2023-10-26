package products

import "github.com/labstack/echo/v4"

func NewRoute(e *echo.Echo) {
	e.GET("/v1/products", GetProduct)
	e.GET("/v1/products/:id", GetProductById)
	e.POST("/v1/products", CreateProduct)
	e.PUT("/v1/products/:id", UpdateProduct)
	e.DELETE("/v1/products/:id", DeleteProduct)
}
