package redis

import (
	"github.com/botscubes/user-service/internal/config"

	"github.com/redis/go-redis/v9"
)

// Get Client for Redis.
func GetClient(c *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.DB,
	})

}
