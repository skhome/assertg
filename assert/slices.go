package assert

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

func SliceMatchPredicateCount[T ~[]E, E any](slice T, predicate Predicate[E]) int {
	matches := 0
	for i := range slice {
		if predicate(slice[i]) {
			matches++
		}
	}
	return matches
}
