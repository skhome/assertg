package assert

import (
	"fmt"
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

// IsEmpty checks that the actual string is empty (has a length of 0)
func (a *StringAssert) IsEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) > 0 {
		if len(a.message) == 0 {
			a.message = fmt.Sprintf("expected string to be empty, but got %q", a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// IsBlank checks that the actual string is empty or contains only whitespace characters.
func (a *StringAssert) IsBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(strings.TrimSpace(a.actual)) > 0 {
		if len(a.message) == 0 {
			a.message = fmt.Sprintf("expected string to be blank, but got %q", a.actual)
		}
	}
	return a
}
