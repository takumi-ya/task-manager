package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DBConnection struct {
	DB    *bun.DB
	SQLDB *sql.DB
}

func NewDB() *DBConnection {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOSTNAME")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	// 接続確認
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	log.Println("Database connected successfully")

	return &DBConnection{
		DB:    db,
		SQLDB: sqldb,
	}
}

func (conn *DBConnection) CloseDB() {
	if err := conn.SQLDB.Close(); err != nil {
		log.Fatalf("Error closing SQLDB: %v", err)
	}

	if err := conn.DB.Close(); err != nil {
		log.Fatalf("Error closing DB: %v", err)
	}

	log.Println("Database connection closed")
}
