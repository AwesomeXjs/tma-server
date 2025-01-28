package portfolio

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UpdatePortfolio - UpdatePortfolio
// @Summary UpdatePortfolio
// @Security TGWebAppToken
// @Tags Portfolio
// @Description update portfolio
// @ID update-portfolio
// @Accept  json
// @Produce  json
// @Param input body model.UpdatePortfolioRequest true "portfolio info"
// @Success 200 {object} schema.NoData
// @Failure 400 {object} schema.NoData
// @Router /api/v1/update-portfolio [patch]
func (p *Portfolio) UpdatePortfolio(ctx echo.Context) error {

	const mark = "Controller.Portfolio.UpdatePortfolio"

	var Request model.UpdatePortfolioRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	err = p.svc.Portfolio.UpdatePortfolio(ctx.Request().Context(), Request)
	if err != nil {
		logger.Error("failed to update portfolio", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "failed to update", err.Error())
	}

	return utils.Response(ctx, http.StatusOK, utils.SuccessMessage, "portfolio updated")
}
