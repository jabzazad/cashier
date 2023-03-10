package validator

import (
	"testing"
)

func TestValidPassword(t *testing.T) {
	var testObjects = []struct {
		input    string
		expected bool
	}{
		{"password", false},
		{"Password", false},
		{"Pass", false},
		{"--------", false},
		{"Passwor4", true},
		{"00000087", false},
	}
	for _, testObject := range testObjects {
		check := validPassword(testObject.input)
		if check != testObject.expected {
			t.Errorf("expected: %v, received: %v, results: %v", testObject.expected, testObject.input, check)
		}
	}
}
