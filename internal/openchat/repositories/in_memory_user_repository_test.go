package repositories

import (
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
)

func TestInMemoryUserRepository(t *testing.T) {
	type fixture struct {
		repository *InMemoryUserRepository
		user       openchat.User
	}

	setup := func() *fixture {
		repository := NewInMemoryUserRepository()
		user := openchat.User{
			ID:       "1",
			Username: "username",
			Password: "password",
			About:    "about",
		}

		return &fixture{
			repository: repository,
			user:       user,
		}
	}

	t.Run("Create", func(t *testing.T) {
		t.Run("Create a user", func(t *testing.T) {
			f := setup()
			f.repository.Create(f.user)
			assert.Equal(t, f.user, f.repository.FindByID(f.user.ID))
		})
	})

	t.Run("FindByUsername", func(t *testing.T) {
		t.Run("Search a user by username", func(t *testing.T) {
			f := setup()
			f.repository.Create(f.user)
			assert.Equal(t, f.user, f.repository.FindByUsername(f.user.Username))
		})

		t.Run("Return NoUSer if user is not found", func(t *testing.T) {
			f := setup()
			assert.Equal(t, openchat.NoUser, f.repository.FindByUsername(f.user.Username))
		})
	})
}
