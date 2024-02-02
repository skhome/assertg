package assert

import (
	"github.com/skhome/assertg/check"
	"golang.org/x/exp/constraints"
)

// FloatAssert provides asseetions on float values.
type FloatAssert[T constraints.Float] struct {
	*BaseAssert[FloatAssert[T]]
	actual T
}

// newFloatAssert creates and returns a new FloatAssert.
func newFloatAssert[T constraints.Float](t TestingT, actual T) *FloatAssert[T] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	integerAssert := &FloatAssert[T]{actual: actual}
	baseAssert := NewBaseAssert(t, NewWritableAssertionInfo(), integerAssert)
	integerAssert.BaseAssert = baseAssert
	return integerAssert
}

// IsEqualTo verifies that the actual value is equal to the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).IsEqualTo(1)
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).IsEqualTo(2)
func (a *FloatAssert[T]) IsEqualTo(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatsAreEqual(a.actual, value) {
		a.FailWithMessage("expected value to equal %s, but got %s", value, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual value is not equal to the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).IsNotEqualTo(-1)
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).IsNotEqualTo(1)
func (a *FloatAssert[T]) IsNotEqualTo(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.FloatsAreEqual(a.actual, value) {
		a.FailWithMessage("expected value not to equal %s, but got %s", value, a.actual)
	}
	return a
}

// IsZero verifies that the actual value is zero.
//
//	// assertion will pass
//	assert.ThatFloat(t, 0).IsZero()
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).IsZero()
func (a *FloatAssert[T]) IsZero() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatsAreEqual(a.actual, 0) {
		a.FailWithMessage("expected value to be zero, but got %s", a.actual)
	}
	return a
}

// IsNonZero verifies that the actual value is non-zero.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).IsNonZero()
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsNonZero()
func (a *FloatAssert[T]) IsNonZero() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.FloatsAreEqual(a.actual, 0) {
		a.FailWithMessage("expected value not to be zero, but got %s", a.actual)
	}
	return a
}

// IsPositive verifies that the actual value is positive.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).IsPositive()
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsPositive()
func (a *FloatAssert[T]) IsPositive() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsGreaterThan(a.actual, 0) {
		a.FailWithMessage("expected value to be positive, but got %s", a.actual)
	}
	return a
}

// IsNegative verifies that the actual value is negative.
//
//	// assertion will pass
//	assert.ThatFloat(t, -1).IsNegative()
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsNegative()
func (a *FloatAssert[T]) IsNegative() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsLessThan(a.actual, 0) {
		a.FailWithMessage("expected value to be negative, but got %s", a.actual)
	}
	return a
}

// IsNonPositive verifies that the actual value is non positive.
//
//	// assertion will pass
//	assert.ThatFloat(t, 0).IsNonPositive()
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).IsNonPositive()
func (a *FloatAssert[T]) IsNonPostive() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsLessThanOrEqualTo(a.actual, 0) {
		a.FailWithMessage("expected value not to be positive, but got %s", a.actual)
	}
	return a
}

// IsNonNegative verifies that the actual value is non negative.
//
//	// assertion will pass
//	assert.ThatFloat(t, 0).IsNonNegative()
//
//	// assertion will fail
//	assert.ThatFloat(t, -1).IsNonNegative()
func (a *FloatAssert[T]) IsNonNegative() *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsGreaterThanOrEqualTo(a.actual, 0) {
		a.FailWithMessage("expected value not to be negative, but got %s", a.actual)
	}
	return a
}

// IsLessThan verifies that the actual value is less then the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 0).IsLessThan(1)
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsLessThan(0)
func (a *FloatAssert[T]) IsLessThan(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsLessThan(a.actual, value) {
		a.FailWithMessage("expected value to be less than %s, but got %s", value, a.actual)
	}
	return a
}

// IsLessThanOrEqualTo verifies that the actual value is less than or equal to the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 0).
//	       IsLessThanOrEqualTo(1).
//	       IsLessThanOrEqualTo(0)
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).IsLessThanOrEqualTo(0)
func (a *FloatAssert[T]) IsLessThanOrEqualTo(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsLessThanOrEqualTo(a.actual, value) {
		a.FailWithMessage("expected value to be less than or equal to %s, but got %s", value, a.actual)
	}
	return a
}

// IsGreaterThan verifies that the actual value is greater then the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).IsGreaterThan(0)
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsGreaterThan(0)
func (a *FloatAssert[T]) IsGreaterThan(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsGreaterThan(a.actual, value) {
		a.FailWithMessage("expected value to be greater than %s, but got %s", value, a.actual)
	}
	return a
}

// IsGreaterThanOrEqualTo verifies that the actual value is greater than or equal to the given one.
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).
//	       IsGreaterThanOrEqualTo(1).
//	       IsGreaterThanOrEqualTo(0)
//
//	// assertion will fail
//	assert.ThatFloat(t, 0).IsGreaterThanOrEqualTo(1)
func (a *FloatAssert[T]) IsGreaterThanOrEqualTo(value T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsGreaterThanOrEqualTo(a.actual, value) {
		a.FailWithMessage("expected value to be greater than or equal to %s, but got %s", value, a.actual)
	}
	return a
}

// IsBetween verifies that the actual value is between the start and end value (inclusive).
//
//	// assertion will pass
//	assert.ThatFloat(t, 1).
//	       IsBetween(0, 2).
//	       IsBetween(1, 2).
//	       IsBetween(0, 1)
//
//	// assertion will fail
//	assert.ThatFloat(t, 1).
//	       IsBetween(2, 3)
func (a *FloatAssert[T]) IsBetween(startInclusive T, endInclusive T) *FloatAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.FloatIsBetween(a.actual, startInclusive, endInclusive) {
		a.FailWithMessage("expected value to be between %s and %s, but got %s", startInclusive, endInclusive, a.actual)
	}
	return a
}
