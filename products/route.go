package products

import "github.com/labstack/echo/v4"

func NewRoute(e *echo.Echo) {
	e.GET("/v1/products", GetProductController)
	e.GET("/v1/products/:id", GetProductByIdController)
	e.POST("/v1/products", CreateProductController)
	e.PUT("/v1/products/:id", UpdateProductController)
	e.DELETE("/v1/products/:id", DeleteProductController)
}
