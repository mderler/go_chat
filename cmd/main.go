package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mderler/go_chat/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	indexHandler := handler.IndexHandler{}
	e.GET("/", indexHandler.ShowIndex)

	userHandler := handler.UserHandler{}
	e.GET("/user", userHandler.ShowUser)

	e.Logger.Fatal(e.Start(":8000"))
}
