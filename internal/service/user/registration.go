package user

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (u *User) Registration(ctx context.Context, user *model.User) error {
	const mark = "Service.User.Registration"

	return u.repo.User.Registration(ctx, user)
}
