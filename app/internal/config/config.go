package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	HttpServerConfig HttpServerConfig `mapstructure:"http_server"`
	PsqlConfig       PsqlConfig       `mapstructure:"psql_config"`
	RedisConfig      RedisConfig      `mapstructure:"redis_config"`
}

type HttpServerConfig struct {
	Host        string
	Port        string
	Timeout     time.Duration
	IdleTimeout time.Duration
	ReadTimeout time.Duration
}

type PsqlConfig struct {
	User     string
	Password string `mapstructure:"password-sql"`
	Host     string `mapstructure:"host-sql"`
	Port     int    `mapstructure:"port-sql"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type RedisConfig struct {
	Addr     string
	Password string `mapstructure:"password-redis"`
	DB       int
	Protocol int
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Println("Config path is not set")
		return &Config{}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists : %s", err)
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("prod")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("failed to read config, Error:", err)
		return &Config{}
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	fmt.Println(cfg.PsqlConfig.Host)
	fmt.Println(cfg.PsqlConfig.Port)
	return &cfg
}
