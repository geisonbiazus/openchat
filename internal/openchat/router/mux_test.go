package router

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
)

func TestRouter(t *testing.T) {
	type fixture struct {
		server *httptest.Server
	}

	setup := func() *fixture {
		mux := NewMux()
		server := httptest.NewServer(mux)
		return &fixture{
			server: server,
		}
	}

	t.Run("/users/", func(t *testing.T) {
		f := setup()
		defer f.server.Close()

		requestBody := `{"username":"username","password":"password","about":"about"}`
		response, _ := http.Post(f.server.URL+"/users/", "application/json", bytes.NewBufferString(requestBody))

		responseBody := `{"id":".*?","username":"username","about":"about"}` + "\n"

		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

		assert.Match(t, responseBody, readAll(response.Body))
	})
}

func readAll(r io.Reader) string {
	content, _ := ioutil.ReadAll(r)
	return string(content)
}
