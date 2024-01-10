package assert

import (
	"fmt"
)

type SliceAssert[E any] struct {
	t           TestingT
	message     string
	description string
	actual      []E
}

// As sets an optional description for this assertion.
func (a *SliceAssert[E]) As(format string, args ...any) *SliceAssert[E] {
	a.description = fmt.Sprintf(format, args...)
	return a
}

// WithFailMessage overrides the failure message.
func (a *SliceAssert[E]) WithFailMessage(format string, args ...any) *SliceAssert[E] {
	a.message = fmt.Sprintf(format, args...)
	return a
}

// failWithMessage records an assertion failure.
func (a *SliceAssert[E]) failWithMessage(format string, args ...any) {
	message := a.message
	if len(message) == 0 {
		message = fmt.Sprintf(format, args...)
	}
	Fail(a.t, message, a.description)
}

// IsNil verifies that the actual slice is nil.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string(nil)).IsNil()
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{}).IsNil()
func (a *SliceAssert[E]) IsNil() *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual != nil {
		a.failWithMessage("expected slice to be nil, but got %#v", a.actual)
	}
	return a
}

// IsNotNil verifies that the actual slice is not nil.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{}).IsNotNil()
//
//	// assertion will fail
//	assert.ThatSlice(t, []string(nil)).IsNotNil()
func (a *SliceAssert[E]) IsNotNil() *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual == nil {
		a.failWithMessage("expected slice to not be nil, but got %#v", a.actual)
	}
	return a
}

// IsEmpty verifies that the actual slice is null or empty.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string(nil)).IsEmpty()
//	assert.ThatSlice(t, []string{}).IsEmpty()
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"Hobbit"}).IsEmpty()
//	assert.ThatSlice(t, []int{1}).IsEmpty()
func (a *SliceAssert[E]) IsEmpty() *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.message == "" {
		a.WithFailMessage("expected slice to be empty, but got %#v", a.actual)
	}
	return a.HasSize(0)
}

// IsNotEmpty verifies that the actual slice is not empty.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"Nenya", "Narya", "Vilya"}).IsNotEmpty()
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{}).IsNotEmpty()
//	assert.ThatSlice(t, []string(nil)).IsNotEmpty()
func (a *SliceAssert[E]) IsNotEmpty() *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.message == "" {
		a.WithFailMessage("expected slice to not be empty, but got %#v", a.actual)
	}
	return a.HasSizeGreaterThan(0)
}

// HasSize verifies that the actual slice has the given size.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"Nenya", "Narya", "Vilya"}).HasSize(3)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{}).HasSize(1)
//	assert.ThatSlice(t, []string{"Nenya"}).HasSize(2)
func (a *SliceAssert[E]) HasSize(size int) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != size {
		a.failWithMessage("expected slice to have a size of %d, but got %#v", size, a.actual)
	}
	return a
}

// HasSizeGreaterThan verifies that the number of values in the actual slice is greater than the given boundary.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"Nenya", "Narya", "Vilya"}).HasSizeGreaterThan(2)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{}).HasSizeGreaterThan(0)
func (a *SliceAssert[E]) HasSizeGreaterThan(boundary int) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) <= boundary {
		a.failWithMessage("expected slice to have a size greater than %d, but got %#v", boundary, a.actual)
	}
	return a
}

// HasSizeLessThan verifies that the actual slice has a size greater then the given threshold.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"Nenya", "Narya", "Vilya"}).HasSizeLessThan(4)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"Frodo"}).HasSizeLessThan(1)
func (a *SliceAssert[E]) HasSizeLessThan(boundary int) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) >= boundary {
		a.failWithMessage("expected slice to have a size less than %d, but got %#v", boundary, a.actual)
	}
	return a
}

// HasSameSizeAs verifies that the actual slice has the same size as the given one.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"Frodo", "Sam"}).
//	       HasSameSizeAs([]string{"Merry", "Pippin"})
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"Frodo", "Bilbo"}).
//	       HasSameSizeAs([]string{"Gandalf"})
func (a *SliceAssert[E]) HasSameSizeAs(other []E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != len(other) {
		a.failWithMessage("expected slice to have the same size as %#v (%d), but got %#v (%d)", other, len(other), a.actual, len(a.actual))
	}
	return a
}

