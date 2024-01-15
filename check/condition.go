package check

// Predicate is a function that returns if a value meets a condition.
type Predicate[T any] func(value T) bool

// Condition is a function to assert a condition.
type Condition[T any] func(value T)
