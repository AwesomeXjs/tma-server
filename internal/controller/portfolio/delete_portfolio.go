package portfolio

import (
	"net/http"

	"github.com/AwesomeXjs/tma-server/internal/model"
	"github.com/AwesomeXjs/tma-server/internal/utils"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// DeletePortfolio - DeletePortfolio
// @Summary DeletePortfolio
// @Security TGWebAppToken
// @Tags Portfolio
// @Description delete portfolio
// @ID delete-portfolio
// @Accept  json
// @Produce  json
// @Param input body model.DeletePortfolioRequest true "portfolio info"
// @Success 200 {object} schema.NoData
// @Failure 400 {object} schema.NoData
// @Router /api/v1/delete-portfolio [delete]
func (p *Portfolio) DeletePortfolio(ctx echo.Context) error {

	const mark = "Controller.Portfolio.DeletePortfolio"

	var Request model.DeletePortfolioRequest
	err := ctx.Bind(&Request)
	if err != nil {
		logger.Error("failed to bind request", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	err = p.svc.Portfolio.DeletePortfolio(ctx.Request().Context(), Request)
	if err != nil {
		logger.Error("failed to delete portfolio", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}

	return utils.Response(ctx, http.StatusOK, utils.SuccessMessage, "portfolio deleted")
}
