package config

import (
	"fmt"
	"net"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	serverHost     = "localhost"
	defaultEnvFile = ".env"
)

// Config is the config struct.
type Config struct {
	ServerPort string `envconfig:"SERVER_PORT" default:"50051"`

	DBName     string `envconfig:"DB_NAME"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`

	SaltPassword string `envconfig:"SALT_PASSWORD"`

	RedisHost              string        `envconfig:"REDIS_HOST"`
	RedisPort              string        `envconfig:"REDIS_PORT"`
	RedisMaxIdleConns      int           `envconfig:"REDIS_MAX_IDLE"`
	RedisConnectionTimeout time.Duration `envconfig:"REDIS_CONNECTION_TIMEOUT"`
	RedisIdleTimeout       time.Duration `envconfig:"REDIS_IDLE_TIMEOUT_SEC"`
}

// New creates a new config.
func New() *Config {
	return &Config{}
}

// Init initializes the config.
func (cfg *Config) Init(path string) error {
	if path == "" {
		path = defaultEnvFile
	}

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed load env: %w", err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("failed parse env: %w", err)
	}

	return nil
}

// GetDbDNS returns the database connection string.
func (cfg *Config) DbDNS() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
}

// GetSalt returns the salt password.
func (cfg *Config) Salt() string {
	return cfg.SaltPassword
}

// GetServerAddress returns the server address.
func (cfg *Config) ServerAddress() string {
	return net.JoinHostPort(serverHost, cfg.ServerPort)
}

type redisConfig struct {
	Host              string
	Port              string
	maxIdleConns      int
	connectionTimeout time.Duration
	idleTimeout       time.Duration
}

// RedisConfig returns the redis config.
func (cfg *Config) RedisConfig() *redisConfig {
	return &redisConfig{
		Host:              cfg.RedisHost,
		Port:              cfg.RedisPort,
		maxIdleConns:      cfg.RedisMaxIdleConns,
		connectionTimeout: cfg.RedisConnectionTimeout,
		idleTimeout:       cfg.RedisIdleTimeout,
	}
}

// Address returns the redis address.
func (r redisConfig) Address() string {
	return net.JoinHostPort(r.Host, r.Port)
}

// MaxIdle returns the max idle connections.
func (r redisConfig) ConnectionTimeout() time.Duration {
	return r.connectionTimeout
}

// MaxIdle returns the max idle connections.
func (r redisConfig) MaxIdle() int {
	return int(r.maxIdleConns)
}

// IdleTimeout returns the idle timeout.
func (r redisConfig) IdleTimeout() time.Duration {
	return r.idleTimeout
}
