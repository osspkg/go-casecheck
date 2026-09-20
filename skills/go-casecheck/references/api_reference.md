# go-casecheck API reference

Import path:

~~~go
import "go.osspkg.com/casecheck"
~~~

All assertions accept testing.TB as their first argument and optional
args ...interface{} after the assertion inputs. A failed assertion reports
through t.Error and stops the current test with t.FailNow.

## Values

| Function | Inputs | Semantics |
| --- | --- | --- |
| Equal | expected, actual | Same dynamic type and reflect.DeepEqual value |
| NotEqual | expected, actual | Not the same dynamic type/value |
| True / False | bool | Exact boolean check |
| Nil / NotNil | value | Nil-interface and typed-nil aware; non-nilable values are safe |

## Errors

~~~go
casecheck.NoError(t, err)
casecheck.Error(t, err)
casecheck.ErrorContains(t, err, "timeout")
casecheck.ErrorIs(t, err, context.DeadlineExceeded)

var target *MyError
casecheck.ErrorAs(t, err, &target)
~~~

ErrorAs requires a non-nil pointer to an error type or interface. It validates
the target before calling errors.As, so malformed targets become assertion
failures instead of reflection panics.

## Collections

Empty, NotEmpty, and Len support string, slice, map, array, and channel values.
A nil collection is empty. For channels, Len uses the current queue length and
never the channel capacity.

Contains and NotContains support:

- string substring;
- byte-slice substring;
- map key;
- slice or array element.

ElementsMatch accepts only slices and arrays. It ignores order, matches each
element at most once, preserves duplicate counts, and requires equal element
types. Nil and allocated-empty slices with the same element type both contain
zero elements and match.

## Execution and numeric/time checks

~~~go
casecheck.Panics(t, func() {
	panic("expected")
})
casecheck.NotPanics(t, func() {
	runOperation()
})

casecheck.InDelta(t, float64(10), result, 0.01)
casecheck.WithinDuration(t, expectedTime, actualTime, time.Second)
~~~

Panics and NotPanics run the supplied function synchronously. InDelta supports
float32 and float64 through a type parameter and rejects negative tolerances.
WithinDuration compares absolute time difference and rejects a negative
duration.

## Text and JSON

Match and NotMatch accept a string and a regular-expression pattern. Compilation
errors are assertion failures:

~~~go
casecheck.Match(t, response, "^status: [0-9]+$")
casecheck.NotMatch(t, username, "\\s")
~~~

JSONEqual accepts string, []byte, or json.RawMessage. It decodes both inputs
using encoding/json, so insignificant whitespace and object-key order are
ignored. Invalid JSON and unsupported input types are assertion failures.

~~~go
casecheck.JSONEqual(
	t,
	"{\"enabled\":true,\"items\":[1,2]}",
	[]byte(`{ "items": [1, 2], "enabled": true }`),
)
~~~

## Message arguments

The first optional argument alone is rendered as a message. With two or more
optional arguments, the first is passed as a fmt.Sprintf format string:

~~~go
casecheck.Equal(t, want, got, "user id: %d", userID)
~~~
