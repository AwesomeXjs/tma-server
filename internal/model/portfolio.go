package model

import (
	"database/sql"
	"time"
)

type Portfolio struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	OwnerID   int          `json:"owner_id" db:"owner_id"`
	Profit    int          `json:"profit" db:"profit"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}

type CreatePortfolioSchema struct {
	Name    string `json:"name" db:"name"`
	OwnerID int    `json:"owner_id" db:"owner_id"`
}
