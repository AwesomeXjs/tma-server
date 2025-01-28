package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type IPortfolio interface {
	CreatePortfolio(ctx context.Context, data *model.CreatePortfolioRequest) (int, error)
	DeletePortfolio(ctx context.Context, portfolio model.DeletePortfolioRequest) error
	GetPortfolios(ctx context.Context, ownerId int) (model.GetPortfoliosResponse, error)
	UpdatePortfolio(ctx context.Context, data model.UpdatePortfolioRequest) error
}

type Portfolio struct {
	db    dbClient.Client
	cache redis.IRedis
}

func New(db dbClient.Client, cache redis.IRedis) IPortfolio {
	return &Portfolio{
		db:    db,
		cache: cache,
	}
}
