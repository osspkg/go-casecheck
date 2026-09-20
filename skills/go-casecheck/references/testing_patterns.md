# Testing patterns

## Table-driven tests

Use named subtests when the same assertion behavior is checked for several
inputs:

~~~go
func TestUnit_Status(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "ready", got: "ready", want: "ready"},
		{name: "stopped", got: "stopped", want: "stopped"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			casecheck.Equal(t, tt.want, tt.got)
		})
	}
}
~~~

Construct the assertion with the subtest's t. Assertions call FailNow, so the
failure terminates only that subtest's test goroutine.

## Choosing exact versus approximate checks

- Use Equal for strings, structs, slices, maps, and exact values.
- Use InDelta for floating-point results where a measured tolerance is part of
  the contract. Do not use a negative tolerance.
- Use WithinDuration for timestamps and deadlines. Pick a duration based on the
  contract, not on a value that merely makes a flaky test pass.
- Use JSONEqual only when JSON representation details such as whitespace and
  object-key order are not part of the contract.

## Failure-path behavior

Do not wrap an assertion in recover to continue after failure. If a test needs
to inspect multiple independent conditions, use separate subtests or ordinary
non-fatal checks for that test.

Do not call assertions from a background goroutine. testing.TB's FailNow
terminates the calling goroutine, and a background assertion would not stop the
test that owns it.

When an assertion's input is unsupported, fix the test input or use a more
specific standard-library check. Do not rely on implementation formatting to
make unsupported values comparable.
