package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"
	"errors"
	"github.com/example/user-service/api/router"
	"github.com/example/user-service/model"
	"github.com/example/user-service/service"
)

type fakeUserRepo struct {
	users []model.User
}

func (f *fakeUserRepo) Create(ctx context.Context, u model.User) error {
	f.users = append(f.users, u)
	return nil
}

func (f *fakeUserRepo) List(ctx context.Context) ([]model.User, error) {
	return f.users, nil
}

func (f *fakeUserRepo) ByID(ctx context.Context, id string) (model.User, error) {
	for _, u := range f.users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("not found")
}

func setupRouter() http.Handler {
	repo := &fakeUserRepo{users: []model.User{{ID: "1", Name: "Alice"}}}
	svc := service.NewUserService(repo)
	return router.New(svc)
}

func TestPingEndpoint(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", w.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp["message"] != "pong" {
		t.Errorf("expected pong, got %s", resp["message"])
	}
}

func TestListUsersEndpoint(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", w.Code)
	}
	var users []model.User
	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 || users[0].Name != "Alice" {
		t.Fatalf("unexpected users response: %+v", users)
	}
}

func TestListCharactersEndpoint(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest("GET", "/users/1/characters", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", w.Code)
	}
	var chars []model.Character
	if err := json.NewDecoder(w.Body).Decode(&chars); err != nil {
		t.Fatal(err)
	}
	if len(chars) != 2 {
		t.Fatalf("expected 2 characters, got %d", len(chars))
	}
}
