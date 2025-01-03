package controller

import (
	_ "github.com/AwesomeXjs/tma-server/docs"
	"github.com/AwesomeXjs/tma-server/internal/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// InitRoutes initializes all the routes for the application.
func (c *Controller) InitRoutes(server *echo.Echo, botToken string) {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/registration", c.Registration)
			v1.POST("/create-portfolio", c.CreatePortfolio)
			v1.POST("/test-validate", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{"message": "success"})
			}, middlewares.TelegramValidationMiddleware(botToken))
			v1.OPTIONS("/*", func(c echo.Context) error {
				return c.NoContent(http.StatusOK)
			})
		}
	}
}
