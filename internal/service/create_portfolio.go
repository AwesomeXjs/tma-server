package service

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (s *Service) CreatePortfolio(ctx context.Context, user *model.Portfolio) error {
	return s.repo.CreatePortfolio(ctx, user)
}
