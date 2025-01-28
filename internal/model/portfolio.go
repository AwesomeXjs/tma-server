package model

import (
	"database/sql"
	"time"
)

type Portfolio struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	OwnerID   int          `json:"owner_id" db:"owner_id"`
	Profit    float64      `json:"profit" db:"profit"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}

type CreatePortfolioRequest struct {
	Name    string `json:"name" db:"name" example:"Portfolio"`
	OwnerID int    `json:"owner_id" db:"owner_id" example:"518774723"`
}

type DeletePortfolioRequest struct {
	ID      int `json:"id" db:"id" example:"1"`
	OwnerID int `json:"ownerID" db:"owner_id"  example:"518774723"`
}

type UpdatePortfolioRequest struct {
	ID      int    `json:"id" db:"id" example:"1"`
	OwnerID int    `json:"owner_id" db:"owner_id" example:"518774723"`
	NewName string `json:"new_name" db:"name" example:"Portfolio"`
}

type PortfolioResponse struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Profit    float64   `json:"profit" db:"profit"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type GetPortfoliosResponse []PortfolioResponse
