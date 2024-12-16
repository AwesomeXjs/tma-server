package repository

import (
	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type Repository struct {
	db    dbClient.Client
	cache redis.IRedis
}

func New(db dbClient.Client, cache redis.IRedis) IRepository {
	return &Repository{
		db:    db,
		cache: cache,
	}
}
