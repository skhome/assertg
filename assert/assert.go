package assert

import "golang.org/x/exp/constraints"

// ThatString starts assertions on a string.
func ThatString(t TestingT, actual string) *StringAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return newStringAssert(t, actual)
}

// ThatSlice starts assertions on a slice.
func ThatSlice[T ~[]E, E any](t TestingT, actual T) *SliceAssert[E] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return newSliceAssert(t, actual)
}

// ThatBool starts assertions on a bool.
func ThatBool(t TestingT, actual bool) *BoolAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return newBoolAssert(t, actual)
}

// ThatInteger starts assertions on an integer.
func ThatInteger[T constraints.Integer](t TestingT, actual T) *IntegerAssert[T] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return newIntegerAssert(t, actual)
}

// ThatError starts assertions on an error.
func ThatError(t TestingT, actual error) *ErrorAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return newErrorAssert(t, actual)
}
