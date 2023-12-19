package assert

import (
	"fmt"
	"slices"
	"strings"
)

// StringAssert provides assertions on strings.
type StringAssert struct {
	t           TestingT
	message     string
	description string
	actual      string
}

// As sets an optional description for this assertion.
func (a *StringAssert) As(format string, args ...any) *StringAssert {
	a.description = fmt.Sprintf(format, args...)
	return a
}

// WithFailMessage overrides the failure message.
func (a *StringAssert) WithFailMessage(format string, args ...any) *StringAssert {
	a.message = fmt.Sprintf(format, args...)
	return a
}

// failWithMessage records an assertion failure.
func (a *StringAssert) failWithMessage(format string, args ...any) {
	message := a.message
	if len(message) == 0 {
		message = fmt.Sprintf(format, args...)
	}
	Fail(a.t, message, a.description)
}

// IsEmpty verifies that the actual value is empty (has a length of 0).
func (a *StringAssert) IsEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) > 0 {
		a.failWithMessage("expected string to be empty, but got %q", a.actual)
	}
	return a
}

// IsBlank verifies that the actual value is blank (empty or only whitespace characters).
func (a *StringAssert) IsBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(strings.TrimSpace(a.actual)) > 0 {
		a.failWithMessage("expected string to be blank, but got %q", a.actual)
	}
	return a
}

// IsEqualTo verifies that the actual value equals the given one.
func (a *StringAssert) IsEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected != a.actual {
		a.failWithMessage("expected string to equal %q, but got %q", expected, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual value does not equal the given one.
func (a *StringAssert) IsNotEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected == a.actual {
		a.failWithMessage("expected string to not equal %q, but got %q", expected, a.actual)
	}
	return a
}

// Matches verifies that the actual value matches the given predicate.
func (a *StringAssert) Matches(predicate Predicate[string]) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !predicate(a.actual) {
		a.failWithMessage("expected string to match predicate, but got %q", a.actual)
	}
	return a
}

// DoesNotMatch verifies that the actual value does not match the given predicate.
func (a *StringAssert) DoesNotMatch(predicate Predicate[string]) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if predicate(a.actual) {
		a.failWithMessage("expected string to not match predicate, but got %q", a.actual)
	}
	return a
}

// IsIn verifies that the actual value is in the given slice.
func (a *StringAssert) IsIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !slices.Contains(slice, a.actual) {
		a.failWithMessage("expected string to be in %#v, but got %q", slice, a.actual)
	}
	return a
}

// IsNotIn verifies that the actual value is not in the given slice.
func (a *StringAssert) IsNotIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if slices.Contains(slice, a.actual) {
		a.failWithMessage("expected string not to be in %#v, but got %q", slice, a.actual)
	}
	return a
}

// StartsWith verifies that the actual value starts with the given prefix.
func (a *StringAssert) StartsWith(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasPrefix(a.actual, prefix) {
		a.failWithMessage("expected string to start with %q, but got %q", prefix, a.actual)
	}
	return a
}

// EndsWith verifies that the actual value ends with the given suffix.
func (a *StringAssert) EndsWith(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasSuffix(a.actual, suffix) {
		a.failWithMessage("expected string to end with %q, but got %q", suffix, a.actual)
	}
	return a
}

// HasLength verifies that the actual value has the expected length.
func (a *StringAssert) HasLength(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != length {
		a.failWithMessage("expected string to have length of %d, but got %q", length, a.actual)
	}
	return a
}

// HasLengthLessThan verifies that the actual value has a length less than expected.
func (a *StringAssert) HasLengthLessThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) >= length {
		a.failWithMessage("expected string to have length less than %d, but got %q", length, a.actual)
	}
	return a
}

// HasLengthGreaterThan verifies that the actual value has a length greater than expected.
func (a *StringAssert) HasLengthGreaterThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) <= length {
		a.failWithMessage("expected string to have length greater than %d, but got %q", length, a.actual)
	}
	return a
}
