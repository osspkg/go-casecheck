package casecheck

import (
	"errors"
	"fmt"
	"testing"
)

var errSentinel = errors.New("sentinel")

type typedError struct {
	message string
}

func (e *typedError) Error() string {
	return e.message
}

func TestUnit_ErrorIs(t *testing.T) {
	err := fmt.Errorf("operation: %w", errSentinel)
	ErrorIs(t, err, errSentinel)
}

func TestUnit_ErrorAs(t *testing.T) {
	want := &typedError{message: "failed"}
	err := fmt.Errorf("operation: %w", want)

	var got *typedError
	ErrorAs(t, err, &got)
	if got != want {
		t.Errorf("ErrorAs() target = %#v, want %#v", got, want)
	}
}

func TestUnit_ErrorAsValidation(t *testing.T) {
	var typedTarget *typedError
	if matched, valid := errorAs(fmt.Errorf("operation: %w", errSentinel), &typedTarget); matched || !valid {
		t.Errorf("errorAs() = (%t, %t), want (false, true)", matched, valid)
	}

	if matched, valid := errorAs(errSentinel, typedTarget); matched || valid {
		t.Errorf("errorAs() = (%t, %t), want (false, false)", matched, valid)
	}

	var invalidTarget int
	if matched, valid := errorAs(errSentinel, &invalidTarget); matched || valid {
		t.Errorf("errorAs() = (%t, %t), want (false, false)", matched, valid)
	}
}
