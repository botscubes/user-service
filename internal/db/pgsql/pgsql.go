package pgsql

import (
	"context"
	"fmt"

	"github.com/botscubes/user-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// New Postgresql pool.
// Сlose after use.
func NewPool(context context.Context, c *config.DBConfig) (*pgxpool.Pool, error) {
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBname,
	)
	pool, err := pgxpool.New(context, connURL)
	if err != nil {
		return nil, err
	}

	return pool, err
}
