package token_storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Implementation of the TokenStorage interface for Redis.
type RedisTokenStorage struct {
	redis *redis.Client
}

// Get redis token storage.
func NewRedisTokenStorage(redis *redis.Client) *RedisTokenStorage {
	return &RedisTokenStorage{redis}
}

// Seve token in Redis.
func (r *RedisTokenStorage) SaveToken(ctx context.Context, token string, lifeTimeInSec int) error {
	err := r.redis.Set(ctx, token, 0, time.Duration(lifeTimeInSec)*time.Second).Err()
	return err
}

// Delete token in Redis.
func (r *RedisTokenStorage) DeleteToken(ctx context.Context, token string) error {
	err := r.redis.Del(ctx, token).Err()
	return err
}

// Check for token existence in redis.
func (r *RedisTokenStorage) CheckToken(ctx context.Context, token string) (bool, error) {
	_, err := r.redis.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil

}

// Close redis client.
func (r *RedisTokenStorage) Close() error {
	return r.redis.Close()
}
