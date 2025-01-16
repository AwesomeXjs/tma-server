package portfolio

import (
	"github.com/AwesomeXjs/tma-server/internal/service"
	"github.com/labstack/echo/v4"
)

type IPortfolio interface {
	CreatePortfolio(ctx echo.Context) error
	DeletePortfolio(ctx echo.Context) error
	GetPortfolios(ctx echo.Context) error
	UpdatePortfolio(ctx echo.Context) error
}

type Portfolio struct {
	svc *service.Service
}

func New(svc *service.Service) IPortfolio {
	return &Portfolio{
		svc: svc,
	}
}
