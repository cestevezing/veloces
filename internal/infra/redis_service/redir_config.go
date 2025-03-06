package redis_service

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
