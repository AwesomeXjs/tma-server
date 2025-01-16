package repository

import (
	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/internal/repository/portfolio"
	"github.com/AwesomeXjs/tma-server/internal/repository/user"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type Repository struct {
	Portfolio portfolio.IPortfolio
	User      user.IUser
}

func New(db dbClient.Client, cache redis.IRedis) *Repository {
	return &Repository{
		Portfolio: portfolio.New(db, cache),
		User:      user.New(db, cache),
	}
}
