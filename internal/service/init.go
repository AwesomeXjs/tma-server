package service

import (
	"github.com/AwesomeXjs/tma-server/internal/repository"
)

type Service struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) IService {
	return &Service{
		repo: repo,
	}
}
