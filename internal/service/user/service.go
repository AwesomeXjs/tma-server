package user

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/repository"
)

type IUser interface {
	Registration(ctx context.Context, user *model.User) error
}

type User struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) IUser {
	return &User{
		repo: repo,
	}
}
