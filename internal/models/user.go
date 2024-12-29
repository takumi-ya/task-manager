package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
}
