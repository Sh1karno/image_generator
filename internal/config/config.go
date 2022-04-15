package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const PortEnvVar = "SERVER_PORT"

type Config struct {
	port string
}

func New() *Config {
	absPath, _ := filepath.Abs("./internal/config/.env")

	err := godotenv.Load(absPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv(PortEnvVar)
	return &Config{
		port: port,
	}
}

func (c *Config) GetPort() string {
	return c.port
}
