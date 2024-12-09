package service

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (s *Service) Registration(ctx context.Context, user *model.User) error {
	return s.repo.Registration(ctx, user)
}
