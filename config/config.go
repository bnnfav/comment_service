package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	Environment      string
	LogLevel         string

	UserServiceHost string
	UserServicePort string

	ProductServiceHost string
	ProductServicePort string

	OrderServiceHost string
	OrderServicePort string

	CommentServiceHost string
	CommentServicePort string
}

func Load() Config {
	c := Config{}

	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "bay"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "123"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "commentdb"))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SERVICE_PORT", "8000"))

	c.ProductServiceHost = cast.ToString(getOrReturnDefault("comment_service_HOST", "localhost"))
	c.ProductServicePort = cast.ToString(getOrReturnDefault("comment_service_PORT", "8010"))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "localhost"))
	c.OrderServicePort = cast.ToString(getOrReturnDefault("ORDER_SERVICE_PORT", "8020"))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "localhost"))
	c.CommentServicePort = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_PORT", ":8030"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
