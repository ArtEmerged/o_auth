package config

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
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

const defaultEnvFile = ".env"

func New() *config {
	return &config{}
}

func (cfg *config) Init(path string) {
	if path == "" {
		path = defaultEnvFile
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Println("failed load env")
		panic(err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		log.Println("failed parse env")
		panic(err)
	}
}

func (cfg *config) GetDbDNS() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
}

func (cfg *config) GetServerAddress() string {
	return net.JoinHostPort("localhost", cfg.ServerPort)
}
