package assert

// ThatString starts assertions on a string.
func ThatString(t TestingT, actual string) *StringAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	stringAssert := &StringAssert{actual: actual}
	baseAssert := &BaseAssert[StringAssert]{t: t, info: NewWritableAssertionInfo(), a: stringAssert}
	stringAssert.BaseAssert = baseAssert
	return stringAssert
}

// ThatSlice starts assertions on a slice.
func ThatSlice[T ~[]E, E any](t TestingT, actual T) *SliceAssert[E] {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	sliceAssert := &SliceAssert[E]{actual: actual}
	baseAssert := &BaseAssert[SliceAssert[E]]{t: t, info: NewWritableAssertionInfo(), a: sliceAssert}
	sliceAssert.BaseAssert = baseAssert
	return sliceAssert
}

// ThatBool starts assertions on a bool.
func ThatBool(t TestingT, actual bool) *BoolAssert {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	boolAssert := &BoolAssert{actual: actual}
	baseAssert := &BaseAssert[BoolAssert]{t: t, info: NewWritableAssertionInfo(), a: boolAssert}
	boolAssert.BaseAssert = baseAssert
	return boolAssert
}
