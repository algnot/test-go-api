package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	AppPort string `envconfig:"APP_PORT" default:"3000"`
	Service string `envconfig:"APP_SERVICE" default:"server"`
}

type DatabaseConfig struct {
	Host   string `envconfig:"DB_HOST" default:"localhost"`
	Port   int    `envconfig:"DB_PORT" default:"3306"`
	User   string `envconfig:"DB_USER" default:"root"`
	Pass   string `envconfig:"DB_PASSWORD" default:""`
	DBName string `envconfig:"DB_NAME" default:"test-go"`
}

func GetConfig() Config {
	var config Config

	err := godotenv.Load()
	if err != nil {
		log.Panicln("No .env file")
	}

	envconfig.MustProcess("app", &config.Server)
	envconfig.MustProcess("database", &config.Database)
	return config
}
