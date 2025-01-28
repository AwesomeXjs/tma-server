package portfolio

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (p *Portfolio) UpdatePortfolio(ctx context.Context, data model.UpdatePortfolioSchema) error {

	const mark = "Repository.Portfolio.UpdatePortfolio"

	builderUpdate := squirrel.Update(TablePortfolios).PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{ColumnID: data.ID}).
		Where(squirrel.Eq{ColumnOwnerID: data.OwnerID}).
		Set(ColumnName, data.NewName).Set(ColumnUpdatedAt, time.Now())

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return err
	}

	q := dbClient.Query{
		Name:     "UpdatePortfolio",
		QueryRaw: query,
	}

	result, err := p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to update portfolio", mark, zap.Error(err))
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("portfolio not found")
	}

	return nil
}
