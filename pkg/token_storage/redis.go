package token_storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Implementation of the TokenStorage interface for Redis.
type RedisTokenStorage struct {
	ctx   context.Context
	redis *redis.Client
}

// Get redis token storage.
func NewRedisTokenStorage(ctx context.Context, redis *redis.Client) *RedisTokenStorage {
	return &RedisTokenStorage{ctx, redis}
}

// Seve token in Redis.
func (r *RedisTokenStorage) SaveToken(token string, lifeTimeInSec int) error {
	err := r.redis.Set(r.ctx, token, 0, time.Duration(lifeTimeInSec)*time.Second).Err()
	return err
}

// Delete token in Redis.
func (r *RedisTokenStorage) DeleteToken(token string) error {
	err := r.redis.Del(r.ctx, token).Err()
	return err
}

// Check for token existence in redis.
func (r *RedisTokenStorage) CheckToken(token string) (bool, error) {
	_, err := r.redis.Get(r.ctx, token).Result()
	if err == nil {
		return false, err
	}
	if err == redis.Nil {
		return false, nil
	}
	return true, nil

}

// Close redis client.
func (r *RedisTokenStorage) Close() error {
	return r.redis.Close()
}
