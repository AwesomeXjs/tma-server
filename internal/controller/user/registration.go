package user

import (
	"net/http"
	"strings"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils"
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
// @Success 200 {object} utils.Body
// @Failure 400 {object} utils.Body
// @Router /api/v1/registration [post]
func (u *User) Registration(ctx echo.Context) error {
	const mark = "Controller.User.Registration"

	var User model.User
	err := ctx.Bind(&User)
	if err != nil {
		logger.Error("failed to bind user", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	err = u.svc.User.Registration(ctx.Request().Context(), &User)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			logger.Info("user already registered", mark, zap.Error(err))
			return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
		}
		logger.Error("failed to register user", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	return utils.Response(ctx, http.StatusOK, "success", "user registered")
}
