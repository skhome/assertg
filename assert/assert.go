package assert

// TestingT is an interface wrapper for *testing.T
type TestingT interface {
	Errorf(format string, args ...any)
}

// Predicate is a function that returns if a value meets a condition.
type Predicate[T any] func(value T) bool

// Condition is a function to assert a condition.
type Condition[T any] func(value T)

// ThatString starts assertions on a string.
func ThatString(t TestingT, actual string) *StringAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return &StringAssert{t: t, actual: actual}
}

// ThatSlice starts assertions on a slice.
func ThatSlice[T ~[]E, E any](t TestingT, actual T) *SliceAssert[T, E] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return &SliceAssert[T, E]{t: t, actual: actual}
}
