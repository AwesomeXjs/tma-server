package controller

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils/response"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (c *Controller) Registration(ctx echo.Context) error {
	var User model.User
	err := ctx.Bind(&User)
	if err != nil {
		logger.Error("failed to bind user", zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	err = c.svc.Registration(ctx.Request().Context(), &User)
	if err != nil {
		logger.Error("failed to register user", zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	return response.Response(ctx, http.StatusOK, "success", "user registered")
}
