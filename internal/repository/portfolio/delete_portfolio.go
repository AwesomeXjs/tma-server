package portfolio

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (p *Portfolio) DeletePortfolio(ctx context.Context, portfolio model.DeletePortfolioRequest) error {

	const mark = "Repository.Portfolio.DeletePortfolio"

	builderDelete := squirrel.Delete(TablePortfolios).PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{ColumnID: portfolio.ID}).Where(squirrel.Eq{ColumnOwnerID: portfolio.OwnerID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return err
	}

	q := dbClient.Query{
		Name:     "DeletePortfolio",
		QueryRaw: query,
	}

	result, err := p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to create portfolio", mark, zap.Error(err))
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("portfolio not found")
	}

	return nil
}
