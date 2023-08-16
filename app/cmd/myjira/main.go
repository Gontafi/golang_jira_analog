package main

/*
TODO:Write handlers, services, repositories
TODO:create api, route for crud changes
TODO:sms code registration or forgot password(redis)
TODO:auth by JWT
TODO:Docker compose
TODO:github
TODO:logic fix
TODO:small testing
...
PROFIT
*/
import (
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func returner() *int {
	var asd int
	asd = 5

	return &asd
}
func main() {
	fmt.Println(*returner())
	//cfg := config.MustLoad()
	//var (
	//	port = cfg.Port
	//	host = cfg.Host
	//)
	//log := setupLogger(cfg.Env)
	//_ = log
	//_, err := database.New("postgres", "postgres", "localhost", 6543, "postgres")
	//if err != nil {
	//	slog.Error("failed to create database", err)
	//}
	//
	//err = ginS.Run(host + ":" + port)
	//if err != nil {
	//	slog.Error("failed to run server", err)
	//}
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
