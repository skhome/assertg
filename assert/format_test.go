package assert_test

import (
	"testing"

	"github.com/skhome/assertg/assert"
)

func TestDefaultDescriptionFormatter(t *testing.T) {
	tests := []struct {
		name        string
		description assert.Description
		expected    string
	}{
		{name: "empty", description: "", expected: ""},
		{name: "text", description: "message", expected: "[message] "},
	}
	formatter := assert.DefaultDescriptionFormatter
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := formatter(test.description)
			if test.expected != actual {
				t.Errorf("expected result to be %q, but got %q", test.expected, actual)
			}
		})
	}
}

func TestCompactMessageFormatter(t *testing.T) {
	tests := []struct {
		name           string
		description    assert.Description
		representation assert.Representation
		format         string
		args           []any
		expected       string
	}{
		{
			name:           "empty",
			description:    "",
			representation: assert.DefaultRepresentation,
			format:         "expected value to be true, but was %s",
			args:           []any{false},
			expected:       "expected value to be true, but was <false>",
		},
		{
			name:           "with description",
			description:    assert.Description("is hobbit"),
			representation: assert.DefaultRepresentation,
			format:         "expected value to be true, but was %s",
			args:           []any{false},
			expected:       "[is hobbit] expected value to be true, but was <false>",
		},
		{
			name:           "with representation",
			description:    "",
			representation: assert.HexadecimalRepresentation,
			format:         "expected value to be %[1]s, but was %[2]s",
			args:           []any{42, 255},
			expected:       "expected value to be <2A>, but was <FF>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			formatter := assert.NewCompactMessageFormatter()
			actual := formatter.Format(test.description, test.representation, test.format, test.args...)
			if test.expected != actual {
				t.Errorf("expected result to be %q, but got %q", test.expected, actual)
			}
		})
	}
}
