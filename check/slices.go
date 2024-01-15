package check

// SliceHasSize returns if the slice has the given size.
func SliceHasSize[T ~[]E, E any](slice T, size int) bool {
	return len(slice) == size
}

// SliceHasSizeGreaterThan returns if the slice has a size greater then the given value.
func SliceHasSizeGreaterThan[T ~[]E, E any](slice T, size int) bool {
	return len(slice) > size
}

// SliceHasSizeLessThan returns if the slice has a size less then the given value.
func SliceHasSizeLessThan[T ~[]E, E any](slice T, size int) bool {
	return len(slice) < size
}

// SliceIsEqual returns if the given slices are equal.
func SliceIsEqual[T ~[]E, E any](a T, b T) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA != lenB {
		return false
	}
	for i := range a {
		if !ObjectsAreEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// SliceContainsEntry returns if a slice contains an entry.
func SliceContainsEntry[T ~[]E, E any](slice T, entry E) bool {
	found := false
	for i := range slice {
		if ObjectsAreEqual(slice[i], entry) {
			found = true
			break
		}
	}
	return found
}

// SliceContainsEntryCount returns how often a slice contains an entry.
func SliceContainsEntryCount[T ~[]E, E any](slice T, entry E) int {
	num := 0
	for i := range slice {
		if ObjectsAreEqual(slice[i], entry) {
			num++
		}
	}
	return num
}

// SliceContainsSequence returns if a slice contains the given sequence.
func SliceContainsSequence[T ~[]E, E any](slice T, sequence T) bool {
	seqLen := len(sequence)
	sliceLen := len(slice)
	if seqLen <= sliceLen {
		for i := 0; i <= sliceLen-seqLen; i++ {
			window := slice[i : i+seqLen]
			if SliceIsEqual(sequence, window) {
				return true
			}
		}
	}
	return false
}

// SliceMatchPredicateCount returns how many elements in slice match the given predicate.
func SliceMatchPredicateCount[T ~[]E, E any](slice T, predicate Predicate[E]) int {
	matches := 0
	for i := range slice {
		if predicate(slice[i]) {
			matches++
		}
	}
	return matches
}

// SliceHasPrecicateMatches returns if the slice has the given number of elements that match the predicate.
func SliceHasPrecicateMatches[T ~[]E, E any](slice T, predicate Predicate[E], times int) bool {
	return SliceMatchPredicateCount(slice, predicate) == times
}
