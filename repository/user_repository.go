package repository

import (
	"context"

	"github.com/example/user-service/model"
)

// UserRepository defines persistence operations for users.
type UserRepository interface {
	// Create persists a new user.
	Create(ctx context.Context, u model.User) error
	// List returns all users.
	List(ctx context.Context) ([]model.User, error)
	// ByID finds a user by identifier.
	ByID(ctx context.Context, id string) (model.User, error)
}
