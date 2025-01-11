package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/internal/models"
	"github.com/uptrace/bun"
)

func GetUser(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		var users []models.User

		err := db.NewSelect().
			Model(&users).
			Scan(ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch users",
			})
		}

		return c.JSON(http.StatusOK, users)
	}
}

func CreateUser(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request payload",
			})
		}

		if user.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Name is required",
			})
		}

		ctx := context.Background()
		_, err := db.NewInsert().
			Model(&user).
			Exec(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create user",
			})
		}

		return c.JSON(http.StatusCreated, user)
	}
}
