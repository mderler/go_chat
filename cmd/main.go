package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gochat "github.com/mderler/go_chat"
	"github.com/mderler/go_chat/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/style.css", func(c echo.Context) error {
		return c.String(200, gochat.Styles)
	})

	indexHandler := handler.IndexHandler{}
	e.GET("/", indexHandler.ShowIndex)

	userHandler := handler.UserHandler{}
	e.GET("/user", userHandler.ShowUser)

	e.Logger.Fatal(e.Start(":8000"))
}
