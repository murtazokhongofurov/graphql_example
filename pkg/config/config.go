package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment        string
	LogLevel           string
	MongoHost          string
	MongoPort          string
	MongoDatabase      string
	MongoUser          string
	MongoPassword      string
	ProductServiceHost string
	ProductServicePort string
}

func Load() *Config {
	c := &Config{}
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.MongoHost = cast.ToString(GetOrReturnDefault("MONGO_HOST", "127.0.0.1"))
	c.MongoPort = cast.ToString(GetOrReturnDefault("MONGO_PORT", 27017))
	c.MongoDatabase = cast.ToString(GetOrReturnDefault("MONGO_DATABASE", "productdb"))
	c.MongoPassword = cast.ToString(GetOrReturnDefault("MONGO_PASSWORD", "2002"))
	c.MongoUser = cast.ToString(GetOrReturnDefault("MONGO_USER", "developer"))
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
