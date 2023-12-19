# AssertG - Fluent Assertions for Go

[![CI](https://github.com/skhome/assertg/actions/workflows/main.yaml/badge.svg)](https://github.com/skhome/assertg/actions/workflows/main.yaml)

Go package to provide an intuitive set of strongly typed assertions for use in unit testing.

## Inspiration and Goals

AssertG takes inspiration from [AssertJ](https://github.com/assertj/assertj) (an assertion library for Java) and [Testify](https://github.com/stretchr/testify) (a toolkit with common assertions for Go).

## The `assert` package

Example:
```go
import (
  "testing"
  "github.com/skhome/assertg/assert"
)

func TestSomething(t *testing.T) {
  assert.ThatString(t, "foobar").
    IsNotEmpty().
    StartsWith("foo").
    EndsWith("bar")

  assert.ThatSlice(t, []int{2, 4, 6}).
    HasSize(3).
    HasAllMatch(func(num int) bool { return num % 2 == 0 })
}
```

