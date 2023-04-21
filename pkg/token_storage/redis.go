package token_storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// implements the TokenStorage interface.
type RedisTokenStorage struct {
	redis *redis.Client
	ctx   *context.Context
}

// Get Client for Redis.
func NewRedisTokenStorage(redis *redis.Client, ctx *context.Context) *RedisTokenStorage {
	return &RedisTokenStorage{redis, ctx}
}

func (r *RedisTokenStorage) SaveToken(token string, lifeTimeInSec int) error {
	err := r.redis.Set(*r.ctx, token, 0, time.Duration(lifeTimeInSec)*time.Second).Err()
	return err
}
func (r *RedisTokenStorage) DeleteToken(token string) error {
	err := r.redis.Del(*r.ctx, token).Err()
	return err
}
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
