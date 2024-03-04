package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gochat "github.com/mderler/go_chat"
	"github.com/mderler/go_chat/handler"
	"github.com/mderler/go_chat/model"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(gochat.EmbedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	queries := model.New(db)

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/style.css", func(c echo.Context) error {
		return c.String(200, gochat.Styles)
	})

	e.GET("/htmx.min.js", func(c echo.Context) error {
		return c.String(200, gochat.HTMX)
	})

	indexHandler := handler.IndexHandler{}
	e.GET("/", indexHandler.ShowIndex)

	userHandler := handler.UserHandler{}
	e.GET("/user", userHandler.ShowUser)

	loginHandler := handler.NewLoginHandler(queries)
	e.GET("/login", loginHandler.ShowLogin)
	e.POST("/login", loginHandler.Login)
	e.GET("/register", loginHandler.ShowRegister)
	e.POST("/register", loginHandler.Register)

	e.Logger.Fatal(e.Start(":8000"))
}
