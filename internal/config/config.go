package config

import (
	"fmt"
	"net"

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
func (cfg *Config) GetDbDNS() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
}

// GetSalt returns the salt password.
func (cfg *Config) GetSalt() string {
	return cfg.SaltPassword
}

// GetServerAddress returns the server address.
func (cfg *Config) GetServerAddress() string {
	return net.JoinHostPort(serverHost, cfg.ServerPort)
}
