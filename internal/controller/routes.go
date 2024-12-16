package controller

import (
	_ "github.com/AwesomeXjs/tma-server/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// InitRoutes initializes all the routes for the application.
func (c *Controller) InitRoutes(server *echo.Echo) {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	api := server.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/registration", c.Registration)
			v1.POST("/create-portfolio", c.CreatePortfolio)
		}
	}
}
