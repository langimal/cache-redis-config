package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host string
	Port int
	DB   int
}

func GetConfig() Config {
	if os.Getenv("REDIS_HOST")!= "" {
		return Config{
			Host: os.Getenv("REDIS_HOST"),
			Port: 6379,
			DB:   0,
		}
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err!= nil {
		log.Fatal(err)
	}

	configPath := filepath.Join(dir, "config", "redis.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Redis config file not found")
	}

	var config Config
	err = os.ReadFile(configPath).UnmarshalJSON(&config)
	if err!= nil {
		log.Fatal(err)
	}

	return config
}

func NewRedisClient(config Config) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: "", // no password set
		DB:       config.DB, // use default DB
	})
}

func NewRedisClientFromEnv() (*redis.Client, error) {
	return NewRedisClient(GetConfig())
}