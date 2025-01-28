package portfolio

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (p *Portfolio) CreatePortfolio(ctx context.Context, data *model.CreatePortfolioRequest) (int, error) {
	const mark = "Repository.Portfolio.CreatePortfolio"

	builderInsert := squirrel.Insert(TablePortfolios).
		PlaceholderFormat(squirrel.Dollar).
		Columns(ColumnName, ColumnOwnerID).
		Values(data.Name, data.OwnerID).
		Suffix(SuffixReturnID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return 0, err
	}

	q := dbClient.Query{
		Name:     "CreatePortfolio",
		QueryRaw: query,
	}

	var ID int
	err = p.db.DB().QueryRowContext(ctx, q, args...).Scan(&ID)
	if err != nil {
		logger.Error("failed to create portfolio", mark, zap.Error(err))
		return 0, err
	}

	return ID, nil
}
