package portfolio

import (
	"context"
	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (p *Portfolio) CreatePortfolio(ctx context.Context, user *model.Portfolio) error {
	return p.repo.Portfolio.CreatePortfolio(ctx, user)
}
