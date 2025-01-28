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
// @Param input body model.DeletePortfolioSchema true "portfolio info"
// @Success 200 {object} utils.Body
// @Failure 400 {object} utils.Body
// @Router /api/v1/delete-portfolio [delete]
func (p *Portfolio) DeletePortfolio(ctx echo.Context) error {

	const mark = "Controller.Portfolio.DeletePortfolio"

	var Request model.DeletePortfolioSchema
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

	return utils.Response(ctx, http.StatusOK, "success", "portfolio deleted")
}
