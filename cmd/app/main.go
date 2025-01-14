package main

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/app"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"go.uber.org/zap"
)

const (
	MarkMain = "Main"
)

// @title TMA API
// @version 1.0
// @description API Server for Authentication
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name TMA
func main() {
	ctx := context.Background()

	myApp, err := app.New(ctx)
	if err != nil {
		logger.Fatal("failed to init app", MarkMain, zap.Error(err))
	}

	err = myApp.Run()
	if err != nil {
		logger.Fatal("failed to run app", MarkMain, zap.Error(err))
	}
}
