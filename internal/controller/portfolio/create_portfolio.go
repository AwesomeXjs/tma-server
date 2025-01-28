package portfolio

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/utils"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// CreatePortfolio - CreatePortfolio
// @Summary CreatePortfolio
// @Security TGWebAppToken
// @Tags Portfolio
// @Description create portfolio for user
// @ID create-portfolio
// @Accept  json
// @Produce  json
// @Param input body model.CreatePortfolioSchema true "portfolio info"
// @Success 200 {object} utils.Body
// @Failure 400 {object} utils.Body
// @Router /api/v1/create-portfolio [get]
func (p *Portfolio) CreatePortfolio(ctx echo.Context) error {
	const mark = "Controller.Portfolio.CreatePortfolio"

	var Request model.Portfolio

	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request 123", err.Error())
	}
	err = p.svc.Portfolio.CreatePortfolio(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to create portfolio", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request 321", err.Error())
	}

	return utils.Response(ctx, http.StatusOK, "success", "portfolio created")
}
