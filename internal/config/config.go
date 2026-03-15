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
	Auth     Auth
	HTTP     HTTP
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
	SessionTTL     time.Duration `env:"SESSION_TTL" env-default:"168h"`
	UpdateInterval time.Duration `env:"UPDATE_INTERVAL" env-default:"5m"`
}

type HTTP struct {
	SecurityHeaders SecurityHeaders
	CORS            CORS
	RateLimit       RateLimit
	BodyLimit       BodyLimit
}

type SecurityHeaders struct {
	EnableHSTS bool `env:"SECURITY_HEADERS_ENABLE_HSTS" env-default:"true"`
	HSTSMaxAge int  `env:"SECURITY_HEADERS_HSTS_MAX_AGE" env-default:"31536000"`
}

type CORS struct {
	AllowedOrigins   []string `env:"HTTP_CORS_ALLOWED_ORIGINS" env-separator:","`
	AllowedMethods   []string `env:"HTTP_CORS_ALLOWED_METHODS" env-separator:","`
	AllowedHeaders   []string `env:"HTTP_CORS_ALLOWED_HEADERS" env-separator:","`
	ExposedHeaders   []string `env:"HTTP_CORS_EXPOSED_HEADERS" env-separator:","`
	AllowCredentials bool     `env:"HTTP_CORS_ALLOW_CREDENTIALS" env-default:"true"`
	MaxAge           int      `env:"HTTP_CORS_MAX_AGE" env-default:"300"`
}

type RateLimit struct {
	GlobalPerMinute int `env:"HTTP_RATE_LIMIT_GLOBAL_PER_MINUTE" env-default:"120"`

	LoginIPRequests int           `env:"HTTP_RATE_LIMIT_LOGIN_IP_REQUESTS" env-default:"20"`
	LoginIPWindow   time.Duration `env:"HTTP_RATE_LIMIT_LOGIN_IP_WINDOW" env-default:"1m"`

	LoginEmailRequests int           `env:"HTTP_RATE_LIMIT_LOGIN_EMAIL_REQUESTS" env-default:"10"`
	LoginEmailWindow   time.Duration `env:"HTTP_RATE_LIMIT_LOGIN_EMAIL_WINDOW" env-default:"15m"`

	LoginIPEmailRequests int           `env:"HTTP_RATE_LIMIT_LOGIN_IPEMAIL_REQUESTS" env-default:"5"`
	LoginIPEmailWindow   time.Duration `env:"HTTP_RATE_LIMIT_LOGIN_IPEMAIL_WINDOW" env-default:"10m"`

	RegisterIPRequests int           `env:"HTTP_RATE_LIMIT_REGISTER_IP_REQUESTS" env-default:"3"`
	RegisterIPWindow   time.Duration `env:"HTTP_RATE_LIMIT_REGISTER_IP_WINDOW" env-default:"10m"`

	RegisterEmailRequests int           `env:"HTTP_RATE_LIMIT_REGISTER_EMAIL_REQUESTS" env-default:"3"`
	RegisterEmailWindow   time.Duration `env:"HTTP_RATE_LIMIT_REGISTER_EMAIL_WINDOW" env-default:"1h"`
}

type BodyLimit struct {
	AuthBytes int64 `env:"HTTP_BODY_LIMIT_AUTH_BYTES" env-default:"16384"`
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
