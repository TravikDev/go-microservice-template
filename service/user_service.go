package service

import "github.com/example/user-service/model"

// UserService contains core business logic.
type UserService struct {
	users []model.User
	chars map[string][]model.Character
}

func NewUserService() *UserService {
	return &UserService{
		users: []model.User{
			{ID: "1", Name: "Alice"},
			{ID: "2", Name: "Bob"},
		},
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
	return s.users
}

// ListCharacters returns all characters for a given user.
func (s *UserService) ListCharacters(userID string) []model.Character {
	return s.chars[userID]
}
// UserService contains core business logic.
type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