// Contains verifies that the actual slice contains the given elements in any order.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       Contains("b", "a").
//	       Contains("b", "a", "b")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       Contains("d")
func (a *SliceAssert[E]) Contains(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	notFound := false
	for i := range elements {
		if !SliceContainsEntry(a.actual, elements[i]) {
			notFound = true
			break
		}
	}
	if notFound {
		a.failWithMessage("expected slice to contain %#v, but got %#v", elements, a.actual)
	}
	return a
}

// ContainsOnly verifies that the actual slice contains only the given elements and nothing else, in any order and ignoring duplicates.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"a", "b", "c"}).ContainsOnly("a", "b", "c")
//	assert.ThatSlice(t, []string{"a", "b", "c"}).ContainsOnly("b", "c", "a")
//	assert.ThatSlice(t, []string{"a", "a", "b"}).ContainsOnly("a", "b")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"a", "b", "c"}).ContainsOnly("a", "b", "c", "d")
//	assert.ThatSlice(t, []string{"a", "b", "c"}).ContainsOnly("a", "b")
//	assert.ThatSlice(t, []string{"a", "b", "c"}).ContainsOnly("d", "e")
func (a *SliceAssert[E]) ContainsOnly(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	missed := false
	for i := range elements {
		if !SliceContainsEntry(a.actual, elements[i]) {
			missed = true
			break
		}
	}
	extraneous := false
	if !missed {
		for i := range a.actual {
			if !SliceContainsEntry(elements, a.actual[i]) {
				extraneous = true
				break
			}
		}
	}
	if missed || extraneous {
		a.failWithMessage("expected slice to contain only %#v, but got %#v", elements, a.actual)
	}
	return a
}

// ContainsOnlyOnce verifies that the actual slice contains the given elements only once.
//
//	// tests will pass
//	assert.ThatSlice(t, "a", "b", "c").ContainsOnlyOnce("a", "b")
//
//	// tests will fail
//	assert.ThatSlice(t, "a", "b", "a").ContainsOnlyOnce("a")
//	assert.ThatSlice(t, "a", "b", "c").ContainsOnlyOnce("d")
//	assert.ThatSlice(t, "a", "b", "c", "c").ContainsOnlyOnce("c", "d")
func (a *SliceAssert[E]) ContainsOnlyOnce(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := false
	for i := range elements {
		if SliceContainsEntryCount(a.actual, elements[i]) != 1 {
			found = true
			break
		}
	}
	if found {
		a.failWithMessage("expected slice to contain %#v only once, but got %#v", elements, a.actual)
	}
	return a
}

// ContainsExactly verifies that the actual slice contains exactly the given elements and nothing else, in order.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"vilya", "nenya", "varya"}).ContainsExactly("vilya", "nenya", "varya")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"vilya", "nenya", "varya"}).ContainsExactly("nenya", "vilya")
func (a *SliceAssert[E]) ContainsExactly(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	containsExactly := SliceIsEqual(elements, a.actual)
	if !containsExactly {
		a.failWithMessage("expected slice to contain exactly %#v, but got %#v", elements, a.actual)
	}
	return a
}

// ContainsExactlyInAnyOrder verifies that the actual slice contains exactly the given elements and nothing else, in any order.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"vilya", "nenya", "varya", "vilya"}).
//	       ContainsExactlyInAnyOrder("vilya", "vilya", "nenya", "varya")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"vilya", "nenya", "varya", "vilya"}).
//	       ContainsExactly("nenya", "vilya", "varya")
func (a *SliceAssert[E]) ContainsExactlyInAnyOrder(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	ok := true
	if len(elements) == len(a.actual) {
		for i := range elements {
			numElements := SliceContainsEntryCount(elements, elements[i])
			numActual := SliceContainsEntryCount(a.actual, elements[i])
			if numElements != numActual {
				ok = false
				break
			}
		}
	} else {
		ok = false
	}
	if !ok {
		a.failWithMessage("expected slice to contain exactly %v in any order, but got %#v", elements, a.actual)
	}
	return a
}

// ContainsSequence verifies that the actual slice contains the given sequence,
// without any extra values between them.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"vilya", "nenya", "narya"}).
//	  ContainsSequence("vilya", "nenya").
//	  ContainsSequence("nenya", "narya")
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"vilya", "nenya", "narya"}).
//	  ContainsSequence("vilya", "narya").
//	  ContainsSequence("nenya", "vilya")
func (a *SliceAssert[E]) ContainsSequence(sequence ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !SliceContainsSequence(a.actual, sequence) {
		a.failWithMessage("expected slice to contain the sequence %#v, but got %#v", sequence, a.actual)
	}
	return a
}

