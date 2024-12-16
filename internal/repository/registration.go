package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) Registration(ctx context.Context, user *model.User) error {
	_, err := r.cache.Get(ctx, strconv.Itoa(user.ID))
	if err == nil {
		logger.Warn("user already registered (info from cache)", zap.String("username", user.Username))
		return fmt.Errorf("duplicate key value violates unique constraint")
	}

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
		logger.Error("failed to register user", zap.Error(err))
		return err
	}

	// cache retries (spam /start protection)
	err = r.cache.Set(ctx, strconv.Itoa(user.ID), user.Username, time.Hour)
	if err != nil {
		return err
	}

	return nil
}
