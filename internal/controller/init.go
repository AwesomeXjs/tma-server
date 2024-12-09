package controller

import "github.com/AwesomeXjs/tma-server/internal/service"

// Controller handles the authentication and header-related operations.
type Controller struct {
	svc service.IService
}

// New creates a new instance of the Controller.
// It takes an authentication client and a header helper as dependencies.
func New(svc service.IService) *Controller {
	return &Controller{
		svc: svc,
	}
}
