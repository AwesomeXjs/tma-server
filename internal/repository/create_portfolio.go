package repository

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) CreatePortfolio(ctx context.Context, user *model.Portfolio) error {
	const mark = "Repository.CreatePortfolio"

	builderInsert := squirrel.Insert(tablePortfolios).
		PlaceholderFormat(squirrel.Dollar).
		Columns(columnName, columnOwnerID).
		Values(user.Name, user.OwnerID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", mark, zap.Error(err))
		return err
	}

	q := dbClient.Query{
		Name:     "CreatePortfolio",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("failed to create portfolio", mark, zap.Error(err))
		return err
	}

	return nil
}
