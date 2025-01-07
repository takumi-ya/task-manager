package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/internal/models"
	"github.com/uptrace/bun"
)

func GetTask(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		var tasks []models.Task

		err := db.NewSelect().
			Model(&tasks).
			Scan(ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch tasks",
			})
		}

		return c.JSON(http.StatusOK, tasks)
	}
}
