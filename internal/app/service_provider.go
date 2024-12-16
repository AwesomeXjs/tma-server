package app

import (
	"context"

	"github.com/AwesomeXjs/tma-server/internal/client/redis"
	"github.com/AwesomeXjs/tma-server/internal/client/redis/go_redis"
	"github.com/AwesomeXjs/tma-server/internal/config"
	"github.com/AwesomeXjs/tma-server/internal/controller"
	"github.com/AwesomeXjs/tma-server/internal/repository"
	"github.com/AwesomeXjs/tma-server/internal/service"
	"github.com/AwesomeXjs/tma-server/pkg/closer"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient"
	"github.com/AwesomeXjs/tma-server/pkg/dbClient/pg"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"go.uber.org/zap"
)

type ServiceProvider struct {
	httpConfig  config.IHTTPConfig
	dbConfig    config.PGConfig
	redisConfig redis.IRedisConfig

	dbClient    dbClient.Client
	redisClient redis.IRedis

	controller *controller.Controller
	svc        service.IService
	repo       repository.IRepository
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

// HTTPConfig returns the HTTP configuration, initializing it if necessary.
func (s *ServiceProvider) HTTPConfig() config.IHTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			logger.Fatal("failed to get http config", zap.Error(err))
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

// PGConfig initializes and returns the PostgresSQL configuration if not already set.
func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.dbConfig == nil {
		cfg, err := config.NewPgConfig()
		if err != nil {
			logger.Fatal("failed to get pg config", zap.Error(err))
		}
		s.dbConfig = cfg
	}
	return s.dbConfig
}

func (s *ServiceProvider) RedisConfig() redis.IRedisConfig {
	if s.redisConfig == nil {
		cfg, err := redis.NewRedisConfig()
		if err != nil {
			logger.Fatal("failed to get redis config", zap.Error(err))
		}
		s.redisConfig = cfg
	}
	return s.redisConfig
}

// RedisClient initializes and returns the Redis client if not already created.
// It also pings Redis to ensure the connection is valid.
func (s *ServiceProvider) RedisClient(ctx context.Context) redis.IRedis {
	if s.redisClient == nil {
		redisClient := go_redis.NewGoRedisClient(s.RedisConfig())
		closer.Add(redisClient.Client.Close)

		err := redisClient.Client.Ping(ctx).Err()
		if err != nil {
			logger.Error("Failed to connect to redis", zap.Error(err))
		}

		s.redisClient = redisClient
	}
	return s.redisClient
}

// DBClient initializes and returns the database client if not already created.
// It also pings the database to ensure the connection is valid.
func (s *ServiceProvider) DBClient(ctx context.Context) dbClient.Client {
	if s.dbClient == nil {
		cfg := s.PGConfig()
		dbc, err := pg.New(ctx, cfg.GetDSN())
		if err != nil {
			logger.Fatal("failed to get db client", zap.Error(err))
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping db", zap.Error(err))
		}

		closer.Add(dbc.Close) // Ensures the database client is closed on shutdown
		s.dbClient = dbc
	}
	return s.dbClient
}

func (s *ServiceProvider) Repository(ctx context.Context) repository.IRepository {
	if s.repo == nil {
		s.repo = repository.New(s.DBClient(ctx), s.RedisClient(ctx))
	}
	return s.repo
}

func (s *ServiceProvider) Service(ctx context.Context) service.IService {
	if s.svc == nil {
		s.svc = service.New(s.Repository(ctx))
	}
	return s.svc
}

func (s *ServiceProvider) Controller(ctx context.Context) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.New(s.Service(ctx))
	}
	return s.controller
}
