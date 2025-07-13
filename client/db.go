package client

import (
	"database/sql"

	"github.com/example/user-service/config"
)

// DB returns a database connection.
func DB(cfg config.Config) (*sql.DB, error) {
	return sql.Open("postgres", cfg.DBURL)
}
