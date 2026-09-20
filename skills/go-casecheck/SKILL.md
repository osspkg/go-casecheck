---
name: go-casecheck
description: Use go-casecheck for Go test assertions, examples, and API-accurate test documentation when the repository imports go.osspkg.com/casecheck.
---

# Go Casecheck

Use this skill when writing or reviewing tests that use go.osspkg.com/casecheck,
or when producing examples and documentation for the library. Keep the result
compatible with the current public package; do not invent testify-style APIs.

## Core contract

- Every assertion accepts testing.TB, calls Helper, reports a useful failure,
  and calls FailNow. Assertions are synchronous and must not create goroutines,
  timers, retries, or external I/O.
- The package is standard-library-only. Do not add a test-framework dependency
  just to use these assertions.
- Preserve the optional-message convention: one message argument is text;
  with multiple arguments, the first is a fmt.Sprintf-style format string.
- Use ErrorIs for wrapped sentinel errors and ErrorAs with a non-nil pointer
  target such as &target.
- Use Equal for exact dynamic type/value comparison, InDelta for float32 or
  float64 tolerances, and WithinDuration for time.Time values.

## Assertion selection

- Collections: Empty, NotEmpty, and Len support string, slice, map, array, and
  channel values. Len reports queued channel elements, not capacity.
- Contains searches string or byte substrings, map keys, and slice/array
  elements. It does not search map values.
- ElementsMatch accepts slices or arrays, ignores order, preserves duplicate
  counts, and requires matching element types.
- Panics and NotPanics execute only the supplied function in the current
  goroutine.
- Match and NotMatch take a regular-expression string. Invalid patterns are
  assertion failures.
- JSONEqual accepts string, []byte, and json.RawMessage; whitespace and object
  key order do not affect equality.
- Nil is safe for non-nilable values and reports a failed assertion rather than
  calling reflect.Value.IsNil on an unsupported kind.

## Workflow

1. Inspect the current module and existing tests before choosing an assertion.
2. Read [references/api_reference.md](references/api_reference.md) for exact
   signatures and edge-case semantics when the choice is non-obvious.
3. Copy the smallest relevant snippet from
   [examples/usage.md](examples/usage.md), adapting names and expected values.
4. Add a focused regression test for new behavior and keep examples compilable.
5. Run gofmt, go test ./..., go test -race ./..., and go vet ./... for code
   changes. For documentation-only changes, at minimum run git diff --check.

Do not mask a failed assertion with a recover, call assertion helpers from a
background goroutine, or pass unsupported values to an assertion merely to
make a test compile.
