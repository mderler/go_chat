package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gochat "github.com/mderler/go_chat"
	"github.com/mderler/go_chat/chat"
	"github.com/mderler/go_chat/handler"
	"github.com/mderler/go_chat/model"
	"github.com/pressly/goose/v3"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

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

	hub := chat.NewHub()
	go hub.Run()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logFile}))

	assetsFs := echo.MustSubFS(gochat.Public, "public")
	e.StaticFS("/public", assetsFs)

	redirectGroup := e.Group("")

	loginHandler := handler.NewLoginHandler(queries)
	redirectGroup.Use(handler.RedirectIfAuthenticated)
	redirectGroup.GET("/login", loginHandler.ShowLogin)
	redirectGroup.GET("/register", loginHandler.ShowRegister)
	e.POST("/login", loginHandler.Login)
	e.POST("/logout", loginHandler.Logout)
	e.POST("/register", loginHandler.Register)

	authGroup := e.Group("")
	authGroup.Use(handler.CookieAuth)

	userHandler := handler.NewUserHandler(queries)
	authGroup.GET("/user", userHandler.ShowUserList)

	indexHandler := handler.NewIndexHandler(queries)

	authGroup.GET("/", indexHandler.ShowIndex)
	authGroup.GET("/chat", indexHandler.ShowChat)
	authGroup.GET("/messages", indexHandler.ShowMessages)

	authGroup.GET("/ws", func(c echo.Context) error {
		userID := c.Get("userID").(int64)
		chat.NewClient(hub, userID, queries, db, c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
