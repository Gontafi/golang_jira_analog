package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" enn-default:"local"`
	HTTPServer  `yaml:"http_server"`
	PSQLConfig  `yaml:"psql_config"`
	RedisConfig `yaml:"redis_config"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" end-default:"60s"`
}

type PSQLConfig struct {
	User        string `yaml:"user" env-default:"postgres"`
	PasswordSQL string `yaml:"password-sql" env-default:"postgres"`
	HostSQL     string `yaml:"host-sql" env-default:"localhost"`
	PortSQL     int    `yaml:"port-sql" env-default:"6543"`
	DBName      string `yaml:"db_name" env-default:"postgres"`
	SSLMode     string `yaml:"ssl-mode" env-default:"disable"`
}

type RedisConfig struct {
	AddrRedis     string `yaml:"addr" env-default:"localhost:6379"`
	PasswordRedis string `yaml:"password-redis" env-default:""`
	DB            int    `yaml:"db" env-default:"0"`
	Protocol      int    `yaml:"protocol" enf-default:"3"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists : %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config %s", err)
	}

	return &cfg
}
