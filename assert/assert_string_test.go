package assert

import (
	"fmt"
	"regexp"
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
		{actual: "Frodo", expectedFail: true},
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

func TestStringIsNotEmpty(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: "Frodo"},
		{actual: " "},
		{actual: "", expectedFail: true},
	}
	messageFormat := "expected string to not be empty, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsNotEmpty()

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

func TestStringIsNotBlank(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: "Frodo"},
		{actual: "", expectedFail: true},
		{actual: " ", expectedFail: true},
		{actual: "\t", expectedFail: true},
	}
	messageFormat := "expected string to not be blank, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsNotBlank()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringContainsWhitespace(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: " "},
		{actual: "Frodo Baggins"},
		{actual: "", expectedFail: true},
		{actual: "Frodo", expectedFail: true},
	}
	messageFormat := "expected string to contain whitespace characters, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsWhitespace()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringDoesNotContainWhitespace(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: ""},
		{actual: "Frodo"},
		{actual: " ", expectedFail: true},
		{actual: "Frodo Baggins", expectedFail: true},
	}
	messageFormat := "expected string to not contain whitespace characters, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotContainWhitespace()

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

func TestStringIsEqualToIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		expected     string
		expectedFail bool
	}{
		{actual: "frodo", expected: "Frodo"},
		{actual: "Frodo", expected: "frodo"},
		{actual: "frodo", expected: "gandalf", expectedFail: true},
	}
	messageFormat := "expected string to equal %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsEqualToIgnoringCase(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringIsNotEqualToIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		expected     string
		expectedFail bool
	}{
		{actual: "Gandalf", expected: "Hobbit"},
		{actual: "Gandalf", expected: "gandalf", expectedFail: true},
	}
	messageFormat := "expected string to not equal %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsNotEqualToIgnoringCase(test.expected)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.expected, test.actual))
		}
	}
}

func TestStringContainsOnlyDigits(t *testing.T) {
	tests := []struct {
		actual       string
		expectedFail bool
	}{
		{actual: "10"},
		{actual: "10€", expectedFail: true},
	}
	messageFormat := "expected string to only contain digits, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsOnlyDigits()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.actual))
		}
	}
}

func TestStringContainsOnlyOnce(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Frodo", substr: "do"},
		{actual: "Frodo", substr: "o", expectedFail: true},
		{actual: "Frodo", substr: "y", expectedFail: true},
	}
	messageFormat := "expected string to contain %q only once, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsOnlyOnce(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringIsIn(t *testing.T) {
	tests := []struct {
		actual       string
		expected     []string
		expectedFail bool
	}{
		{actual: "nenya", expected: []string{"nenya"}},
		{actual: "nenya", expected: []string{"nenya", "nenya"}},
		{actual: "nenya", expected: []string{"vilya", "nenya", "varya"}},
		{actual: "one", expected: []string{}, expectedFail: true},
		{actual: "one", expected: nil, expectedFail: true},
		{actual: "one", expected: []string{"vilya", "nenya", "varya"}, expectedFail: true},
	}
	messageFormat := "expected string to be present in %#v, but got %q"

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
		{actual: "one", expected: []string{"vilya", "nenya", "varya"}},
		{actual: "one", expected: []string{}},
		{actual: "one", expected: nil},
		{actual: "nenya", expected: []string{"nenya"}, expectedFail: true},
		{actual: "nenya", expected: []string{"vilya", "nenya", "varya"}, expectedFail: true},
	}
	messageFormat := "expected string not to be present in %#v, but got %q"

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
		{actual: "Frodo", prefix: "Fro"},
		{actual: "Frodo", prefix: "Frodo"},
		{actual: "Frodo", prefix: ""},
		{actual: "Frodo", prefix: "fro", expectedFail: true},
		{actual: "", prefix: "fro", expectedFail: true},
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

func TestStringDoesNotStartWith(t *testing.T) {
	tests := []struct {
		actual       string
		prefix       string
		expectedFail bool
	}{
		{actual: "Frodo", prefix: "fro"},
		{actual: "Frodo", prefix: "Sam"},
		{actual: "Frodo", prefix: "", expectedFail: true},
		{actual: "Frodo", prefix: "Fro", expectedFail: true},
	}
	messageFormat := "expected string to not start with %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotStartWith(test.prefix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.prefix, test.actual))
		}
	}
}

func TestStringStartsWithIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		prefix       string
		expectedFail bool
	}{
		{actual: "Frodo", prefix: "Fro"},
		{actual: "Frodo", prefix: "fro"},
		{actual: "Frodo", prefix: "gan", expectedFail: true},
	}
	messageFormat := "expected string to start with %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).StartsWithIgnoringCase(test.prefix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.prefix, test.actual))
		}
	}
}

func TestStringDoesNotStartWithIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		prefix       string
		expectedFail bool
	}{
		{actual: "Gandalf", prefix: "Fro"},
		{actual: "Gandalf", prefix: "gan", expectedFail: true},
	}
	messageFormat := "expected string to not start with %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotStartWithIgnoringCase(test.prefix)

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

func TestStringDoesNotEndWith(t *testing.T) {
	tests := []struct {
		actual       string
		suffix       string
		expectedFail bool
	}{
		{actual: "Frodo", suffix: "Fro"},
		{actual: "Frodo", suffix: "do", expectedFail: true},
	}
	messageFormat := "expected string not to end with %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotEndWith(test.suffix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.suffix, test.actual))
		}
	}
}

func TestStringEndsWithIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		suffix       string
		expectedFail bool
	}{
		{actual: "foobar", suffix: "bar"},
		{actual: "FooBar", suffix: "bar"},
		{actual: "bar", suffix: "foo", expectedFail: true},
	}
	messageFormat := "expected string to end with %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).EndsWithIgnoringCase(test.suffix)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.suffix, test.actual))
		}
	}
}

func TestStringDoesNotEndWithIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		suffix       string
		expectedFail bool
	}{
		{actual: "Frodo", suffix: "fro"},
		{actual: "Frodo", suffix: "do", expectedFail: true},
		{actual: "Frodo", suffix: "Do", expectedFail: true},
	}
	messageFormat := "expected string not to end with %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotEndWithIgnoringCase(test.suffix)

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

func TestStringHasSameLengthAs(t *testing.T) {
	tests := []struct {
		actual       string
		other        string
		expectedFail bool
	}{
		{actual: "foo", other: "bar"},
		{actual: "", other: ""},
		{actual: "foo", other: "bazy", expectedFail: true},
	}
	messageFormat := "expected string to have the same length as %q (%d), but got %q (%d)"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).HasSameLengthAs(test.other)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.other, len(test.other), test.actual, len(test.actual)))
		}
	}
}

func TestStringHasLineCount(t *testing.T) {
	tests := []struct {
		actual       string
		lineCount    int
		expectedFail bool
	}{
		{actual: "foo", lineCount: 1},
		{actual: "", lineCount: 1},
		{actual: "foo\nbar", lineCount: 2},
		{actual: "foo\nbar\n", lineCount: 3},
		{actual: "foo", lineCount: 2, expectedFail: true},
	}
	messageFormat := "expected string to have %d lines, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).HasLineCount(test.lineCount)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.lineCount, test.actual))
		}
	}
}

func TestStringContains(t *testing.T) {
	tests := []struct {
		actual       string
		value        string
		expectedFail bool
	}{
		{actual: "foobar", value: "oba"},
		{actual: "foobar", value: "foo"},
		{actual: "foobar", value: ""},
		{actual: "FooBar", value: "foo", expectedFail: true},
		{actual: "", value: "foo", expectedFail: true},
	}
	messageFormat := "expected string to contain %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).Contains(test.value)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.value, test.actual))
		}
	}
}

func TestStringContainsIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Gandalf the grey", substr: "gandalf"},
		{actual: "Gandalf the grey", substr: "white", expectedFail: true},
	}
	messageFormat := "expected string to contain %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsIgnoringCase(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringContainsIgnoringWhitespace(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Gandalf the grey", substr: "thegrey"},
		{actual: "Gandalf the grey", substr: "thegr ey"},
		{actual: "Gandalf the grey", substr: "Grey", expectedFail: true},
	}
	messageFormat := "expected string to contain %q ignoring whitespace, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsIgnoringWhitespace(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringContainAllOf(t *testing.T) {
	tests := []struct {
		actual       string
		values       []string
		expectedFail bool
	}{
		{actual: "foobar", values: []string{"foo", "bar"}},
		{actual: "FooBar", values: []string{"foo"}, expectedFail: true},
		{actual: "foobar", values: []string{"foo", "baz"}, expectedFail: true},
	}
	messageFormat := "expected string to contain all of %#v, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsAllOf(test.values...)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.values, test.actual))
		}
	}
}

