package response

import (
	"time"

	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Body struct defines the structure of the response body.
// It includes a title, detail about the error or message, the request URI, and the current time.
type Body struct {
	Title   string `json:"title"`
	Detail  string `json:"detail"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

// Response is a utility function to send JSON responses.
// It accepts the HTTP context, a status code, a message, and detailed information to send in the response.
func Response(ctx echo.Context, statusCode int, message, detail string) error {
	// Construct the response object using the provided parameters.
	err := ctx.JSON(statusCode, Body{
		Title:   message,
		Detail:  detail,
		Request: ctx.Request().RequestURI,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	})

	// If there's an error in writing the response, log it and return the error.
	if err != nil {
		logger.Info("failed to write response", zap.String("error", err.Error()))
		return err
	}

	// Return nil indicating that the response was successfully written.
	return nil
}
