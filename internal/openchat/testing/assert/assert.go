package assert

import (
	"reflect"
	"regexp"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("\nExpected: %v\n  Actual: %v", expected, actual)
	}
}

func DeepEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected: %v\n  Actual: %v", expected, actual)
	}
}

func Match(t *testing.T, regex, actual string) {
	t.Helper()

	matched, _ := regexp.MatchString(regex, actual)
	if matched != true {
		t.Errorf("Expected %v to match %v.", actual, regex)
	}
}
