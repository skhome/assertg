package check

import "golang.org/x/exp/constraints"

// FloatsAreEqual returns wether both values are equal.
func FloatsAreEqual[T constraints.Float](a, b T) bool {
	return a == b
}

// FloatIsGreaterThan returns wether a float value is greater then another.
func FloatIsGreaterThan[T constraints.Float](a, b T) bool {
	return a > b
}

// FloatIsGreaterThanOrEqualTo returns wether a float value is greater than or equal to another.
func FloatIsGreaterThanOrEqualTo[T constraints.Float](a, b T) bool {
	return a >= b
}

// FloatIsLessThan returns wether a float value is less then another.
func FloatIsLessThan[T constraints.Float](a, b T) bool {
	return a < b
}

// FloatIsLessThanOrEqualTo returns wether a float value is less than or equal to another.
func FloatIsLessThanOrEqualTo[T constraints.Float](a, b T) bool {
	return a <= b
}

// FloatIsBetween returns wether a float value is greater of equal to the start and less or equal to the end value.
func FloatIsBetween[T constraints.Float](value, start, end T) bool {
	return (start <= value) && (value <= end)
}
