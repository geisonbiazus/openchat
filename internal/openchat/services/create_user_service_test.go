package services

import (
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/doubles"
)

func TestCreateUserService(t *testing.T) {
	type fixture struct {
		userRepository      *doubles.UserRepositorySpy
		service             *CreateUserService
		input               CreateUserInput
		successOutput       CreateUserOutput
		usernameTakenOutput CreateUserOutput
		user                openchat.User
	}

	setup := func() *fixture {
		uuid := "a6c3f22d-d67e-4973-a98e-2026a56b3116"
		userRepository := doubles.NewUserRepositorySpy()
		user := openchat.User{
			ID:       uuid,
			Username: "username",
			Password: "password",
			About:    "about",
		}
		idGenerator := doubles.NewIDGeneratorStub(uuid)
		service := NewCreateUserService(idGenerator, userRepository)
		input := CreateUserInput{
			Username: "username",
			Password: "password",
			About:    "about",
		}

		successOutput := CreateUserOutput{
			Status:   StatusSuccess,
			ID:       uuid,
			Username: "username",
			About:    "about",
		}

		usernameTakenOutput := CreateUserOutput{
			Status: StatusError,
			Errors: []openchat.Error{
				openchat.Error{Field: "username", Type: "ALREADY_TAKEN"},
			},
		}

		return &fixture{
			userRepository:      userRepository,
			service:             service,
			input:               input,
			user:                user,
			successOutput:       successOutput,
			usernameTakenOutput: usernameTakenOutput,
		}
	}

	t.Run("Create a user", func(t *testing.T) {
		f := setup()
		f.service.Run(f.input)
		assert.DeepEqual(t, f.userRepository.CreatedUser, f.user)
	})

	t.Run("Return the created user data", func(t *testing.T) {
		f := setup()
		output := f.service.Run(f.input)
		assert.DeepEqual(t, f.successOutput, output)
	})

	t.Run("Return error message when username has already been taken", func(t *testing.T) {
		f := setup()
		f.userRepository.Create(f.user)
		output := f.service.Run(f.input)
		assert.DeepEqual(t, f.usernameTakenOutput, output)
	})
}
