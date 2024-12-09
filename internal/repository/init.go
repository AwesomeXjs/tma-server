package repository

import (
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
)

type Repository struct {
	db dbClient.Client
}

func New(db dbClient.Client) IRepository {
	return &Repository{
		db: db,
	}
}
