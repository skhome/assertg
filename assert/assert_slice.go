package assert

import (
	"fmt"
)

type SliceAssert[T ~[]E, E any] struct {
	t           TestingT
	message     string
	description string
	actual      T
}

// As sets an optional description for this assertion.
func (a *SliceAssert[T, E]) As(format string, args ...any) *SliceAssert[T, E] {
	a.description = fmt.Sprintf(format, args...)
	return a
}

// WithFailMessage overrides the failure message.
func (a *SliceAssert[T, E]) WithFailMessage(format string, args ...any) *SliceAssert[T, E] {
	a.message = fmt.Sprintf(format, args...)
	return a
}

// hasFailMessage returns if a failure message has been set already.
func (a *SliceAssert[T, E]) hasFailMessage() bool {
	return len(a.message) > 0
}

// IsNil checks that the actual slice is nil
func (a *SliceAssert[T, E]) IsNil() *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual != nil {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to be nil, but got %#v", a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// IsNotNil checks that the actual slice is not nil
func (a *SliceAssert[T, E]) IsNotNil() *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual == nil {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to not be nil, but got %#v", a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasSize checks that the actual slice has the given size.
func (a *SliceAssert[T, E]) HasSize(size int) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != size {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have a size of %d, but got %#v", size, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasSizeGreaterThan checks that the actual slice has a size greater then the given threshold.
func (a *SliceAssert[T, E]) HasSizeGreaterThan(threshold int) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) <= threshold {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have a size greater than %d, but got %#v", threshold, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasSizeLessThan checks that the actual slice has a size greater then the given threshold.
func (a *SliceAssert[T, E]) HasSizeLessThan(threshold int) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) >= threshold {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have a size less than %d, but got %#v", threshold, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// IsEmpty checks that the actual slice is empty.
func (a *SliceAssert[T, E]) IsEmpty() *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	a.WithFailMessage("expected slice to be empty, but got %#v", a.actual).HasSize(0)
	return a
}

// IsNotEmpty checks that the actual slice is not empty.
func (a *SliceAssert[T, E]) IsNotEmpty() *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	a.WithFailMessage("expected slice to not be empty, but got %#v", a.actual).HasSizeGreaterThan(0)
	return a
}

func containsEntry[T ~[]E, E any](slice T, entry E) bool {
	found := false
	for i := range slice {
		if ObjectsAreEqual(slice[i], entry) {
			found = true
		}
	}
	return found
}

// Contains checks that the actual slice contains the given element.
func (a *SliceAssert[T, E]) Contains(element E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := containsEntry(a.actual, element)
	if !found {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to contain %#v, but got %#v", element, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// DoesNotContain checks that the actual slice does not contain the given element.
func (a *SliceAssert[T, E]) DoesNotContain(element E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := containsEntry(a.actual, element)
	if found {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to not contain %#v, but got %#v", element, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// ContainsAnyOf checks that the actual slice contains at least one of the elements.
func (a *SliceAssert[T, E]) ContainsAnyOf(elements ...E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := false
	for i := range elements {
		if containsEntry(a.actual, elements[i]) {
			found = true
		}
	}
	if !found {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to contain any of %#v, but got %#v", elements, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// ContainsAllOf checks that the actual slice contains all of the elements.
func (a *SliceAssert[T, E]) ContainsAllOf(elements ...E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := 0
	for i := range elements {
		if containsEntry(a.actual, elements[i]) {
			found++
		}
	}
	if found != len(elements) || len(a.actual) == 0 || len(elements) == 0 {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to contain any of %#v, but got %#v", elements, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// ContainsNoneOf checks that the actual slice contains none of the elements.
func (a *SliceAssert[T, E]) ContainsNoneOf(elements ...E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := 0
	for i := range elements {
		if containsEntry(a.actual, elements[i]) {
			found++
		}
	}
	if found != 0 {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to contain none of %#v, but got %#v", elements, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// ContainsExactly checks that the actual slice contains exactly the elements.
func (a *SliceAssert[T, E]) ContainsExactly(elements ...E) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := 0
	for i := range elements {
		if containsEntry(a.actual, elements[i]) {
			found++
		}
	}
	if found != len(elements) || len(a.actual) != len(elements) {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to contain exactly %#v, but got %#v", elements, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

func numberOfPredicateMatches[T ~[]E, E any](slice T, predicate Predicate[E]) int {
	matches := 0
	for i := range slice {
		if predicate(slice[i]) {
			matches++
		}
	}
	return matches
}

// HasExactlyMatch checks that there are exactly n elements in the actual slice that match the given predicate.
func (a *SliceAssert[T, E]) HasExactlyMatch(n int, predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := numberOfPredicateMatches(a.actual, predicate)
	if matches != n {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have exactly %d entries match the predicate, but got %#v", n, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasAtLeastMatch checks that there are at least n elements in the actual slice that match the given predicate.
func (a *SliceAssert[T, E]) HasAtLeastMatch(n int, predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := numberOfPredicateMatches(a.actual, predicate)
	if matches < n {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have at least %d entries match the predicate, but got %#v", n, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasAtMostMatch checks that there are at most n elements in the actual slice that match the given predicate.
func (a *SliceAssert[T, E]) HasAtMostMatch(n int, predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := numberOfPredicateMatches(a.actual, predicate)
	if matches > n {
		if !a.hasFailMessage() {
			a.WithFailMessage("expected slice to have at most %d entries match the predicate, but got %#v", n, a.actual)
		}
		Fail(a.t, a.message, a.description)
	}
	return a
}

// HasAnyMatch checks whether any element of the actual slice matches the given predicate.
func (a *SliceAssert[T, E]) HasAnyMatch(predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !a.hasFailMessage() {
		a.WithFailMessage("expected slice to have any entry match the predicate, but got %#v", a.actual)
	}
	a.HasAtLeastMatch(1, predicate)
	return a
}

// HasNoneMatch checks whether no element of the actual slice match the given predicate.
func (a *SliceAssert[T, E]) HasNoneMatch(predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !a.hasFailMessage() {
		a.WithFailMessage("expected slice to have no entry match the predicate, but got %#v", a.actual)
	}
	a.HasExactlyMatch(0, predicate)
	return a
}

// HasAllMatch checks whether all elements of the actual slice match the given predicate.
func (a *SliceAssert[T, E]) HasAllMatch(predicate Predicate[E]) *SliceAssert[T, E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !a.hasFailMessage() {
		a.WithFailMessage("expected slice to have all entries match the predicate, but got %#v", a.actual)
	}
	a.HasExactlyMatch(len(a.actual), predicate)
	return a
}

// ExtractingStrings extracts a new slice of strings from the actual slice using the given extractor function.
// The extracted slice becomes the the new object under test.
func (a *SliceAssert[T, E]) ExtractingStrings(extractor func(elem E) string) *SliceAssert[[]string, string] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	var extracted []string
	for _, elem := range a.actual {
		extracted = append(extracted, extractor(elem))
	}
	return &SliceAssert[[]string, string]{t: a.t, actual: extracted}
}
