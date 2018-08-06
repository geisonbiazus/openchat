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
		successOutput  CreateUserOutput
		user           openchat.User
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
			Status:   "SUCCESS",
			ID:       uuid,
			Username: "username",
			About:    "about",
		}

		return &fixture{
			userRepository: userRepository,
			service:        service,
			input:          input,
			user:           user,
			successOutput:  successOutput,
		}
	}

	t.Run("Create a user", func(t *testing.T) {
		f := setup()
		f.service.Run(f.input)
		assert.Equal(t, f.userRepository.CreatedUser, f.user)
	})

	t.Run("Return the created user data", func(t *testing.T) {
		f := setup()
		output := f.service.Run(f.input)
		assert.Equal(t, f.successOutput, output)
	})
}
