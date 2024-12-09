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

	tableUsers = "users"

	suffixReturnID = "RETURNING id"
)

type IRepository interface {
	Registration(ctx context.Context, user *model.User) error
}
