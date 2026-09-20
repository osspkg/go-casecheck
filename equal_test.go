package casecheck

import (
	"math"
	"testing"
)

type constantStringer struct {
	value string
}

func (constantStringer) String() string {
	return "constant"
}

func TestUnit_ValuesEqual(t *testing.T) {
	type firstEmpty struct{}
	type secondEmpty struct{}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		want     bool
	}{
		{
			name:     "identical values",
			expected: struct{ Value int }{Value: 1},
			actual:   struct{ Value int }{Value: 1},
			want:     true,
		},
		{
			name:     "different named types",
			expected: firstEmpty{},
			actual:   secondEmpty{},
			want:     false,
		},
		{
			name:     "NaN values",
			expected: math.NaN(),
			actual:   math.NaN(),
			want:     false,
		},
		{
			name:     "different values with matching string representation",
			expected: constantStringer{value: "first"},
			actual:   constantStringer{value: "second"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valuesEqual(tt.expected, tt.actual); got != tt.want {
				t.Errorf("valuesEqual(%#v, %#v) = %t, want %t", tt.expected, tt.actual, got, tt.want)
			}
		})
	}
}
