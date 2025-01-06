package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/takumi-ya/taskmanager/configs"
	"github.com/takumi-ya/taskmanager/internal/handlers"
	"github.com/takumi-ya/taskmanager/internal/models"
	"github.com/uptrace/bun"
)

func main() {
	// 環境変数をロード
	configs.LoadEnv()

	// Echoのインスタンス作成
	e := echo.New()

	// DB接続
	conn := configs.NewDB()
	defer conn.CloseDB()

	if err := migrate(conn.DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", handlers.GetUser(conn.DB))
	e.POST("/users", handlers.CreateUser(conn.DB))

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8082"
	}

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + appPort))
}

func migrate(db *bun.DB) error {
	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*models.User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatalf("Failed to create table of users: %v", err)
	}

	_, err = db.NewCreateTable().Model((*models.Task)(nil)).
		ForeignKey(`("user_id") REFERENCES "users" ("id") ON DELETE CASCADE`).
		IfNotExists().
		Exec(ctx)

	return err
}
