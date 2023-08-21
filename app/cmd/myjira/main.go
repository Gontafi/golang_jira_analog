package main

import (
	"context"
	app "github.com/Gontafi/golang_jira_analog"
	"github.com/Gontafi/golang_jira_analog/internal/config"
	"github.com/Gontafi/golang_jira_analog/pkg/database"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/handler"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/interfaces"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("loaded config")
	log.Info("Connecting to PostgresSQL")

	db, err := database.ConnectPSQL(database.PSQLConfig{
		Host:     cfg.HostSQL,
		Port:     cfg.PortSQL,
		User:     cfg.User,
		Password: cfg.PasswordSQL,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
	})
	if err != nil {
		log.Error("failed to connect postgres", err)
	}

	log.Info("Connected to PostgresSQL")
	log.Info("Connecting to Redis")

	rdb, err := database.ConnectRedis(database.RedisConfig{
		Addr:     cfg.AddrRedis,
		Password: cfg.PasswordRedis,
		DB:       cfg.DB,
		Protocol: cfg.Protocol,
	})
	if err != nil {
		log.Error("failed to connect redis", err)
	}

	log.Info("Connected to Redis")

	ctx := context.Background()

	repo := interfaces.NewRepository(ctx, db, rdb)
	service := interfaces.NewServices(repo)
	handlers := handler.NewHandler(service)

	log.Info("Starting server")

	srv := new(app.Server)
	err = srv.Run(cfg.Host, cfg.Port, cfg.Timeout, cfg.IdleTimeout, handlers.InitRoutes())
	if err != nil {
		log.Error("failed to run server", err)
	}
	log.Info("Server closed")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
