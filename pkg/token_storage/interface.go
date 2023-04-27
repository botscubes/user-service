package token_storage

import "context"

type TokenStorage interface {
	SaveToken(ctx context.Context, token string, lifeTime int) error
	DeleteToken(ctx context.Context, token string) error
	CheckToken(ctx context.Context, token string) (bool, error)
	Close() error
}
