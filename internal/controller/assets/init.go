package assets

import (
	"github.com/AwesomeXjs/tma-server/internal/service"
	"github.com/labstack/echo/v4"
)

type IAssets interface {
	AddAssetToPortfolio(ctx echo.Context) error
}

type Assets struct {
	svc *service.Service
}

func New(svc *service.Service) IAssets {
	return &Assets{
		svc: svc,
	}
}
