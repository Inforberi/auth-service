package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppEnv   string `env:"APP_ENV" default:"prod"`
	Logger   Logger
	Postgres Postgres
}

type Logger struct {
	Level  string `env:"LOG_LEVEL" default:"info"`
	Format string `env:"LOG_FORMAT" default:"json"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Port     int    `env:"POSTGRES_PORT" env-required:"true"`
	User     string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBName   string `env:"POSTGRES_NAME_DB" env-required:"true"`
	SSLMode  string `env:"PG_SSL_MODE" env-default:"disable"`

	// Пул (опционально)
	MaxConns        int32         `env:"PG_MAX_CONNS" env-default:"20"`
	MinConns        int32         `env:"PG_MIN_CONNS" env-default:"2"`
	HealthTimeout   time.Duration `env:"PG_HEALTH_TIMEOUT" env-default:"3s"`
	MaxConnLifetime time.Duration `env:"PG_MAX_CONN_LIFETIME" env-default:"30m"`
	MaxConnIdleTime time.Duration `env:"PG_MAX_CONN_IDLE_TIME" env-default:"5m"`
}

type Auth struct {
	SessionTTL time.Duration `env:"SESSION_TTL" env-default:"168h"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}

func (p Postgres) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode,
	)
}
