package main

import (
	"simple-echo-backend/common"
	"simple-echo-backend/config"
	"simple-echo-backend/products"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	common.NewValidator()
	e := echo.New()
	products.NewRoute(e)

	e.Logger.Fatal(e.Start(":4000"))
}
