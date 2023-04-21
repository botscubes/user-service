package token_storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Implementation of the TokenStorage interface for Redis.
type RedisTokenStorage struct {
	redis *redis.Client
	ctx   *context.Context
}

// Get redis token storage.
func NewRedisTokenStorage(redis *redis.Client, ctx *context.Context) *RedisTokenStorage {
	return &RedisTokenStorage{redis, ctx}
}

// Seve token in Redis.
func (r *RedisTokenStorage) SaveToken(token string, lifeTimeInSec int) error {
	err := r.redis.Set(*r.ctx, token, 0, time.Duration(lifeTimeInSec)*time.Second).Err()
	return err
}

// Delete token in Redis.
func (r *RedisTokenStorage) DeleteToken(token string) error {
	err := r.redis.Del(*r.ctx, token).Err()
	return err
}

// Check for token existence in redis.
func (r *RedisTokenStorage) CheckToken(token string) (bool, error) {
	_, err := r.redis.Get(*r.ctx, token).Result()
	if err == nil {
		return false, err
	}
	if err == redis.Nil {
		return false, nil
	}
	return true, nil

}
