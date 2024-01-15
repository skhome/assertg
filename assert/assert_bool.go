package assert

// BoolAssert provides assertions on boolean values.
type BoolAssert struct {
	*BaseAssert[BoolAssert]
	actual bool
}

// IsTrue verifies that the actual value is true.
//
//	// assertions will pass
//	assert.ThatBool(t, true).IsTrue()
//
//	// assertions will fail
//	assert.ThatBool(t, false).IsTrue()
func (a *BoolAssert) IsTrue() *BoolAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !a.actual {
		a.FailWithMessage("expected value to be true, but got %s", a.actual)
	}
	return a
}

// IsFalse verifies that the actual value is true.
//
//	// assertions will pass
//	assert.ThatBool(t, false).IsFalse()
//
//	// assertions will fail
//	assert.ThatBool(t, true).IsFalse()
func (a *BoolAssert) IsFalse() *BoolAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual {
		a.FailWithMessage("expected value to be false, but got %s", a.actual)
	}
	return a
}

// IsEqualTo verifies that the actual value is equal to the given one.
//
//	// assertions will pass
//	assert.ThatBool(t, false).IsEqualTo(false)
//
//	// assertions will fail
//	assert.ThatBool(t, true).IsEqualTo(false)
func (a *BoolAssert) IsEqualTo(expected bool) *BoolAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual {
		a.FailWithMessage("expected value to be %s, but got %s", expected, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual value is not equal to the given one.
//
//	// assertions will pass
//	assert.ThatBool(t, false).IsNotEqualTo(true)
//
//	// assertions will fail
//	assert.ThatBool(t, true).IsNotEqualTo(true)
func (a *BoolAssert) IsNotEqualTo(expected bool) *BoolAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual {
		a.FailWithMessage("expected value not to be %s, but got %s", expected, a.actual)
	}
	return a
}
