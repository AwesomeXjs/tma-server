package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/repository"
)

type IPortfolio interface {
	CreatePortfolio(ctx context.Context, data *model.CreatePortfolioRequest) (int, error)
	GetPortfolios(ctx context.Context, ownerId int) (model.GetPortfoliosResponse, error)
	UpdatePortfolio(ctx context.Context, data model.UpdatePortfolioRequest) error
	DeletePortfolio(ctx context.Context, portfolio model.DeletePortfolioRequest) error
}

type Portfolio struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) IPortfolio {
	return &Portfolio{
		repo: repo,
	}
}
