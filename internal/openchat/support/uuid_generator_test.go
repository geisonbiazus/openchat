package support

import (
	"regexp"
	"testing"
)

const uuidv4Regex = "(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"

func TestUUIDGenerator(t *testing.T) {
	t.Run("Generate a UUID", func(t *testing.T) {
		generator := NewUUIDGenerator()
		uuid := generator.Generate()

		matched, _ := regexp.MatchString(uuidv4Regex, uuid)
		if matched != true {
			t.Errorf("Expected %v to match %v.", uuid, uuidv4Regex)
		}
	})
}
