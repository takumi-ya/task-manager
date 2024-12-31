package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID    int64     `bun:"id,pk,autoincrement"`
	Name  string    `bun:"name,notnull"`
	Done  bool      `bun:"done"`
	Until time.Time `bun:"until"`
}
