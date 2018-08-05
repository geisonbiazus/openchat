package doubles

import "github.com/geisonbiazus/openchat/internal/openchat"

type UserRepositorySpy struct {
	CreatedUser openchat.User
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}

func (r *UserRepositorySpy) Create(user openchat.User) {
	r.CreatedUser = user
}
