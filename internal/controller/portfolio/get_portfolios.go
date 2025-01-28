package portfolio

import (
	"net/http"
	"strconv"

	"github.com/AwesomeXjs/tma-server/internal/utils"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// GetPortfolios - GetPortfolios
// @Summary GetPortfolios
// @Security TGWebAppToken
// @Tags Portfolio
// @Description get portfolios
// @ID get-portfolio
// @Param id path int false "owner id"
// @Accept  json
// @Produce  json
// @Success 200 {object} schema.GetPortfolios
// @Failure 400 {object} schema.NoData
// @Router /api/v1/portfolios/{id} [get]
func (p *Portfolio) GetPortfolios(ctx echo.Context) error {

	const mark = "Controller.Portfolio.GetPortfolios"

	id := ctx.Param("id")

	if len(id) == 0 {
		return utils.Response(ctx, http.StatusBadRequest, "bad request", "invalid id")
	}

	ownerID, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("failed to convert id to int", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", "invalid id")
	}

	portfolios, err := p.svc.Portfolio.GetPortfolios(ctx.Request().Context(), ownerID)
	if err != nil {
		logger.Error("failed to get portfolios", mark, zap.Error(err))
		return utils.Response(ctx, http.StatusBadRequest, "bad request", err.Error())
	}
	return utils.Response(ctx, http.StatusOK, utils.SuccessMessage, portfolios)
}
