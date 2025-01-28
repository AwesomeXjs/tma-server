package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
)

func (p *Portfolio) UpdatePortfolio(ctx context.Context, data model.UpdatePortfolioSchema) error {

	const mark = "Service.Portfolio.UpdatePortfolio"

	return p.repo.Portfolio.UpdatePortfolio(ctx, data)
}
