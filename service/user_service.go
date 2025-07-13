package service

import (
	"context"

	"github.com/example/user-service/model"
	"github.com/example/user-service/repository"
)

// UserService contains core business logic.
type UserService struct {
	repo  repository.UserRepository
	chars map[string][]model.Character
}

// NewUserService constructs a UserService with the provided repository.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
		chars: map[string][]model.Character{
			"1": {
				{ID: "c1", UserID: "1", Name: "Warrior"},
				{ID: "c2", UserID: "1", Name: "Mage"},
			},
			"2": {
				{ID: "c3", UserID: "2", Name: "Rogue"},
			},
		},
	}
}

// ListUsers returns all users.
func (s *UserService) ListUsers() []model.User {
	users, err := s.repo.List(context.Background())
	if err != nil {
		return nil
	}
	return users
}

// ListCharacters returns all characters for a given user.
func (s *UserService) ListCharacters(userID string) []model.Character {
	return s.chars[userID]
}
