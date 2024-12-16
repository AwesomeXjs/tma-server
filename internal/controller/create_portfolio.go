package controller

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils/response"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// CreatePortfolio - CreatePortfolio
// @Summary CreatePortfolio
// @Tags Portfolio
// @Description create portfolio for user
// @ID create-portfolio
// @Accept  json
// @Produce  json
// @Param input body model.CreatePortfolioSchema true "portfolio info"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.Body
// @Router /api/v1/create-portfolio [post]
func (c *Controller) CreatePortfolio(ctx echo.Context) error {
	var Request model.Portfolio

	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request 123", err.Error())
	}
	err = c.svc.CreatePortfolio(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to create portfolio", zap.Error(err))
		return response.Response(ctx, http.StatusBadRequest, "bad request 321", err.Error())
	}

	return response.Response(ctx, http.StatusOK, "success", "portfolio created")
}
