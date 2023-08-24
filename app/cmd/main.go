package main

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/internal/config"
	"github.com/Gontafi/golang_jira_analog/pkg/database"
	"github.com/Gontafi/golang_jira_analog/pkg/handler"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
	"github.com/Gontafi/golang_jira_analog/pkg/services"
	"log"
)

func main() {

	cfg := config.MustLoad()

	db, err := database.ConnectPSQL(database.PSQLConfig{
		Host:     cfg.PsqlConfig.Host,
		Port:     cfg.PsqlConfig.Port,
		User:     cfg.PsqlConfig.User,
		Password: cfg.PsqlConfig.Password,
		DBName:   cfg.PsqlConfig.DBName,
		SSLMode:  cfg.PsqlConfig.SSLMode,
	})

	rdb, err := database.ConnectRedis(database.RedisConfig{
		Addr:     cfg.RedisConfig.Addr,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
		Protocol: cfg.RedisConfig.Protocol,
	})

	ctx := context.Background()

	repo := repos.NewRepository(ctx, db, rdb)
	service := services.NewServices(repo)
	handlers := handler.NewHandler(service)

	app := handlers.InitRoutes()

	err = app.Listen(cfg.HttpServerConfig.Host + ":" + cfg.HttpServerConfig.Port)
	if err != nil {
		log.Println("Error:", err)
	}

}
