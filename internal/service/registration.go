package service

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (s *Service) Registration(ctx context.Context, user *model.User) error {
	const mark = "Service.Registration"

	return s.repo.Registration(ctx, user)
}
