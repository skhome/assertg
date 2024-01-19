package assert

import (
	"github.com/skhome/assertg/check"
	"golang.org/x/exp/constraints"
)

// IntegerAssert provides asseetions on integer values.
type IntegerAssert[T constraints.Integer] struct {
	*BaseAssert[IntegerAssert[T]]
	actual T
}

// newIntegerAssert creates and returns a new IntegerAssert.
func newIntegerAssert[T constraints.Integer](t TestingT, actual T) *IntegerAssert[T] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	integerAssert := &IntegerAssert[T]{actual: actual}
	baseAssert := NewBaseAssert(t, NewWritableAssertionInfo(), integerAssert)
	integerAssert.BaseAssert = baseAssert
	return integerAssert
}

// IsEqualTo verifies that the actual value is equal to the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsEqualTo(1)
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsEqualTo(2)
func (a *IntegerAssert[T]) IsEqualTo(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegersAreEqual(a.actual, value) {
		a.FailWithMessage("expected value to equal %s, but got %s", value, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual value is not equal to the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsNotEqualTo(-1)
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsNotEqualTo(1)
func (a *IntegerAssert[T]) IsNotEqualTo(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.IntegersAreEqual(a.actual, value) {
		a.FailWithMessage("expected value not to equal %s, but got %s", value, a.actual)
	}
	return a
}

// IsZero verifies that the actual value is zero.
//
//	// assertion will pass
//	assert.ThatInteger(t, 0).IsZero()
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsZero()
func (a *IntegerAssert[T]) IsZero() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegersAreEqual(a.actual, 0) {
		a.FailWithMessage("expected value to be zero, but got %s", a.actual)
	}
	return a
}

// IsNonZero verifies that the actual value is non-zero.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsNonZero()
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsNonZero()
func (a *IntegerAssert[T]) IsNonZero() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.IntegersAreEqual(a.actual, 0) {
		a.FailWithMessage("expected value not to be zero, but got %s", a.actual)
	}
	return a
}

// IsPositive verifies that the actual value is positive.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsPositive()
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsPositive()
func (a *IntegerAssert[T]) IsPositive() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsGreaterThan(a.actual, 0) {
		a.FailWithMessage("expected value to be positive, but got %s", a.actual)
	}
	return a
}

// IsNegative verifies that the actual value is negative.
//
//	// assertion will pass
//	assert.ThatInteger(t, -1).IsNegative()
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsNegative()
func (a *IntegerAssert[T]) IsNegative() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsLessThan(a.actual, 0) {
		a.FailWithMessage("expected value to be negative, but got %s", a.actual)
	}
	return a
}

// IsNonPositive verifies that the actual value is non positive.
//
//	// assertion will pass
//	assert.ThatInteger(t, 0).IsNonPositive()
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsNonPositive()
func (a *IntegerAssert[T]) IsNonPostive() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsLessThanOrEqualTo(a.actual, 0) {
		a.FailWithMessage("expected value not to be positive, but got %s", a.actual)
	}
	return a
}

// IsNonNegative verifies that the actual value is non negative.
//
//	// assertion will pass
//	assert.ThatInteger(t, 0).IsNonNegative()
//
//	// assertion will fail
//	assert.ThatInteger(t, -1).IsNonNegative()
func (a *IntegerAssert[T]) IsNonNegative() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsGreaterThanOrEqualTo(a.actual, 0) {
		a.FailWithMessage("expected value not to be negative, but got %s", a.actual)
	}
	return a
}

// IsLessThan verifies that the actual value is less then the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 0).IsLessThan(1)
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsLessThan(0)
func (a *IntegerAssert[T]) IsLessThan(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsLessThan(a.actual, value) {
		a.FailWithMessage("expected value to be less than %s, but got %s", value, a.actual)
	}
	return a
}

// IsLessThanOrEqualTo verifies that the actual value is less than or equal to the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 0).
//	       IsLessThanOrEqualTo(1).
//	       IsLessThanOrEqualTo(0)
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsLessThanOrEqualTo(0)
func (a *IntegerAssert[T]) IsLessThanOrEqualTo(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsLessThanOrEqualTo(a.actual, value) {
		a.FailWithMessage("expected value to be less than or equal to %s, but got %s", value, a.actual)
	}
	return a
}

// IsGreaterThan verifies that the actual value is greater then the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsGreaterThan(0)
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsGreaterThan(0)
func (a *IntegerAssert[T]) IsGreaterThan(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsGreaterThan(a.actual, value) {
		a.FailWithMessage("expected value to be greater than %s, but got %s", value, a.actual)
	}
	return a
}

// IsGreaterThanOrEqualTo verifies that the actual value is greater than or equal to the given one.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).
//	       IsGreaterThanOrEqualTo(1).
//	       IsGreaterThanOrEqualTo(0)
//
//	// assertion will fail
//	assert.ThatInteger(t, 0).IsGreaterThanOrEqualTo(1)
func (a *IntegerAssert[T]) IsGreaterThanOrEqualTo(value T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsGreaterThanOrEqualTo(a.actual, value) {
		a.FailWithMessage("expected value to be greater than or equal to %s, but got %s", value, a.actual)
	}
	return a
}

// IsBetween verifies that the actual value is between the start and end value (inclusive).
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).
//	       IsBetween(0, 2).
//	       IsBetween(1, 2).
//	       IsBetween(0, 1)
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).
//	       IsBetween(2, 3)
func (a *IntegerAssert[T]) IsBetween(startInclusive T, endInclusive T) *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsBetween(a.actual, startInclusive, endInclusive) {
		a.FailWithMessage("expected value to be between %s and %s, but got %s", startInclusive, endInclusive, a.actual)
	}
	return a
}

// IsEven verifies that the actual value is even.
//
//	// assertion will pass
//	assert.ThatInteger(t, 2).IsEven()
//	assert.ThatInteger(t, 0).IsEven()
//
//	// assertion will fail
//	assert.ThatInteger(t, 1).IsEven()
func (a *IntegerAssert[T]) IsEven() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsEven(a.actual) {
		a.FailWithMessage("expected value to be even, but got %s", a.actual)
	}
	return a
}

// IsOdd verifies that the actual value is odd.
//
//	// assertion will pass
//	assert.ThatInteger(t, 1).IsOdd()
//
//	// assertion will fail
//	assert.ThatInteger(t, 2).IsOdd()
//	assert.ThatInteger(t, 0).IsOdd()
func (a *IntegerAssert[T]) IsOdd() *IntegerAssert[T] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.IntegerIsOdd(a.actual) {
		a.FailWithMessage("expected value to be odd, but got %s", a.actual)
	}
	return a
}
