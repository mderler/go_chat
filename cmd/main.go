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

	assetsFs := echo.MustSubFS(gochat.Public, "public")

	e := echo.New()
	e.Use(middleware.Recover())

	e.StaticFS("/public", assetsFs)

	indexHandler := handler.NewIndexHandler(queries)

	authGroup := e.Group("")
	authGroup.Use(handler.CookieAuth)
	authGroup.GET("/", indexHandler.ShowIndex)

	userHandler := handler.NewUserHandler(queries)
	e.GET("/user", userHandler.ShowUserList)

	loginHandler := handler.NewLoginHandler(queries)

	redirectGroup := e.Group("")
	redirectGroup.Use(handler.RedirectIfAuthenticated)
	redirectGroup.GET("/login", loginHandler.ShowLogin)
	redirectGroup.GET("/register", loginHandler.ShowRegister)
	e.POST("/login", loginHandler.Login)
	e.POST("/logout", loginHandler.Logout)
	e.POST("/register", loginHandler.Register)

	e.Logger.Fatal(e.Start(":8000"))
}
