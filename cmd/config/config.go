package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		MYSQL   MySql
		GRPC    GRPC
		RMQ     RMQ
		Metrics Metrics
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// MySql -.
	MySql struct {
		User   string `env:"DB_USER,required"`
		Pass   string `env:"DB_PASS,required"`
		URL    string `env:"DB_URL,required"`
		DBName string `env:"DB_NAME,required"`
	}

	// GRPC -.
	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	// Metrics -.
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	// Swagger -.
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
