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
				"error": "failed to fetch users",
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
				"error": "invalid request payload",
			})
		}

		if user.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "name is required",
			})
		}

		ctx := context.Background()
		_, err := db.NewInsert().
			Model(&user).
			Exec(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to create user",
			})
		}

		return c.JSON(http.StatusCreated, user)
	}
}

func DeleteUser(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := ParseID(c, "user")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		var user models.User
		ctx := context.Background()

		if _, err := db.NewDelete().Model(&user).Where("id = ?", id).Exec(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to delete user",
			})
		}

		return c.JSON(http.StatusNoContent, map[string]string{})
	}
}
