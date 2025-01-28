package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/repository"
)

type IPortfolio interface {
	CreatePortfolio(ctx context.Context, user *model.Portfolio) error
	GetPortfolios(ctx context.Context, ownerId int) (model.GetPortfoliosResponse, error)
	UpdatePortfolio(ctx context.Context, data model.UpdatePortfolioSchema) error
	DeletePortfolio(ctx context.Context, portfolio model.DeletePortfolioSchema) error
}

type Portfolio struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) IPortfolio {
	return &Portfolio{
		repo: repo,
	}
}
