package redis

import (
	"github.com/botscubes/user-service/internal/config"

	"github.com/go-redis/redis"
)

// Get Client for Redis.
func getClient(c *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.DB,
	})

}
