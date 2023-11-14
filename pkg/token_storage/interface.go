package token_storage

import (
	"context"
	"time"
)

type TokenStorage interface {
	SaveToken(ctx context.Context, token string, lifeDuration time.Duration) error
	DeleteToken(ctx context.Context, token string) error
	CheckToken(ctx context.Context, token string) (bool, error)
	Close() error
}
