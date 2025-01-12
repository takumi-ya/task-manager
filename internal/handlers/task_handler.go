package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/internal/models"
	"github.com/uptrace/bun"
)

func GetTasks(db *bun.DB) echo.HandlerFunc {
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

func GetTask(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "ID must be integer",
			})
		}

		ctx := context.Background()

		var task models.Task

		err = db.NewSelect().
			Model(&task).
			Where("id = ?", idInt).
			Scan(ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update task id: " + id,
			})
		}

		return c.JSON(http.StatusOK, task)
	}
}

type createTaskRequest struct {
	Name   string `json:"name" validate:"required"`
	Until  string `json:"until" validate:"required"`
	UserID int64  `json:"user_id" validate:"required"`
}

func CreateTask(db *bun.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req createTaskRequest
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
		untilTime, err := time.Parse("2006/1/2 15:04:05", req.Until)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Until is required",
			})
		}
		if req.UserID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "User ID is required",
			})
		}

		task := &models.Task{
			Name:   req.Name,
			Until:  untilTime,
			UserID: req.UserID,
		}

		ctx := context.Background()

		if _, err := db.NewInsert().Model(task).Exec(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create task",
			})
		}

		return c.JSON(http.StatusCreated, task)
	}
}
