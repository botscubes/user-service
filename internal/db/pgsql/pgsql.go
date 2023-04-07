package pgsql

import (
	"database/sql"

	"github.com/botscubes/user-service/internal/config"

	_ "github.com/lib/pq"
)

// Open Postgresql connection.
// Сlose after use.
func Open(c *config.DBConfig) (*sql.DB, error) {
	connStr := "user=" + c.User +
		" password=" + c.Password +
		" dbname=" + c.DBname +
		" port=" + c.Port +
		" host=" + c.Host
	db, err := sql.Open("postgres", connStr)

	return db, err
}
