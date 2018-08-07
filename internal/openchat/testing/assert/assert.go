package assert

import (
	"reflect"
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
