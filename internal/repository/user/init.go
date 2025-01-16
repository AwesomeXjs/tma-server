package user

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type IUser interface {
	Registration(ctx context.Context, user *model.User) error
}

type User struct {
	db    dbClient.Client
	cache redis.IRedis
}

func New(db dbClient.Client, cache redis.IRedis) IUser {
	return &User{
		db:    db,
		cache: cache,
	}
}
