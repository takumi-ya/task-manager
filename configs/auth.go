package configs

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAuth(e *echo.Echo) {
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().RequestURI == "/users" && c.Request().Method == "POST" {
				return false
			}
			return true
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			// Be careful to use constant time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(username), []byte("Joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		},
	}))
}
