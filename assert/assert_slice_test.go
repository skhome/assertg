package assert

import (
	"fmt"
	"testing"
)

func TestSliceIsNil(t *testing.T) {
	tests := []struct {
		slice        []string
		expectedFail bool
		message      string
	}{
		{
			slice:        nil,
			expectedFail: false,
		},
		{
			slice:        []string(nil),
			expectedFail: false,
		},
		{
			slice:        []string{"foo"},
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to be nil, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).IsNil()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceIsNotNil(t *testing.T) {
	tests := []struct {
		slice        []string
		expectedFail bool
		message      string
	}{
		{
			slice:        []string{"foo"},
			expectedFail: false,
		},
		{
			slice:        nil,
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to not be nil, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).IsNotNil()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceHasSize(t *testing.T) {
	tests := []struct {
		slice        []string
		size         int
		expectedFail bool
	}{
		{
			slice:        []string{"foo"},
			size:         1,
			expectedFail: false,
		},
		{
			slice:        []string(nil),
			size:         0,
			expectedFail: false,
		},
		{
			slice:        []string(nil),
			size:         1,
			expectedFail: true,
		},
		{
			slice:        []string{"foo"},
			size:         2,
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to have a size of %d, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasSize(test.size)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.size, test.slice))
		}
	}
}

func TestSliceHasSizeGreaterThan(t *testing.T) {
	tests := []struct {
		slice        []string
		threshold    int
		expectedFail bool
	}{
		{
			slice:        []string{"foo"},
			threshold:    0,
			expectedFail: false,
		},
		{
			slice:        []string{"foo", "bar"},
			threshold:    1,
			expectedFail: false,
		},
		{
			slice:        nil,
			threshold:    0,
			expectedFail: true,
		},
		{
			slice:        []string(nil),
			threshold:    0,
			expectedFail: true,
		},
		{
			slice:        []string{"foo"},
			threshold:    1,
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to have a size greater than %d, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasSizeGreaterThan(test.threshold)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.threshold, test.slice))
		}
	}
}

func TestSliceHasSizeLessThan(t *testing.T) {
	tests := []struct {
		slice        []string
		threshold    int
		expectedFail bool
	}{
		{
			slice:        []string{"foo"},
			threshold:    2,
			expectedFail: false,
		},
		{
			slice:        []string{"foo", "bar"},
			threshold:    3,
			expectedFail: false,
		},
		{
			slice:        nil,
			threshold:    1,
			expectedFail: false,
		},
		{
			slice:        []string(nil),
			threshold:    1,
			expectedFail: false,
		},
		{
			slice:        []string{"foo"},
			threshold:    1,
			expectedFail: true,
		},
		{
			slice:        nil,
			threshold:    0,
			expectedFail: true,
		},
		{
			slice:        []string(nil),
			threshold:    0,
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to have a size less than %d, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasSizeLessThan(test.threshold)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.threshold, test.slice))
		}
	}
}

func TestSliceIsEmpty(t *testing.T) {
	tests := []struct {
		slice        []string
		size         int
		expectedFail bool
	}{
		{
			slice:        nil,
			expectedFail: false,
		},
		{
			slice:        []string(nil),
			expectedFail: false,
		},
		{
			slice:        []string{},
			expectedFail: false,
		},
		{
			slice:        []string{"foo"},
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to be empty, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).IsEmpty()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceIsNotEmpty(t *testing.T) {
	tests := []struct {
		slice        []string
		size         int
		expectedFail bool
	}{
		{
			slice:        []string{"foo"},
			expectedFail: false,
		},
		{
			slice:        nil,
			expectedFail: true,
		},
		{
			slice:        []string(nil),
			expectedFail: true,
		},
		{
			slice:        []string{},
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to not be empty, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).IsNotEmpty()

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceContains(t *testing.T) {
	type bit struct{ name string }
	tests := []struct {
		slice        []*bit
		element      *bit
		expectedFail bool
	}{
		{
			slice:        []*bit{{name: "foo"}},
			element:      &bit{name: "foo"},
			expectedFail: false,
		},
		{
			slice:        []*bit{{name: "foo"}, {name: "bar"}},
			element:      &bit{name: "foo"},
			expectedFail: false,
		},
		{
			slice:        []*bit{nil},
			element:      nil,
			expectedFail: false,
		},
		{
			slice:        []*bit{{name: "bar"}},
			element:      &bit{name: "foo"},
			expectedFail: true,
		},
		{
			slice:        []*bit(nil),
			element:      nil,
			expectedFail: true,
		},
		{
			slice:        nil,
			element:      nil,
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to contain %#v, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).Contains(test.element)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.element, test.slice))
		}
	}
}

func TestSliceDoesNotContain(t *testing.T) {
	type bit struct{ name string }
	tests := []struct {
		slice        []*bit
		element      *bit
		expectedFail bool
	}{
		{
			slice:        []*bit{{name: "bar"}},
			element:      &bit{name: "foo"},
			expectedFail: false,
		},
		{
			slice:        []*bit{{name: "foo"}, {name: "bar"}},
			element:      &bit{name: "baz"},
			expectedFail: false,
		},
		{
			slice:        []*bit{},
			element:      &bit{name: "foo"},
			expectedFail: false,
		},
		{
			slice:        nil,
			element:      nil,
			expectedFail: false,
		},
		{
			slice:        []*bit{},
			element:      nil,
			expectedFail: false,
		},
		{
			slice:        []*bit{{name: "foo"}},
			element:      &bit{name: "foo"},
			expectedFail: true,
		},
		{
			slice:        []*bit{{name: "foo"}, {name: "bar"}},
			element:      &bit{name: "foo"},
			expectedFail: true,
		},
	}
	messageFormat := "expected slice to not contain %#v, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).DoesNotContain(test.element)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.element, test.slice))
		}
	}
}

func TestSliceHasAnyMatch(t *testing.T) {
	tests := []struct {
		slice        []string
		expectedFail bool
	}{
		{slice: []string{"foo"}},
		{slice: []string{"foo", "bar"}},
		{slice: []string{"foo", "foo"}},
		{slice: nil, expectedFail: true},
		{slice: []string(nil), expectedFail: true},
		{slice: []string{}, expectedFail: true},
		{slice: []string{"bar"}, expectedFail: true},
	}
	predicate := func(element string) bool { return element == "foo" }
	messageFormat := "expected slice to have any entry match the predicate, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasAnyMatch(predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceHasAllMatch(t *testing.T) {
	tests := []struct {
		slice        []string
		expectedFail bool
	}{
		{slice: []string{"foo"}},
		{slice: []string{"foo", "foo"}},
		{slice: nil},
		{slice: []string(nil)},
		{slice: []string{}},
		{slice: []string{"bar"}, expectedFail: true},
		{slice: []string{"foo", "bar"}, expectedFail: true},
	}
	predicate := func(element string) bool { return element == "foo" }
	messageFormat := "expected slice to have all entries match the predicate, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasAllMatch(predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.slice))
		}
	}
}

