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

type config struct {
	ServerPort string `envconfig:"SERVER_PORT" default:"50051"`

	DBName     string `envconfig:"DB_NAME"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`

	SaltPassword string `envconfig:"SALT_PASSWORD"`
}

// New creates a new config.
func New() *config {
	return &config{}
}

// Init initializes the config.
func (cfg *config) Init(path string) error {
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
func (cfg *config) GetDbDNS() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
}

// GetServerAddress returns the server address.
func (cfg *config) GetServerAddress() string {
	return net.JoinHostPort(serverHost, cfg.ServerPort)
}
