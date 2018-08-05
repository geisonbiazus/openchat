package assert

import "testing"

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("\nExpected: %v\n  Actual: %v", expected, actual)
	}
}
