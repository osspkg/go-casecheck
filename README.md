# go-casecheck

[![Go Reference](https://pkg.go.dev/badge/go.osspkg.com/casecheck.svg)](https://pkg.go.dev/go.osspkg.com/casecheck)
[![Go Version](https://img.shields.io/github/go-mod/go-version/osspkg/go-casecheck)](https://go.dev/)
[![CI](https://github.com/osspkg/go-casecheck/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/osspkg/go-casecheck/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/osspkg/go-casecheck)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/osspkg/go-casecheck)](https://goreportcard.com/report/github.com/osspkg/go-casecheck)

go-casecheck is a small, standard-library-only assertion library for Go
tests. It builds on testing.TB, marks assertion functions as helpers, and
stops the current test goroutine when an assertion fails.

## Getting started

Install the module:

~~~bash
go get go.osspkg.com/casecheck
~~~

Use the assertions from a regular Go test:

~~~go
package user_test

import (
	"errors"
	"fmt"
	"testing"

	"go.osspkg.com/casecheck"
)

var errNotReady = errors.New("not ready")

func TestUnit_User(t *testing.T) {
	err := fmt.Errorf("load user: %w", errNotReady)
	casecheck.ErrorIs(t, err, errNotReady)

	roles := []string{"reader", "admin"}
	casecheck.NotEmpty(t, roles)
	casecheck.Len(t, roles, 2)
	casecheck.Contains(t, roles, "admin")
	casecheck.ElementsMatch(t, []string{"admin", "reader"}, roles)

	casecheck.InDelta(t, 3.14, 3.141, 0.01)
	casecheck.JSONEqual(t, "{\"active\":true,\"roles\":[\"admin\"]}", "{\"roles\":[\"admin\"],\"active\":true}")
}
~~~

## Assertions

The table summarizes which assertion to choose and what it verifies:

| Method               | Use it for                                                                                                                                  |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `Equal`              | Exact comparison of dynamic types and values via `reflect.DeepEqual`.                                                                       |
| `NotEqual`           | Verifying that values do not have the same type and value.                                                                                  |
| `True` / `False`     | Checking a boolean condition.                                                                                                               |
| `Nil` / `NotNil`     | Checking nil interfaces and typed-nil values; non-nilable values are handled safely.                                                        |
| `Contains`           | Finding a substring, byte-slice substring, map key, or element in a slice/array.                                                            |
| `NotContains`        | Verifying that a string, byte slice, map, slice, or array does not contain a value.                                                         |
| `Empty` / `NotEmpty` | Checking whether a string, slice, map, array, or channel has no elements.                                                                   |
| `Len`                | Comparing the length of a string, slice, map, array, or channel queue.                                                                      |
| `ElementsMatch`      | Comparing slices/arrays without order while preserving duplicate counts. Element types must match.                                          |
| `NoError`            | Requiring an error result to be `nil`.                                                                                                      |
| `Error`              | Requiring an error result to be non-`nil`.                                                                                                  |
| `ErrorContains`      | Checking that an error message contains specific text.                                                                                      |
| `ErrorIs`            | Matching a sentinel error through wrapping with `errors.Is`.                                                                                |
| `ErrorAs`            | Extracting a typed error from an error chain with `errors.As`; pass a pointer such as `&target`.                                            |
| `Panics`             | Verifying that a supplied function panics.                                                                                                  |
| `NotPanics`          | Verifying that a supplied function completes without panicking.                                                                             |
| `InDelta`            | Comparing `float32` or `float64` values within a non-negative tolerance.                                                                    |
| `WithinDuration`     | Comparing two `time.Time` values within a non-negative duration.                                                                            |
| `Match`              | Verifying that a string matches a regular-expression pattern.                                                                               |
| `NotMatch`           | Verifying that a string does not match a regular-expression pattern.                                                                        |
| `JSONEqual`          | Comparing JSON semantically while ignoring insignificant whitespace and object-key order. Accepts `string`, `[]byte`, or `json.RawMessage`. |

All assertions accept testing.TB and optional message arguments. With one
message argument it is displayed as text; with multiple arguments the first
argument is treated as a fmt.Sprintf-style format string.

## Development

The module targets Go 1.26 and intentionally has no runtime dependencies.
Run the local quality gates from the repository root:

~~~bash
gofmt -w *.go
go test ./...
go test -race ./...
go vet ./...
go mod verify
~~~

The repository also provides make ci, which uses the pinned goppy tool to run
the project checks. Known dependency vulnerabilities can be checked with:

~~~bash
govulncheck ./...
~~~

## Contributing

Keep changes focused and preserve the public assertion semantics. Add a
regression test for every new behavior or bug fix, keep exported API comments
accurate, and run the quality gates above before opening a pull request.

## License

This project is distributed under the [BSD 3-Clause License](./LICENSE).
