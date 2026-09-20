package casecheck

import "testing"

func TestUnit_Panics(t *testing.T) {
	Panics(t, func() { panic("expected") })
	NotPanics(t, func() {})

	if !panics(func() { panic("expected") }) {
		t.Error("panics() = false, want true")
	}
	if panics(func() {}) {
		t.Error("panics() = true, want false")
	}
}
