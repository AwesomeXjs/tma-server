package controller

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils/response"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Registration - Registration
// @Summary Registration
// @Tags User
// @Description Registration new user
// @ID registration
// @Accept  json
// @Produce  json
// @Param input body model.User true "user info"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Router /api/v1/registration [post]
func (c *Controller) Registration(ctx echo.Context) error {
	const mark = "Controller.Registration"

	var User model.User
	err := ctx.Bind(&User)
	if err != nil {
		logger.Error("failed to bind user", mark, zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}
	
	err = c.svc.Registration(ctx.Request().Context(), &User)
	if err != nil {
		logger.Error("failed to register user", mark, zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	return response.Response(ctx, http.StatusOK, "success", "user registered")
}