// DoesNotContainSequence verifies the actual slice does not contain the given sequence,
// without any extra values between them.
//
//	// assertion will pass
//	assert.ThatSlice(t, []string{"vilya", "nenya", "narya"}).
//	  DoesNotContainSequence("vilya", "narya").
//	  DoesNotContainSequence("nenya", "vilya")
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"vilya", "nenya", "narya"}).
//	  DoesNotContainSequence("vilya", "nenya").
//	  DoesNotContainSequence("nenya", "narya")
func (a *SliceAssert[E]) DoesNotContainSequence(sequence ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if SliceContainsSequence(a.actual, sequence) {
		a.failWithMessage("expected slice not to contain the sequence %#v, but got %#v", sequence, a.actual)
	}
	return a
}

// DoesNotContain verifies that the actual slice does not contain the given elements.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       DoesNotContain("d").
//	       DoesNotContain("d", "e")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       DoesNotContain("a").
//	       DoesNotContain("a", "b").
//	       DoesNotContain("c", "d")
func (a *SliceAssert[E]) DoesNotContain(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := false
	for i := range elements {
		if SliceContainsEntry(a.actual, elements[i]) {
			found = true
			break
		}
	}
	if found {
		a.failWithMessage("expected slice not to contain %#v, but got %#v", elements, a.actual)
	}
	return a
}

func (a *SliceAssert[E]) DoesNotHaveDuplicates() *SliceAssert[E] {
	return a
}

func (a *SliceAssert[E]) StartsWith(sequence ...E) *SliceAssert[E] {
	return a
}

func (a *SliceAssert[E]) EndsWith(sequence ...E) *SliceAssert[E] {
	return a
}

// ContainsAnyOf verifies that the actual slice contains at least one of the given elements.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       ContainsAnyOf("b").
//	       ContainsAnyOf("b", "c").
//	       ContainsAnyOf("a", "b", "c").
//	       ContainsAnyOf("a", "b", "c", "d").
//	       ContainsAnyOf("e","f", "c")
//
//	// assertions will fail
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       ContainsAnyOf("d").
//	       ContainsAnyOf("d", "e", "f", "g")
func (a *SliceAssert[E]) ContainsAnyOf(elements ...E) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	found := false
	for i := range elements {
		if SliceContainsEntry(a.actual, elements[i]) {
			found = true
			break
		}
	}
	if !found {
		a.failWithMessage("expected slice to contain any of %#v, but got %#v", elements, a.actual)
	}
	return a
}

// HasAll verifies that each element of the actual slice matches the given predicate.
//
//	// predicate
//	isSingleCharacter = func(value string) bool { return len(value) == 1 }
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"a", "b", "c"}).
//	       HasAll(isSingleCharacter)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"a", "b", "cc"}).
//	       HasAll(isSingleCharacter)
func (a *SliceAssert[E]) HasAll(predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.message) == 0 {
		a.WithFailMessage("expected slice to have all entries match the predicate, but got %#v", a.actual)
	}
	a.HasExactly(len(a.actual), predicate)
	return a
}

// HasNone verifies that no element of the actual slice matches the given predicate.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"Gandalf", "Elrond", "Galadriel"}).
//	       HasNone(isHobbit)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"Gandalf", "Frodo", "Elrond"}).
//	       HasNone(isHobbit)
func (a *SliceAssert[E]) HasNone(predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.message) == 0 {
		a.WithFailMessage("expected slice to have no entry match the predicate, but got %#v", a.actual)
	}
	a.HasExactly(0, predicate)
	return a
}

// HasAny verifies that any element of the actual slice matches the given predicate.
//
//	// assertions will pass
//	assert.ThatSlice(t, []string{"Luke", "Leia", "Han"}).
//	       HasAny(isJedi)
//
//	// assertion will fail
//	assert.ThatSlice(t, []string{"Leia", "Han", "Lando"}).
//	       HasAllMatch(isJedi)
func (a *SliceAssert[E]) HasAny(predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.message) == 0 {
		a.WithFailMessage("expected slice to have any entry match the predicate, but got %#v", a.actual)
	}
	a.HasAtLeast(1, predicate)
	return a
}

// HasAtLeast verifies that there are at least n elements in the actual slice that match the given predicate.
//
//	// assertions will pass
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtLeast(2, isOddNumber)
//
//	// assertion will fail
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtLeast(3, isOddNumber)
func (a *SliceAssert[E]) HasAtLeast(n int, predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := SliceMatchPredicateCount(a.actual, predicate)
	if matches < n {
		a.failWithMessage("expected slice to have at least %d entries match the predicate, but got %#v", n, a.actual)
	}
	return a
}

