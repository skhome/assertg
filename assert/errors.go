package assert

// tHelper is a helper interface to signal that this function is a test helper.
type tHelper interface {
	Helper()
}

// TestingT is an interface wrapper for *testing.T
type TestingT interface {
	Errorf(format string, args ...any)
}
