package service

import (
	"github.com/AwesomeXjs/tma-server/internal/repository"
	"github.com/AwesomeXjs/tma-server/internal/service/portfolio"
	"github.com/AwesomeXjs/tma-server/internal/service/user"
)

type Service struct {
	User      user.IUser
	Portfolio portfolio.IPortfolio
}

func New(repo *repository.Repository) *Service {
	return &Service{
		User:      user.New(repo),
		Portfolio: portfolio.New(repo),
	}
}
