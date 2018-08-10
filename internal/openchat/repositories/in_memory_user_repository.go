package repositories

import "github.com/geisonbiazus/openchat/internal/openchat"

type InMemoryUserRepository struct {
	users map[string]openchat.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]openchat.User),
	}
}

func (r *InMemoryUserRepository) Create(user openchat.User) {
	r.users[user.ID] = user
}

func (r *InMemoryUserRepository) FindByUsername(username string) openchat.User {
	for _, u := range r.users {
		if u.Username == username {
			return u
		}
	}
	return openchat.NoUser
}

func (r *InMemoryUserRepository) FindByID(id string) openchat.User {
	return r.users[id]
}
