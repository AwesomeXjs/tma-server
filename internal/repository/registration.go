package repository

import (
	"context"
	"strings"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) Registration(ctx context.Context, user *model.User) error {
	builderInsert := squirrel.Insert(tableUsers).
		PlaceholderFormat(squirrel.Dollar).
		Columns(columnID, columnUsername, columnFirstName, columnLastName, columnIsPremium).
		Values(user.ID, user.Username, user.FirstName, user.LastName, user.IsPremium).
		Suffix(suffixReturnID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		logger.Error("failed to build insert query", zap.Error(err))
		return err
	}

	q := dbClient.Query{
		Name:     "Registration",
		QueryRaw: query,
	}

	var ID int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			logger.Warn("username already registered", zap.Error(err))
			return err
		}
		logger.Error("failed to register user", zap.Error(err))
		return err
	}
	return nil
}
