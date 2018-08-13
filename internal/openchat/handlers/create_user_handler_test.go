package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat"
	"github.com/geisonbiazus/openchat/internal/openchat/services"
	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
)

func TestCreateUserHandler(t *testing.T) {
	type fixture struct {
		service             *CreateUserServiceSpy
		handler             *CreateUserHandler
		w                   *httptest.ResponseRecorder
		validRequest        *http.Request
		validInput          services.CreateUserInput
		userJSON            string
		invalidResponseJSON string
		validOutput         services.CreateUserOutput
		invalidOutput       services.CreateUserOutput
	}

	setup := func() *fixture {
		service := NewCreateUserServiceSpy()
		handler := NewCreateUserHandler(service)
		w := httptest.NewRecorder()
		validRequestBody := `{"username":"username","password":"password","about":"about"}`
		validRequest := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(validRequestBody))
		validInput := services.CreateUserInput{
			Username: "username",
			Password: "password",
			About:    "about",
		}
		userJSON := `{"id":"id","username":"username","about":"about"}` + "\n"
		validOutput := services.CreateUserOutput{
			Status:   services.StatusSuccess,
			ID:       "id",
			Username: "username",
			About:    "about",
		}

		invalidResponseJSON := `{"errors":[{"field":"username","type":"ALREADY_TAKEN"}]}` + "\n"
		invalidOutput := services.CreateUserOutput{
			Status: services.StatusError,
			Errors: []openchat.Error{
				openchat.Error{Field: "username", Type: "ALREADY_TAKEN"},
			},
		}

		return &fixture{
			service:             service,
			handler:             handler,
			validRequest:        validRequest,
			w:                   w,
			validInput:          validInput,
			userJSON:            userJSON,
			validOutput:         validOutput,
			invalidOutput:       invalidOutput,
			invalidResponseJSON: invalidResponseJSON,
		}
	}

	t.Run("Create a user", func(t *testing.T) {
		f := setup()
		f.handler.ServeHTTP(f.w, f.validRequest)
		assert.Equal(t, f.validInput, f.service.Input)
	})

	t.Run("Respond with the created user", func(t *testing.T) {
		f := setup()
		f.service.Output = f.validOutput
		f.handler.ServeHTTP(f.w, f.validRequest)
		assert.Equal(t, http.StatusCreated, f.w.Code)
		assert.Equal(t, "application/json", f.w.Header().Get("content-type"))
		assert.Equal(t, f.userJSON, readAll(f.w.Body))
	})

	t.Run("Respond with error messages with invalid output", func(t *testing.T) {
		f := setup()
		f.service.Output = f.invalidOutput
		f.handler.ServeHTTP(f.w, f.validRequest)
		assert.Equal(t, http.StatusBadRequest, f.w.Code)
		assert.Equal(t, "application/json", f.w.Header().Get("content-type"))
		assert.Equal(t, f.invalidResponseJSON, readAll(f.w.Body))
	})
}

func readAll(r io.Reader) string {
	content, _ := ioutil.ReadAll(r)
	return string(content)
}

type CreateUserServiceSpy struct {
	Input  services.CreateUserInput
	Output services.CreateUserOutput
}

func NewCreateUserServiceSpy() *CreateUserServiceSpy {
	return &CreateUserServiceSpy{}
}

func (s *CreateUserServiceSpy) Run(input services.CreateUserInput) services.CreateUserOutput {
	s.Input = input
	return s.Output
}
