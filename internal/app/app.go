package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/AwesomeXjs/tma-server/internal/middlewares"
	"github.com/AwesomeXjs/tma-server/pkg/closer"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"os"
)

const (
	// EnvPath is the path to the .env file that contains environment variables.
	EnvPath = ".env"
)

// logLevel is a command-line flag for specifying the log level.
var logLevel = flag.String("l", "info", "log level")
var mode = flag.String("m", "dev", "build mode")

type App struct {
	ServiceProvider *ServiceProvider
	server          *echo.Echo
}

// New creates and initializes the App with dependencies.
func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		// Fatal log in case of failure during dependency initialization
		logger.Fatal("failed to init deps", zap.Error(err))
	}
	return app, nil
}
func (app *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		app.InitConfig,
		app.initServiceProvider,
		app.InitEchoServer,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			logger.Error("failed to init deps", zap.Error(err))
			return err
		}
	}

	app.InitRoutes(ctx, app.server)
	return nil
}

// Run starts the HTTP server and handles cleanup on shutdown.
func (app *App) Run() error {
	defer func() {
		closer.CloseAll() // Close all services/resources
		closer.Wait()     // Wait for all services to close
	}()
	err := app.runHTTPServer() // Run the HTTP server
	if err != nil {
		logger.Fatal("failed to run http server", zap.Error(err))
	}
	return nil
}

// InitConfig loads environment variables for the application.
func (app *App) InitConfig(_ context.Context) error {
	err := godotenv.Load(EnvPath)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", EnvPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return err
}

// InitEchoServer sets up the Echo server and its middleware.
func (app *App) InitEchoServer(_ context.Context) error {
	flag.Parse()                                                 // Parse command-line flags
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel))) // Initialize logger with the specified log level

	app.server = echo.New()              // Create a new Echo server
	app.server.Use(middleware.Recover()) // Middleware for recovering from panics
	app.server.Use(middlewares.Logger)   // Custom logging middleware

	app.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://awesomexjs.github.io", "http://192.168.0.105:3000"},                     // Allow CORS from this origin
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS}, // Allowed HTTP methods
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlRequestHeaders,
			echo.HeaderAccessControlAllowOrigin,
		}, // Allowed headers for CORS
	}))
	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.ServiceProvider = NewServiceProvider()
	return nil
}

// runHTTPServer starts the Echo server and listens for requests.
func (app *App) runHTTPServer() error {
	logger.Info("server listening at %v", zap.String("start", app.ServiceProvider.HTTPConfig().Address())) // Log the server address
	return app.server.Start(app.ServiceProvider.HTTPConfig().Address())                                    // Start the server at the configured address
}

// InitRoutes sets up the application routes.
func (app *App) InitRoutes(ctx context.Context, server *echo.Echo) {
	app.ServiceProvider.Controller(ctx).InitRoutes(server, checkMode(mode)) // Initialize routes using the controller
}

func checkMode(mode *string) string {
	var token string
	switch *mode {
	case "dev":
		logger.Warn("DEV MODE ACTIVATED")
		token = os.Getenv("TEST_BOT_TOKEN")
	case "prod":
		logger.Warn("PRODUCTION MODE ACTIVATED")
		token = os.Getenv("REAL_BOT_TOKEN")
	default:
		token = os.Getenv("TEST_BOT_TOKEN")
	}

	return token
}