func TestSliceHasAtLeastMatch(t *testing.T) {
	tests := []struct {
		slice        []string
		threshold    int
		expectedFail bool
	}{
		{slice: []string{"foo"}, threshold: 1},
		{slice: []string{"foo", "bar"}, threshold: 1},
		{slice: []string{"foo", "foo"}, threshold: 2},
		{slice: []string{"foo", "foo", "foo"}, threshold: 2},
		{slice: nil, threshold: 0},
		{slice: []string(nil), threshold: 0},
		{slice: []string{}, threshold: 0},
		{slice: []string{"bar"}, threshold: 0},
		{slice: nil, threshold: 1, expectedFail: true},
		{slice: []string(nil), threshold: 1, expectedFail: true},
		{slice: []string{}, threshold: 1, expectedFail: true},
		{slice: []string{"bar"}, threshold: 1, expectedFail: true},
	}
	predicate := func(element string) bool { return element == "foo" }
	messageFormat := "expected slice to have at least %d entries match the predicate, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasAtLeastMatch(test.threshold, predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.threshold, test.slice))
		}
	}
}

func TestSliceHasAtMostMatch(t *testing.T) {
	tests := []struct {
		slice        []string
		threshold    int
		expectedFail bool
	}{
		{slice: []string{"foo"}, threshold: 1},
		{slice: []string{"foo", "bar"}, threshold: 1},
		{slice: []string{"foo"}, threshold: 2},
		{slice: []string{"foo", "foo", "bar"}, threshold: 2},
		{slice: nil, threshold: 0},
		{slice: []string(nil), threshold: 0},
		{slice: []string{}, threshold: 0},
		{slice: []string{"bar"}, threshold: 0},
		{slice: []string{"foo"}, threshold: 0, expectedFail: true},
		{slice: []string{"foo", "foo"}, threshold: 1, expectedFail: true},
	}
	predicate := func(element string) bool { return element == "foo" }
	messageFormat := "expected slice to have at most %d entries match the predicate, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasAtMostMatch(test.threshold, predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.threshold, test.slice))
		}
	}
}

func TestSliceHasExactlyMatch(t *testing.T) {
	tests := []struct {
		slice        []string
		num          int
		expectedFail bool
	}{
		{slice: []string{"foo"}, num: 1},
		{slice: []string{"foo", "bar"}, num: 1},
		{slice: []string{"foo", "foo"}, num: 2},
		{slice: []string{"bar"}, num: 0},
		{slice: nil, num: 0},
		{slice: []string(nil), num: 0},
		{slice: []string{}, num: 0},
		{slice: []string{"bar"}, num: 1, expectedFail: true},
		{slice: []string{"foo"}, num: 0, expectedFail: true},
		{slice: []string{"foo", "foo"}, num: 1, expectedFail: true},
		{slice: nil, num: 1, expectedFail: true},
		{slice: []string(nil), num: 1, expectedFail: true},
		{slice: []string{}, num: 1, expectedFail: true},
	}
	predicate := func(element string) bool { return element == "foo" }
	messageFormat := "expected slice to have exactly %d entries match the predicate, but got %#v"

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		ThatSlice(fixture, test.slice).HasExactlyMatch(test.num, predicate)

		// then
		if !test.expectedFail {
			assertNoError(t, fixture)
		} else {
			assertErrorMessage(t, fixture, fmt.Sprintf(messageFormat, test.num, test.slice))
		}
	}
}

func TestSliceExtractingString(t *testing.T) {
	type tuple struct{ name string }
	tests := []struct {
		slice    []tuple
		expected []string
	}{
		{slice: nil, expected: []string(nil)},
		{slice: []tuple(nil), expected: []string(nil)},
		{slice: []tuple{{name: "foo"}}, expected: []string{"foo"}},
	}
	extractor := func(elem tuple) string { return elem.name }

	for _, test := range tests {
		// given
		fixture := new(fixtureT)

		// when
		assert := ThatSlice(fixture, test.slice).ExtractingStrings(extractor)

		// then
		if !ObjectsAreEqual(assert.actual, test.expected) {
			t.Errorf("expected new actual slice to equal %#v, but got %#v", test.expected, assert.actual)
		}
	}
}
