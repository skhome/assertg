package assert

type BaseAssert[T any] struct {
	t    TestingT
	info *WritableAssertionInfo
	a    *T
}

// NewBaseAssert creates and returns a new BaseAssert for constructing derived assertions.
func NewBaseAssert[T any](t TestingT, info *WritableAssertionInfo, assertion *T) *BaseAssert[T] {
	return &BaseAssert[T]{t, info, assertion}
}

// DescribedAs sets an optional description for the following assertion.
func (a *BaseAssert[T]) DescribedAs(description Description) *T {
	a.info.WithDescription(description)
	return a.a
}

// WithFailMessage overrides the default error message for the following assertions.
func (a *BaseAssert[T]) WithFailMessage(message string, args ...any) *T {
	a.info.WithOverridingFailureMessage(message, args...)
	return a.a
}

// WithFailMessageSupplier overrides the default error message for the following assertions.
// The new error message is built if the assertion fails by consuming the given Supplier function.
func (a *BaseAssert[T]) WithFailMessageSupplier(supplier Supplier[string]) *T {
	a.info.WithOverridingFailureMessageSupplier(supplier)
	return a.a
}

// WithRepresentation uses the given representation to describe values in error messages.
//
//	currencyRepresentation := func(value any) string { return fmt.Sprintf("€%d", value) }
//
//	// assertion will fail with message:
//	// expected value to be zero, but got €42
//	assert.ThatInt(t, 42).
//	       WithRepresentation(currencyRepresentation).
//	       IsZero()
func (a *BaseAssert[T]) WithRepresentation(representation Representation) *T {
	a.info.UsingRepresentation(representation)
	return a.a
}

// InHexadecimal uses hexadecimal representation to describe values in error messages.
func (a *BaseAssert[T]) InHexadecimal() *T {
	a.info.UsingHexadecimalRepresentation()
	return a.a
}

// InBinary uses binary representation to describe values in error messages.
func (a *BaseAssert[T]) InBinary() *T {
	a.info.UsingBinaryRepresentation()
	return a.a
}

// FailWithMessage records an assertion error
func (a *BaseAssert[T]) FailWithMessage(message string, args ...any) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	description := a.info.Description()
	overridingErrorMessage := a.info.OverridingFailureMessage()
	representation := a.info.Representation()
	if overridingErrorMessage != "" {
		a.t.Errorf(messageFormatter.Format(description, representation, overridingErrorMessage))
	} else {
		a.t.Errorf(messageFormatter.Format(description, representation, message, args...))
	}
}
