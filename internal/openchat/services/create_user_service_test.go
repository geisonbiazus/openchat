package services

import (
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/doubles"
)

func TestCreateUserService(t *testing.T) {
	type fixture struct {
		userRepository *doubles.UserRepositorySpy
		service        *CreateUserService
		input          CreateUserInput
		user           openchat.User
	}

	setup := func() *fixture {
		userRepository := doubles.NewUserRepositorySpy()
		user := openchat.User{
			Username: "username",
			Password: "password",
			About:    "about",
		}
		service := NewCreateUserService(userRepository)
		input := CreateUserInput{
			Username: "username",
			Password: "password",
			About:    "about",
		}

		return &fixture{
			userRepository: userRepository,
			service:        service,
			input:          input,
			user:           user,
		}
	}

	t.Run("Create a user", func(t *testing.T) {
		f := setup()
		f.service.Run(f.input)
		assert.Equal(t, f.userRepository.CreatedUser, f.user)
	})
}
