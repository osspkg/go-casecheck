package casecheck

import (
	"testing"
	"time"
)

func TestUnit_WithinDuration(t *testing.T) {
	expected := time.Date(2026, time.September, 21, 12, 0, 0, 0, time.UTC)
	WithinDuration(t, expected, expected.Add(time.Second), time.Second)

	if !withinDuration(expected, expected.Add(time.Second), time.Second) {
		t.Error("withinDuration() = false, want true")
	}
	if withinDuration(expected, expected.Add(2*time.Second), time.Second) {
		t.Error("withinDuration() = true, want false")
	}
	if withinDuration(expected, expected, -time.Second) {
		t.Error("withinDuration() = true, want false")
	}
}
