package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/handler"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userHandler := handler.UserHandler{}
	e.GET("/user", userHandler.ShowUser)

	e.Logger.Fatal(e.Start(":8000"))
}
