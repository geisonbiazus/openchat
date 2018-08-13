package support

import (
	"testing"

	"github.com/geisonbiazus/openchat/internal/openchat/testing/assert"
)

const uuidv4Regex = "(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"

func TestUUIDGenerator(t *testing.T) {
	t.Run("Generate a UUID", func(t *testing.T) {
		generator := NewUUIDGenerator()
		uuid := generator.Generate()
		assert.Match(t, uuidv4Regex, uuid)
	})
}
