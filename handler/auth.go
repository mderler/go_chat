package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/view/layout"
)

func generateJWTCookie(id int64) (*http.Cookie, error) {
	// TODO: Move to env file
	secretKey := []byte("my_secret_key")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cookie := &http.Cookie{Name: "auth", Value: tokenString, HttpOnly: true, Expires: time.Now().Add(24 * time.Hour)}

	return cookie, nil
}

func CookieAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			fmt.Println(err)
			c.Response().Header().Set("HX-Redirect", "/login")
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})
		if err != nil || !token.Valid {
			fmt.Println(err)
			c.SetCookie(&http.Cookie{Name: "auth", Value: "", HttpOnly: true, Expires: time.Unix(0, 0)})
			c.Response().Header().Set("HX-Redirect", "/login")
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		userID, ok := token.Claims.(jwt.MapClaims)["id"].(float64)
		if !ok {
			log.Println("Error getting userID from token")
			return render(c, layout.ErrorBase(internalServerError))
		}

		c.Set("userID", int64(userID))

		return next(c)
	}
}

func RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("auth")
		if err == nil {
			c.Response().Header().Set("HX-Redirect", "/")
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.Redirect(http.StatusMovedPermanently, "/")
		}

		return next(c)
	}
}
