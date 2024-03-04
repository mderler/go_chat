package handler

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func generateJWTToken(sub int64) (string, error) {
	// TODO: Move to env file
	secretKey := []byte("my_secret_key")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["sub"] = sub

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func cookieAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("username")
		if err != nil {
			return c.Redirect(302, "/login")
		}
		return next(c)
	}
}
