package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/configs"
	"github.com/takumi-ya/taskmanager/internal/handlers"
)

func main() {
	// 環境変数をロード
	configs.LoadEnv()

	// Echoのインスタンス作成
	e := echo.New()

	// DB接続
	db := configs.NewDB()
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", handlers.GetUser)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + appPort))
}
