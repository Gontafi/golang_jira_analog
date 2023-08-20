package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type PSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	Protocol int
}

func ConnectPSQL(config PSQLConfig) (*pgx.Conn, error) {

	psqlUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)

	dbCon, err := pgx.Connect(context.Background(), psqlUrl)
	if err != nil {
		log.Fatalf("Unable to connect database %s", err)
	}

	return dbCon, nil
}

func ConnectRedis(config RedisConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Unable to connect Redis: %s", err)
	}

	return redisClient, nil
}
