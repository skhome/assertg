package assert

import (
	"fmt"
	"testing"
)

func TestStringDescription(t *testing.T) {
	// given
	fixture := new(fixtureT)

	// when
	assert := ThatString(fixture, "foo").As("custom description for %s", "bar")

	// then
	expected := "custom description for bar"
	if assert.description != expected {
		t.Errorf("expected description to be %q, but got %q", expected, assert.description)
	}
}

func TestStringIsEmpty(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: ""},
		{actual: " ", expectedFail: true},
		{actual: "foo", expectedFail: true},
	}
	messageFormat := "expected string to be empty, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsEmpty()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringIsBlank(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: ""},
		{actual: " "},
		{actual: "\t"},
		{actual: " \t"},
		{actual: "foo", expectedFail: true},
	}
	messageFormat := "expected string to be blank, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsBlank()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringMatches(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: "foo"},
		{actual: "bar", expectedFail: true},
	}
	predicate := func(str string) bool { return str[0] == 'f' }
	messageFormat := "expected string to match predicate, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).Matches(predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringDoesNotMatches(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: "bar"},
		{actual: "foo", expectedFail: true},
	}
	predicate := func(str string) bool { return str[0] == 'f' }
	messageFormat := "expected string to not match predicate, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotMatch(predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringIsEqualTo(t *testing.T) {
	tests := []struct {
		actual       string
		expected     string
		expectedFail bool
	}{
		{actual: "foo", expected: "foo"},
		{actual: "foo", expected: "bar", expectedFail: true},
	}
	messageFormat := "expected string to equal %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsEqualTo(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringIsNotEqualTo(t *testing.T) {
	tests := []struct {
		actual       string
		expected     string
		expectedFail bool
	}{
		{actual: "foo", expected: "bar"},
		{actual: "foo", expected: "foo", expectedFail: true},
	}
	messageFormat := "expected string to not equal %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsNotEqualTo(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringIsIn(t *testing.T) {
	tests := []struct {
		actual       string
		expected     []string
		expectedFail bool
	}{
		{actual: "foo", expected: []string{"foo"}},
		{actual: "foo", expected: []string{"foo", "foo"}},
		{actual: "foo", expected: []string{"foo", "bar"}},
		{actual: "foo", expected: []string{}, expectedFail: true},
		{actual: "foo", expected: []string(nil), expectedFail: true},
		{actual: "foo", expected: []string{"bar"}, expectedFail: true},
	}
	messageFormat := "expected string to be in %#v, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsIn(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringIsNotIn(t *testing.T) {
	tests := []struct {
		actual       string
		expected     []string
		expectedFail bool
	}{
		{actual: "foo", expected: []string{"bar"}},
		{actual: "foo", expected: []string{}},
		{actual: "foo", expected: []string(nil)},
		{actual: "foo", expected: []string{"foo"}, expectedFail: true},
		{actual: "foo", expected: []string{"foo", "bar"}, expectedFail: true},
	}
	messageFormat := "expected string not to be in %#v, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsNotIn(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringStartsWith(t *testing.T) {
	tests := []struct {
		actual       string
		prefix       string
		expectedFail bool
	}{
		{actual: "foobar", prefix: "foo"},
		{actual: "foo", prefix: "foo"},
		{actual: "foo", prefix: ""},
		{actual: "bar", prefix: "foo", expectedFail: true},
		{actual: "", prefix: "f", expectedFail: true},
	}
	messageFormat := "expected string to start with %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).StartsWith(test.prefix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.prefix, test.actual))
		}
	}
}

func TestStringEndsWith(t *testing.T) {
	tests := []struct {
		actual       string
		suffix       string
		expectedFail bool
	}{
		{actual: "foobar", suffix: "bar"},
		{actual: "foo", suffix: "foo"},
		{actual: "foo", suffix: ""},
		{actual: "bar", suffix: "foo", expectedFail: true},
		{actual: "", suffix: "f", expectedFail: true},
	}
	messageFormat := "expected string to end with %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).EndsWith(test.suffix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.suffix, test.actual))
		}
	}
}

func TestStringHasLength(t *testing.T) {
	tests := []struct {
		actual       string
		length       int
		expectedFail bool
	}{
		{actual: "foo", length: 3},
		{actual: "", length: 0},
		{actual: "foo", length: 2, expectedFail: true},
	}
	messageFormat := "expected string to have length of %d, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).HasLength(test.length)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.length, test.actual))
		}
	}
}

func TestStringHasLengthLessThan(t *testing.T) {
	tests := []struct {
		actual       string
		length       int
		expectedFail bool
	}{
		{actual: "foo", length: 4},
		{actual: "", length: 1},
		{actual: "foo", length: 2, expectedFail: true},
	}
	messageFormat := "expected string to have length less than %d, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).HasLengthLessThan(test.length)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.length, test.actual))
		}
	}
}

func TestStringHasLengthGreaterThan(t *testing.T) {
	tests := []struct {
		actual       string
		length       int
		expectedFail bool
	}{
		{actual: "foo", length: 2},
		{actual: "", length: 0, expectedFail: true},
		{actual: "foo", length: 3, expectedFail: true},
	}
	messageFormat := "expected string to have length greater than %d, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).HasLengthGreaterThan(test.length)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.length, test.actual))
		}
	}
}
