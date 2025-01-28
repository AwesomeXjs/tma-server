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
// @Param input body model.CreatePortfolioRequest true "portfolio info"
// @Success 200 {object} schema.CreatePortfolio
// @Failure 400 {object} schema.NoData
// @Router /api/v1/create-portfolio [post]
func (p *Portfolio) CreatePortfolio(ctx echo.Context) error {
	const mark = "Controller.Portfolio.CreatePortfolio"

	var Request model.CreatePortfolioRequest

	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to create portfolio", err.Error())
	}
	ID, err := p.svc.Portfolio.CreatePortfolio(ctx.Request().Context(), &Request)
	if err != nil {
		logger.Error("failed to create portfolio", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to create portfolio", err.Error())
	}

	return utils.Response(ctx, http.StatusOK, utils.SuccessMessage, ID)
}
