package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID       int64     `bun:"id,pk,autoincrement"`
	Name     string    `bun:"name,notnull"`
	Done     bool      `bun:"done"`
	Until    time.Time `bun:"until"`
	CreateAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UserID   int64     `bun:"user_id"`
	User     *User     `bun:"rel:belongs-to,join:user_id=id"`
}
