package repository

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

const (
	columnID        = "id"
	columnUsername  = "username"
	columnFirstName = "first_name"
	columnLastName  = "last_name"
	columnIsPremium = "is_premium"

	columnName      = "name"
	columnOwnerID   = "owner_id"
	columnProfit    = "profit"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"

	tableUsers      = "users"
	tablePortfolios = "portfolios"

	suffixReturnID = "RETURNING id"
)

type IRepository interface {
	Registration(ctx context.Context, user *model.User) error
	CreatePortfolio(ctx context.Context, user *model.Portfolio) error
}
