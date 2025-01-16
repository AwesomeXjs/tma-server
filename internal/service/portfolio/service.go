package portfolio

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/repository"
)

type IPortfolio interface {
	CreatePortfolio(ctx context.Context, user *model.Portfolio) error
	GetPortfolios(ctx context.Context) error
	UpdatePortfolio(ctx context.Context) error
	DeletePortfolio(ctx context.Context) error
}

type Portfolio struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) IPortfolio {
	return &Portfolio{
		repo: repo,
	}
}
