# Usage examples

These snippets are intentionally small and can be copied into a Go test after
adjusting the package-specific values.

## Errors and typed errors

~~~go
func TestUnit_Load(t *testing.T) {
	err := fmt.Errorf("load config: %w", context.DeadlineExceeded)
	casecheck.ErrorIs(t, err, context.DeadlineExceeded)

	pathErrValue := &os.PathError{Op: "open", Path: "config", Err: fs.ErrNotExist}
	err = fmt.Errorf("load config: %w", pathErrValue)
	var pathErr *os.PathError
	casecheck.ErrorAs(t, err, &pathErr)
}
~~~

Use ErrorAs with a pointer-to-target. Passing target directly is invalid because
errors.As needs a location where it can store the matched value.

## Collections

~~~go
func TestUnit_Roles(t *testing.T) {
	roles := []string{"reader", "admin", "admin"}

	casecheck.NotEmpty(t, roles)
	casecheck.Len(t, roles, 3)
	casecheck.Contains(t, roles, "admin")
	casecheck.NotContains(t, roles, "owner")
	casecheck.ElementsMatch(t, []string{"admin", "reader", "admin"}, roles)
}
~~~

Contains searches elements in slices and arrays, keys in maps, and substrings
in strings or byte slices. ElementsMatch still detects a missing duplicate.

## Panic and tolerance checks

~~~go
func TestUnit_Operation(t *testing.T) {
	casecheck.Panics(t, func() {
		mustParse("invalid")
	})
	casecheck.NotPanics(t, func() {
		parse("valid")
	})

	casecheck.InDelta(t, 100.0, measured, 0.5)
	casecheck.WithinDuration(t, started, finished, 2*time.Second)
}
~~~

The supplied functions run synchronously in the current test goroutine. Keep
the function body focused; use a separate test when setup or cleanup is
substantial.

## JSON and regular expressions

~~~go
func TestUnit_Response(t *testing.T) {
	casecheck.JSONEqual(
		t,
		"{\"ok\":true,\"items\":[1,2]}",
		[]byte("{ \"items\": [1, 2], \"ok\": true }"),
	)
	casecheck.Match(t, "request-id: 42", "^request-id: [0-9]+$")
	casecheck.NotMatch(t, username, "\\s")
}
~~~

JSONEqual parses both values with encoding/json. It is useful when formatting
or object-key order is irrelevant; it is not a schema validator.

## Named assertion messages

~~~go
casecheck.Equal(t, expected, actual, "record id: %d", recordID)
~~~

With one optional argument, the argument is plain message text. With multiple
arguments, the first argument is a fmt.Sprintf-style format string.
