package repository

import (
	"context"
	"database/sql"

	"github.com/example/user-service/model"
)

// PgUserRepository implements UserRepository backed by Postgres.
type PgUserRepository struct {
	db *sql.DB
}

// NewPgUserRepository creates a new PgUserRepository.
func NewPgUserRepository(db *sql.DB) *PgUserRepository {
	return &PgUserRepository{db: db}
}

// Create inserts a new user into the database.
func (r *PgUserRepository) Create(ctx context.Context, u model.User) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO users (id, name) VALUES ($1, $2)`, u.ID, u.Name)
	return err
}

// List returns all users from the database.
func (r *PgUserRepository) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// ByID returns a user by id.
func (r *PgUserRepository) ByID(ctx context.Context, id string) (model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(ctx, `SELECT id, name FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name)
	return u, err
}
