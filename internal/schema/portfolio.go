package schema

import "time"

type Portfolio struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Portfolio"`
	Profit    float64   `json:"profit" example:"1000.0"`
	CreatedAt time.Time `json:"created_at" example:"2023-08-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-08-01T00:00:00Z"`
}

type GetPortfolios struct {
	Base
	Data []Portfolio
}

type CreatePortfolio struct {
	Base
	Data int `json:"data" example:"1"`
}
