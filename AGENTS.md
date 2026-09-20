# go-casecheck contributor instructions

## Project overview

go-casecheck is a small Go assertion library. The module path is
go.osspkg.com/casecheck, the package name is casecheck, and the repository
intentionally uses a flat layout because it contains one public package and no
application binaries.

The module targets Go 1.26 and has no runtime dependencies. Keep the package
stdlib-only unless the user explicitly approves a dependency and its API need
cannot be met by the standard library.

## Project memory

Use the Chroma collection `chat_go-casecheck_memory` for durable project
decisions. Ensure the collection exists before querying or writing, query it
before non-trivial work, and store only concise, reusable API or architecture
facts. Never store credentials, tokens, or transient command output.

## Public API contract

- Assertion functions accept testing.TB, call t.Helper(), report a useful
  failure, and call t.FailNow() on failure.
- Do not start background goroutines, timers, or retry loops from assertions.
- Preserve the existing optional-message convention implemented by
  errorMessage: one argument is text; multiple arguments use the first as a
  format string.
- Equal and NotEqual compare dynamic types and values with reflect.DeepEqual.
- Contains and NotContains support strings, byte-substrings, map keys, and
  slice/array elements.
- ElementsMatch ignores order but counts duplicate elements and requires
  matching element types.
- JSONEqual accepts string, []byte, and json.RawMessage; malformed or
  unsupported input must produce an assertion failure, not a library panic.
- ErrorAs requires a valid non-nil pointer target and must use errors.As.
- InDelta and WithinDuration reject negative tolerances.

When adding an exported function, add a Go doc comment beginning with its
name and update README.md when the user-visible API or semantics change.

## Repository layout

- *.go: the public casecheck package.
- *_test.go: unit and regression tests colocated with the package.
- README.md: GitHub-facing usage and API documentation.
- Makefile: project automation; the goppy tool version is intentionally pinned
  for reproducible CI.
- .github/workflows/ci.yml: CI workflow with actions pinned to commit SHAs.

Do not create cmd/, internal/, or pkg/ directories unless the project gains a
separate executable or public package that justifies them.

## Implementation and review rules

1. Inspect git status and preserve unrelated worktree changes.
2. Prefer the standard library and the smallest change that satisfies the API
   contract.
3. Validate reflection kinds before calling methods such as IsNil.
4. Keep assertion failure paths deterministic and avoid exposing implementation
   panics to callers when input can be checked safely.
5. Do not use unsafe, mutable package state, network calls, or filesystem
   writes in assertion helpers.
6. Use apply_patch for source edits and gofmt for Go formatting.

## Required validation

For Go changes, run:

~~~bash
gofmt -w *.go
go test ./...
go test -race ./...
go test -shuffle=on -count=5 ./...
go vet ./...
go mod verify
git diff --check
~~~

For dependency or release-related changes, also run govulncheck ./... and
report if the vulnerability database is unavailable. Do not claim a clean
vulnerability scan when the command could not access its database.

Documentation-only changes should still be checked with git diff --check and
should not modify Go behavior.