// HasAtMost verifies that there are at most n elements in the actual slice that match the given predicate.
//
//	// assertions will pass
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtMost(2, isOddNumber).
//	       HasAtMost(3, isOddNumber)
//
//	// assertion will fail
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtMost(1, isOddNumber)
func (a *SliceAssert[E]) HasAtMost(n int, predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := SliceMatchPredicateCount(a.actual, predicate)
	if matches > n {
		a.failWithMessage("expected slice to have at most %d entries match the predicate, but got %#v", n, a.actual)
	}
	return a
}

// HasExactly verifies that there are exactly n elements in the actual slice that match the given predicate.
//
//	// assertions will pass
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtMost(2, isOddNumber).
//
//	// assertion will fail
//	assert.ThatSlice(t, []int{1, 2, 3}).
//	       HasAtMost(1, isOddNumber)
//	       HasAtMost(3, isOddNumber)
func (a *SliceAssert[E]) HasExactly(n int, predicate Predicate[E]) *SliceAssert[E] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	matches := SliceMatchPredicateCount(a.actual, predicate)
	if matches != n {
		a.failWithMessage("expected slice to have exactly %d entries match the predicate, but got %#v", n, a.actual)
	}
	return a
}

// ExtractingStrings extracts a new slice of strings from the actual slice using the given extractor function.
// The extracted slice becomes the the new object under test.
//
//	type TolkienCharacter struct {
//	  name    string
//	  age     int
//	  species string
//	}
//
//	fellowship := []TolkienCharacter{
//	  {name: "Frodo", age: 33, species: "Hobbit"},
//	  {name: "Sam", age: 38, species: "Hobbit"},
//	  {name: "Gandalf", age: 2020, species: "Maia"},
//	  {name: "Legolas", age: 1000, species: "Elf"},
//	}
//
//	characterName := func(character TolkienCharacter) string { return character.name }
//
//	assert.ThatSlice(t, fellowship).
//	       ExtractingStrings(characterName).
//	       ContainsExactly("Frodo", "Sam", "Gandalf", "Legolas").
//	       DoesNotContain("Elrond")
func (a *SliceAssert[E]) ExtractingStrings(extractor func(elem E) string) *SliceAssert[string] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	var extracted []string
	for _, elem := range a.actual {
		extracted = append(extracted, extractor(elem))
	}
	return &SliceAssert[string]{t: a.t, actual: extracted}
}

// ExtractingInts extracts a new slice of ints from the actual slice using the given extractor function.
// The extracted slice becomes the the new object under test.
//
//	type TolkienCharacter struct {
//	  name    string
//	  age     int
//	  species string
//	}
//
//	fellowship := []TolkienCharacter{
//	  {name: "Frodo", age: 33, species: "Hobbit"},
//	  {name: "Sam", age: 38, species: "Hobbit"},
//	  {name: "Gandalf", age: 2020, species: "Maia"},
//	  {name: "Legolas", age: 1000, species: "Elf"},
//	}
//
//	characterAge := func(character TolkienCharacter) int { return character.age }
//
//	assert.ThatSlice(t, fellowship).
//	       ExtractingInts(characterAge).
//	       Contains(33, 2020).
//	       HasSize(4)
func (a *SliceAssert[E]) ExtractingInts(extractor func(elem E) int) *SliceAssert[int] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	var extracted []int
	for _, elem := range a.actual {
		extracted = append(extracted, extractor(elem))
	}
	return &SliceAssert[int]{t: a.t, actual: extracted}
}

// Extracting extracts a new slice from the actual slice using the given extractor function.
// The extracted slice becomes the the new object under test.
//
//	type Species string
//	type TolkienCharacter struct {
//	  name    string
//	  species Species
//	}
//
//	fellowship := []TolkienCharacter{
//	  {name: "Frodo", species: "Hobbit"},
//	  {name: "Legolas", species: "Elf"},
//	}
//
//	characterSpecies := func(character TolkienCharacter) any { return character.species }
//
//	assert.ThatSlice(t, fellowship).
//	       Extracting(characterSpecies).
//	       Contains(Species("Elf"))
func (a *SliceAssert[E]) Extracting(extractor func(elem E) any) *SliceAssert[any] {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	var extracted []any
	for _, elem := range a.actual {
		extracted = append(extracted, extractor(elem))
	}
	return &SliceAssert[any]{t: a.t, actual: extracted}
}
