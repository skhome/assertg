package check

import "golang.org/x/exp/constraints"

// IntegersAreEqual returns wether both integer values are equal.
func IntegersAreEqual[T constraints.Integer](a, b T) bool {
	return a == b
}

// IntegerIsGreaterThan returns wether an integer value is greater then another.
func IntegerIsGreaterThan[T constraints.Integer](a, b T) bool {
	return a > b
}

// IntegerIsGreaterThanOrEqualTo returns wether an integer value is greater than or equal to another.
func IntegerIsGreaterThanOrEqualTo[T constraints.Integer](a, b T) bool {
	return a >= b
}

// IntegerIsLessThan returns wether an integer value is less then another.
func IntegerIsLessThan[T constraints.Integer](a, b T) bool {
	return a < b
}

// IntegerIsLessThanOrEqualTo returns wether an integer value is less than or equal to another.
func IntegerIsLessThanOrEqualTo[T constraints.Integer](a, b T) bool {
	return a <= b
}

// IntegerIsBetween returns wether an integer value is greater of equal to the start and less or equal to the end value.
func IntegerIsBetween[T constraints.Integer](value, start, end T) bool {
	return (start <= value) && (value <= end)
}

// IntegerIsEven returns wether an integer value is even.
func IntegerIsEven[T constraints.Integer](value T) bool {
	return value&1 == 0
}

// IntegerIsOdd returns wether an integer value is odd.
func IntegerIsOdd[T constraints.Integer](value T) bool {
	return value&1 != 0
}
