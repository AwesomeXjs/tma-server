package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (p *Portfolio) DeletePortfolio(ctx context.Context, portfolio model.DeletePortfolioSchema) error {
	const mark = "Service.Portfolio.DeletePortfolio"

	return p.repo.Portfolio.DeletePortfolio(ctx, portfolio)
}
