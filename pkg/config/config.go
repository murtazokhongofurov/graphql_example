package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment        string
	LogLevel           string
	PostgresHost       string
	PostgresPort       string
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	ProductServiceHost string
	ProductServicePort string
	PGXPoolMax         int
}

func Load() *Config {
	c := &Config{}
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "productdata"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "2002"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "developer"))
	c.ProductServiceHost = cast.ToString(GetOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.ProductServicePort = cast.ToString(GetOrReturnDefault("PRODUCT_SERVICE_PORT", "8080"))
	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
