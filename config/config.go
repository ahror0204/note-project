package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	LogLevel string
	HTTPport string

	PostgresHost string
	PostgresPort int
	PostgresDatabase string
	PostgresUser string
	PostgresPassword string

	RedisHost string
	RedisPort int
}

func Load() Config {
	c := Config{}

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPport = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8282"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "test"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
