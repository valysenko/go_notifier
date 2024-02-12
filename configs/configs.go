package configs

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppEnv string `env:"APP_ENV" env-default:"loc"`
	DBConfig
	HttpServerConfig
	RabbitConfig
}

type DBConfig struct {
	Host            string `env:"DB_HOST"`
	Port            string `env:"DB_PORT"`
	Username        string `env:"DB_USERNAME" `
	Password        string `env:"DB_PASSWORD"`
	DbName          string `env:"DB_NAME"`
	MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type RabbitConfig struct {
	Host     string `env:"RABBITMQ_HOST" env-default:"localhost"`
	Port     string `env:"RABBITMQ_PORT" env-default:"5672"`
	User     string `env:"RABBITMQ_USER" env-default:"guest"`
	Password string `env:"RABBITMQ_PASSWORD" env-default:"guest"`
}

func (cfg *RabbitConfig) ProvideDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		url.QueryEscape(cfg.User),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
	)
}

type HttpServerConfig struct {
	ServerPort string `env:"HTTP_PORT"`
}

func (cfg *DBConfig) ProvideDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
}

func InitConfig() *AppConfig {
	cfg := &AppConfig{}
	godotenv.Load("deployments/.env") // TODO: figure out, if run go run cmd/main.go
	// godotenv.Load("../deployments/.env") // TODO: figure out, if run go run main.go

	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatalf("Error reading environment variables: %v", err)
		help, _ := cleanenv.GetDescription(&cfg, nil)
		fmt.Println(help)
		panic(err)
	}
	return cfg
}