func TestStringContainsAnyOf(t *testing.T) {
	tests := []struct {
		actual       string
		values       []string
		expectedFail bool
	}{
		{actual: "foobar", values: []string{"foo", "bar"}},
		{actual: "foobar", values: []string{"foo", "baz"}},
		{actual: "FooBar", values: []string{"foo"}, expectedFail: true},
	}
	messageFormat := "expected string to contain any of %#v, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).ContainsAnyOf(test.values...)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.values, test.actual))
		}
	}
}

func TestStringDoesNotContain(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Frodo", substr: "Pippin"},
		{actual: "Frodo", substr: "do", expectedFail: true},
	}
	messageFormat := "expected string to not contain %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotContain(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringDoesNotContainIgnoringCase(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Frodo", substr: "sam"},
		{actual: "Frodo", substr: "fro", expectedFail: true},
	}
	messageFormat := "expected string to not contain %q ignoring case, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotContainIgnoringCase(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringDoesNotContainIgnoringWhitespace(t *testing.T) {
	tests := []struct {
		actual       string
		substr       string
		expectedFail bool
	}{
		{actual: "Gandalf the grey", substr: "TheGrey"},
		{actual: "Gandalf the grey", substr: "thegrey", expectedFail: true},
	}
	messageFormat := "expected string to not contain %q ignoring whitespace, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotContainIgnoringWhitespace(test.substr)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.substr, test.actual))
		}
	}
}

func TestStringMatches(t *testing.T) {
	tests := []struct {
		actual       string
		pattern      string
		expectedFail bool
	}{
		{actual: "Frodo", pattern: `..o.o`},
		{actual: "Frodo", pattern: `.*f`, expectedFail: true},
	}
	messageFormat := "expected string to match %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).Matches(test.pattern)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.pattern, test.actual))
		}
	}
}

func TestStringDoesNotMatch(t *testing.T) {
	tests := []struct {
		actual       string
		pattern      string
		expectedFail bool
	}{
		{actual: "Frodo", pattern: `.*d$`},
		{actual: "Frodo", pattern: `F.*`, expectedFail: true},
	}
	messageFormat := "expected string not to match %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotMatch(test.pattern)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.pattern, test.actual))
		}
	}
}

func TestStringMatchesRegexp(t *testing.T) {
	tests := []struct {
		actual       string
		regex        *regexp.Regexp
		expectedFail bool
	}{
		{actual: "Frodo", regex: regexp.MustCompile(`..o.o`)},
		{actual: "Frodo", regex: regexp.MustCompile(`.*f`), expectedFail: true},
	}
	messageFormat := "expected string to match \"%v\", but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).MatchesRegexp(test.regex)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.regex, test.actual))
		}
	}
}

func TestStringDoesNotMatchRegexp(t *testing.T) {
	tests := []struct {
		actual       string
		regex        *regexp.Regexp
		expectedFail bool
	}{
		{actual: "Frodo", regex: regexp.MustCompile(`.*f`)},
		{actual: "Frodo", regex: regexp.MustCompile(`..o.o`), expectedFail: true},
	}
	messageFormat := "expected string to match \"%v\", but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).DoesNotMatchRegexp(test.regex)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.regex, test.actual))
		}
	}
}

func TestStringIsSubstringOf(t *testing.T) {
	tests := []struct {
		actual       string
		str          string
		expectedFail bool
	}{
		{actual: "Lego", str: "Legolas"},
		{actual: "Frodo", str: "Fro", expectedFail: true},
	}
	messageFormat := "expected string to be a substring of %q, but got %q"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatString(fixture, test.actual).IsSubstringOf(test.str)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.str, test.actual))
		}
	}
}
