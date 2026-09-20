package casecheck

import "testing"

func TestUnit_IsNil(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  bool
	}{
		{name: "nil interface", value: nil, want: true},
		{name: "nil pointer", value: (*int)(nil), want: true},
		{name: "nil map", value: (map[string]int)(nil), want: true},
		{name: "nil slice", value: ([]string)(nil), want: true},
		{name: "nil channel", value: (chan int)(nil), want: true},
		{name: "nil function", value: (func())(nil), want: true},
		{name: "integer", value: 0, want: false},
		{name: "boolean", value: false, want: false},
		{name: "struct", value: struct{}{}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNil(tt.value); got != tt.want {
				t.Errorf("isNil(%#v) = %t, want %t", tt.value, got, tt.want)
			}
		})
	}
}
