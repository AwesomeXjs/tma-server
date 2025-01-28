package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (p *Portfolio) CreatePortfolio(ctx context.Context, data *model.CreatePortfolioRequest) (int, error) {
	return p.repo.Portfolio.CreatePortfolio(ctx, data)
}
