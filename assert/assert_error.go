package assert

import (
	"errors"

	"github.com/skhome/assertg/check"
)

// ErrprAssert provides assertions on error values.
type ErrorAssert struct {
	*BaseAssert[ErrorAssert]
	actual error
}

// newErrorAssert creates and returns a new ErrorAssert.
func newErrorAssert(t TestingT, actual error) *ErrorAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	errorAssert := &ErrorAssert{actual: actual}
	baseAssert := NewBaseAssert(t, NewWritableAssertionInfo(), errorAssert)
	errorAssert.BaseAssert = baseAssert
	return errorAssert
}

// IsNil verifies that the actual error is nil.
//
//	// assertions will pass
//	assert.ThatError(t, nil).IsNil()
//
//	// assertions will fail
//	assert.ThatError(t, errors.New("message")).IsNil()
func (a *ErrorAssert) IsNil() *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual != nil {
		a.FailWithMessage("expected error to be nil, but got %s", a.actual)
	}
	return a
}

// IsNil verifies that the actual error is nil.
//
//	// assertions will pass
//	assert.ThatError(t, nil).IsNil()
//
//	// assertions will fail
//	assert.ThatError(t, errors.New("message")).IsNil()
func (a *ErrorAssert) IsNotNil() *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual == nil {
		a.FailWithMessage("expected error not to be nil, but got %s", a.actual)
	}
	return a
}

// Is verifies that at least one error in the actual error's chain matches the given error.
//
//	 cause := errors.New("cause")
//
//	// assertions will pass
//	assert.ThatError(t, cause).Is(cause)
//	assert.ThatError(t, fmt.Errorf("failed because of: %w", cause)).Is(cause)
//
//	// assertions will fail
//	assert.ThatError(t, errors.New("message")).Is(cause)
func (a *ErrorAssert) Is(err error) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !errors.Is(a.actual, err) {
		a.FailWithMessage("expected error to have %s in its error chain, but got %s", err, a.actual)
	}
	return a
}

// IsNot verifies that the actual error does not have the given error in its error chain.
//
//	 cause := errors.New("cause")
//
//	// assertions will fail
//	assert.ThatError(t, errors.New("message")).IsNot(cause)
//
//	// assertions will pass
//	assert.ThatError(t, cause).IsNot(cause)
//	assert.ThatError(t, fmt.Errorf("failed because of: %w", cause)).IsNot(cause)
func (a *ErrorAssert) IsNot(err error) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if errors.Is(a.actual, err) {
		a.FailWithMessage("expected error not to have %s in its error chain, but got %s", err, a.actual)
	}
	return a
}

// HasMessage verifies that the actual error has the given error message.
//
//	 err := errors.New("no such file")
//
//	// assertions will pass
//	assert.ThatError(t, err).HasMessage("no such file")
//
//	// assertions will fail
//	assert.ThatError(t, err).HasMessage("file")
func (a *ErrorAssert) HasMessage(message string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringIsEqual(a.actual.Error(), message) {
		a.FailWithMessage("expected error to have message %s, but got %s", message, a.actual.Error())
	}
	return a
}

// DoesNotHaveMessage verifies that the actual error does not have the given error message.
//
//	 err := errors.New("no such file")
//
//	// assertions will pass
//	assert.ThatError(t, err).DoesNotHaveMessage("file")
//
//	// assertions will fail
//	assert.ThatError(t, err).DoesNotHaveMessage("no such file")
func (a *ErrorAssert) DoesNotHaveMessage(message string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringIsEqual(a.actual.Error(), message) {
		a.FailWithMessage("expected error not to have message %s, but got %s", message, a.actual.Error())
	}
	return a
}

// HasMessageContaining verifies that the actual error message contains all given values.
//
//	 err := errors.New("invalid users frodo, merry, pippin")
//
//	// assertions will pass
//	assert.ThatError(t, err).
//	       HasMessageContaining("frodo").
//	       HasMessageContaining("frodo", "merry")
//
//	// assertions will fail
//	assert.ThatError(t, err).
//	       HasMessageContaining("sam").
//	       HasMessageContaining("frodo", "sam")
func (a *ErrorAssert) HasMessageContaining(values ...string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContains(a.actual.Error(), values) {
		a.FailWithMessage("expected error to have message containing %s, but got %s", values, a.actual.Error())
	}
	return a
}

// HasMessageContainingAnyOf verifies that the actual error message contains any of the given values.
//
//	 err := errors.New("invalid users frodo, merry, pippin")
//
//	// assertions will pass
//	assert.ThatError(t, err).
//	       HasMessageContaining("frodo", "merry").
//	       HasMessageContaining("frodo", "sam")
//
//	// assertions will fail
//	assert.ThatError(t, err).
//	       HasMessageContaining("sam").
func (a *ErrorAssert) HasMessageContainingAnyOf(values ...string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsAny(a.actual.Error(), values) {
		a.FailWithMessage("expected error to have message containing any of %s, but got %s", values, a.actual.Error())
	}
	return a
}

// HasMessageNotContaining verifies that the actual error message does not contain the given content.
//
//	 err := errors.New("wrong amount 123")
//
//	// assertions will pass
//	assert.ThatError(t, err).
//	       HasMessageNotContaining("234")
//
//	// assertions will fail
//	assert.ThatError(t, err).
//	       HasMessageNotContaining("123")
func (a *ErrorAssert) HasMessageNotContaining(content string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringContains(a.actual.Error(), []string{content}) {
		a.FailWithMessage("expected error not to have message containing %s, but got %s", content, a.actual.Error())
	}
	return a
}

// HasMessageStartingWith verifies that the actual error message starts with the given prefix.
//
//	 err := errors.New("wrong amount 123")
//
//	// assertions will pass
//	assert.ThatError(t, err).
//	       HasMessageStartingWith("wrong amount")
//
//	// assertions will fail
//	assert.ThatError(t, err).
//	       HasMessageStartingWith("right amount")
func (a *ErrorAssert) HasMessageStartingWith(prefix string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringStartsWith(a.actual.Error(), prefix) {
		a.FailWithMessage("expected error to have message starting with %s, but got %s", prefix, a.actual.Error())
	}
	return a
}

// HasMessageEndingWith verifies that the actual error message ends with the given suffix.
//
//	 err := errors.New("wrong amount 123")
//
//	// assertions will pass
//	assert.ThatError(t, err).
//	       HasMessageEndingWith("wrong amount")
//
//	// assertions will fail
//	assert.ThatError(t, err).
//	       HasMessageEndingWith("right amount")
func (a *ErrorAssert) HasMessageEndingWith(suffix string) *ErrorAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringEndsWith(a.actual.Error(), suffix) {
		a.FailWithMessage("expected error to have message ending with %s, but got %s", suffix, a.actual.Error())
	}
	return a
}
