package controller

import (
	"github.com/AwesomeXjs/tma-server/internal/controller/assets"
	"github.com/AwesomeXjs/tma-server/internal/controller/portfolio"
	"github.com/AwesomeXjs/tma-server/internal/controller/user"
	"github.com/AwesomeXjs/tma-server/internal/service"
)

// Controller handles the authentication and header-related operations.
type Controller struct {
	User      user.IUser
	Portfolio portfolio.IPortfolio
	Assets    assets.IAssets
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New(svc *service.Service) *Controller {
	return &Controller{
		User:      user.New(svc),
		Portfolio: portfolio.New(svc),
		Assets:    assets.New(svc),
	}
}
