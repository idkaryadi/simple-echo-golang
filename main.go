package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/products", GetProduct)
	e.GET("/products/:id", GetProductById)

	e.Logger.Fatal(e.Start(":4000"))
}

func GetProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": "success"})
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"id": id}})
}
