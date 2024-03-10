package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/layout"
	"github.com/mderler/go_chat/view/login"
	"golang.org/x/exp/rand"
)

var userColors = []string{"#aa2712", "#01b522", "#0176b5", "#b5019a", "#1601b5", "#e29104"}

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
	fullName := c.FormValue("full-name")

	user := UserRequest{
		Username: username,
		Password: password,
		FullName: fullName,
	}

	var passwordConfirmError string

	if password != c.FormValue("password-confirm") {
		passwordConfirmError = "Passwords do not match"
	}

	err := validate.Struct(user)
	if err != nil {
		var usernameError, fullNameError, passwordError string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				usernameError = fmt.Sprintf("%s %s", "Username", getErrorMessage(err))
			case "FullName":
				fullNameError = fmt.Sprintf("%s %s", "Full Name", getErrorMessage(err))
			case "Password":
				passwordError = fmt.Sprintf("%s %s", "Password", getErrorMessage(err))
			}
		}
		return render(c, login.RegisterForm(username, fullName, usernameError, fullNameError, passwordError, passwordConfirmError))
	}
	if passwordConfirmError != "" {
		return render(c, login.RegisterForm(username, fullName, "", "", "", passwordConfirmError))
	}

	randomIndex := rand.Intn(len(userColors))

	createUserParams := model.CreateUserParams{
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Color:    userColors[randomIndex],
	}

	userID, err := h.queries.CreateUser(c.Request().Context(), createUserParams)
	if err != nil {
		log.Println(err)
		return render(c, layout.ErrorBase(internalServerError))
	}

	cookie, err := generateJWTCookie(userID)
	if err != nil {
		log.Println(err)
		return render(c, layout.ErrorBase(internalServerError))
	}

	c.Response().Header().Set("HX-Redirect", "/")
	c.SetCookie(cookie)
	return nil
}

func (h *LoginHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.queries.GetUserForLogin(c.Request().Context(), username)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return render(c, login.LoginForm(username, incorrectPassword))
		}
		return render(c, layout.ErrorBase(internalServerError))
	}

	if user.Password != password {
		return render(c, login.LoginForm(username, incorrectPassword))
	}

	cookie, err := generateJWTCookie(user.ID)
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	c.Response().Header().Set("HX-Redirect", "/")
	c.Response().Header().Set("Cache-Control", "no-store")
	c.SetCookie(cookie)
	return nil
}

func (h *LoginHandler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "auth", Value: "", HttpOnly: true, Expires: time.Unix(0, 0)})
	c.Response().Header().Set("HX-Redirect", "/login")
	c.Response().Header().Set("Cache-Control", "no-store")
	return nil
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "min":
		return "is too short"
	case "max":
		return "is too long"
	default:
		return "is invalid"
	}
}
