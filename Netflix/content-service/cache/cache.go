package cache

import (
	"content-service/config"

	"github.com/go-redis/redis/v8"
)

func NewClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
}
