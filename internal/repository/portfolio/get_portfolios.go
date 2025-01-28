package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (p *Portfolio) GetPortfolios(ctx context.Context, ownerId int) (model.GetPortfoliosResponse, error) {

	const mark = "Repository.Portfolio.GetPortfolios"

	builderSelect := squirrel.Select(
		ColumnID,
		ColumnName,
		ColumnProfit,
		ColumnCreatedAt,
		ColumnUpdatedAt,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(TablePortfolios).
		Where(squirrel.Eq{ColumnOwnerID: ownerId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return nil, err
	}

	q := dbClient.Query{
		Name:     "GetPortfolios",
		QueryRaw: query,
	}

	rows, err := p.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to get portfolios", mark, zap.Error(err))
		return nil, err
	}

	var portfolios model.GetPortfoliosResponse

	for rows.Next() {
		var portfolio model.PortfolioResponse
		err = rows.Scan(
			&portfolio.ID,
			&portfolio.Name,
			&portfolio.Profit,
			&portfolio.CreatedAt,
			&portfolio.UpdatedAt,
		)
		if err != nil {
			logger.Error("failed to scan portfolio", mark, zap.Error(err))
			return nil, err
		}

		portfolios = append(portfolios, portfolio)
	}

	return portfolios, nil
}
