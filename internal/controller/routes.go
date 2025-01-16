package controller

import (
	_ "github.com/AwesomeXjs/tma-server/docs"
	"github.com/AwesomeXjs/tma-server/internal/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

const (
	baseBath = "/api"

	version1        = "/v1"
	registration    = "/registration"
	createPortfolio = "/create-portfolio"
	deletePortfolio = "/delete-portfolio"
	updatePortfolio = "/update-portfolio"
	getPortfolios   = "/portfolios"

	addAsset = "/add-asset"

	test = "/test-validate"
)

// InitRoutes initializes all the routes for the application.
func (c *Controller) InitRoutes(server *echo.Echo, botToken string) {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group(baseBath)
	{
		v1 := api.Group(version1)
		{
			v1.POST(registration, c.User.Registration)

			secureRoutes := v1.Group("")
			{
				secureRoutes.POST(createPortfolio, c.Portfolio.CreatePortfolio)
				secureRoutes.DELETE(deletePortfolio, c.Portfolio.DeletePortfolio)
				secureRoutes.PATCH(updatePortfolio, c.Portfolio.UpdatePortfolio)
				secureRoutes.GET(getPortfolios, c.Portfolio.GetPortfolios)

				secureRoutes.POST(addAsset, c.Assets.AddAssetToPortfolio)

				secureRoutes.POST(test, func(c echo.Context) error {
					return c.JSON(http.StatusOK, map[string]string{"message": "success"})
				}, middlewares.TelegramValidationMiddleware(botToken))
			}

		}
	}
}
