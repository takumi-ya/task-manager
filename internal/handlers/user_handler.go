package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/internal/models"
)

func GetUser(c echo.Context) error {
	users := []models.User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}

	return c.JSON(http.StatusOK, users)
}
