package user

import (
	"github.com/AwesomeXjs/tma-server/internal/service"
	"github.com/labstack/echo/v4"
)

type IUser interface {
	Registration(ctx echo.Context) error
}

type User struct {
	svc *service.Service
}

func New(svc *service.Service) IUser {
	return &User{
		svc: svc,
	}
}
