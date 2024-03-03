package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/login"
)

type LoginHandler struct {
	queries *model.Queries
}

func NewLoginHandler(queries *model.Queries) *LoginHandler {
	return &LoginHandler{queries: queries}
}

func (h *LoginHandler) ShowLogin(c echo.Context) error {
	return render(c, login.ShowLogin())
}

func (h *LoginHandler) ShowRegister(c echo.Context) error {
	return render(c, login.ShowRegister())
}

func (h *LoginHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := UserRequest{
		Username: username,
		Password: password,
	}

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			panic(err)
		}
	}

	h.queries.CreateUser(c.Request().Context(), model.CreateUserParams(user))

	return render(c, login.RegisterForm("too long", "test", "not matching"))
}
