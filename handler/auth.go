package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func generateJWTCookie(id int64) (*http.Cookie, error) {
	// TODO: Move to env file
	secretKey := []byte("my_secret_key")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{Name: "auth", Value: tokenString, HttpOnly: true, Expires: time.Now().Add(24 * time.Hour)}

	return cookie, nil
}

func CookieAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			c.Response().Header().Set("HX-Redirect", "/login")
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})
		if err != nil || !token.Valid {
			c.SetCookie(&http.Cookie{Name: "auth", Value: "", HttpOnly: true, Expires: time.Unix(0, 0)})
			c.Response().Header().Set("HX-Redirect", "/login")
			c.Response().Header().Set("Cache-Control", "no-store")
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		userID, ok := token.Claims.(jwt.MapClaims)["id"].(int64)
		if !ok {
			log.Println("Error getting userID from token")
		} else {
			c.Set("userID", userID)
		}

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
