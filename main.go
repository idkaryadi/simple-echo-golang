package main

import (
	"simple-echo-backend/common"
	"simple-echo-backend/config"
	"simple-echo-backend/products"
	"simple-echo-backend/products/handler"
	"simple-echo-backend/products/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	common.NewValidator()
	e := echo.New()

	repo := repository.NewRepository(config.DB)
	handler := handler.NewHandler(&repo)
	products.NewRoute(e, handler)

	e.Logger.Fatal(e.Start(":4000"))
}
