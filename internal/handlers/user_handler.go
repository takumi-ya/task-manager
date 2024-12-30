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

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

func CreateUser(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request payload",
			})
		}

		if req.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Name is required",
			})
		}

		user := &models.User{
			Name: req.Name,
		}

		ctx := context.Background()
		_, err := db.NewInsert().
			Model(user).
			Exec(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create user",
			})
		}

		return c.JSON(http.StatusCreated, user)
	}
}
