package service

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

type IService interface {
	Registration(ctx context.Context, user *model.User) error
	CreatePortfolio(ctx context.Context, user *model.Portfolio) error
}
