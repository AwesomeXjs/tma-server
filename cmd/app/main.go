package main

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/app"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	myApp, err := app.New(ctx)
	if err != nil {
		logger.Fatal("failed to init app", zap.Error(err))
	}

	err = myApp.Run()
	if err != nil {
		logger.Fatal("failed to run app", zap.Error(err))
	}
}
