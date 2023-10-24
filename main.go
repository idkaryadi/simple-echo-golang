package main

import (
	"net/http"

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
	return c.JSON(http.StatusOK, echo.Map{"status": "success"})
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"id": id}})
}

func CreateProduct(c echo.Context) error {
	return c.JSON(http.StatusCreated, echo.Map{"status": "success"})
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "updateId": id})
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "deletedId": id})
}
