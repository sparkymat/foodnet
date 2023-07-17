package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/foodnet/internal/route"
)

func main() {
	e := echo.New()

	route.Setup(e)

	e.Logger.Fatal(e.Start(":8080"))
}
