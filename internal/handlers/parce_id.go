package handlers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseID(c echo.Context, apiName string) (int64, error) {
	taskID := c.Param("id")
	if taskID == "" {
		return 0, fmt.Errorf(apiName + " ID is required")
	}

	taskIDInt, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid " + apiName + " ID")
	}

	return taskIDInt, nil
}
