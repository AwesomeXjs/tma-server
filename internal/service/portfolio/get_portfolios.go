package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"go.uber.org/zap"
)

func (p *Portfolio) GetPortfolios(ctx context.Context, ownerId int) (model.GetPortfoliosResponse, error) {

	const mark = "Service.Portfolio.GetPortfolios"

	portfolios, err := p.repo.Portfolio.GetPortfolios(ctx, ownerId)
	if err != nil {
		logger.Error("failed to get portfolios", mark, zap.Error(err))
		return nil, err
	}

	return portfolios, nil
}
