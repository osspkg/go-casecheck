package casecheck

import (
	"math"
	"testing"
)

func TestUnit_InDelta(t *testing.T) {
	InDelta(t, 1.0, 1.05, 0.1)
	InDelta(t, float32(1.0), float32(1.05), float32(0.1))

	tests := []struct {
		name     string
		expected float64
		actual   float64
		delta    float64
		want     bool
	}{
		{name: "within delta", expected: 1, actual: 1.05, delta: 0.1, want: true},
		{name: "outside delta", expected: 1, actual: 1.2, delta: 0.1, want: false},
		{name: "negative delta", expected: 1, actual: 1, delta: -1, want: false},
		{name: "same infinity", expected: math.Inf(1), actual: math.Inf(1), delta: 0, want: true},
		{name: "NaN", expected: math.NaN(), actual: math.NaN(), delta: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inDelta(tt.expected, tt.actual, tt.delta); got != tt.want {
				t.Errorf("inDelta(%v, %v, %v) = %t, want %t", tt.expected, tt.actual, tt.delta, got, tt.want)
			}
		})
	}
}
