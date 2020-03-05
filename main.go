package main

import (
	"echo-mongo/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", controller.Home)
	e.POST("/register", controller.Register)
	e.Logger.Fatal(e.Start(":1234"))
}
