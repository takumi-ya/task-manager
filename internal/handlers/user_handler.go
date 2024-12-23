package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUser(c echo.Context) error {
	users := []User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}

	return c.JSON(http.StatusOK, users)
}
