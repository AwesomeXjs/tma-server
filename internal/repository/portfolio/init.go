package portfolio

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type IPortfolio interface {
	CreatePortfolio(ctx context.Context, user *model.Portfolio) error
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
