package casecheck

import "testing"

func TestUnit_Empty(t *testing.T) {
	Empty(t, "")
	Empty(t, []int(nil))
	Empty(t, map[string]int{})
	Empty(t, make(chan int))
	NotEmpty(t, "value")
	NotEmpty(t, []int{1})
	NotEmpty(t, map[string]int{"key": 1})
}

func TestUnit_Length(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  int
		ok    bool
	}{
		{name: "string", value: "go", want: 2, ok: true},
		{name: "slice", value: []int{1, 2}, want: 2, ok: true},
		{name: "array", value: [2]int{}, want: 2, ok: true},
		{name: "map", value: map[string]int{"key": 1}, want: 1, ok: true},
		{name: "channel", value: make(chan int, 2), want: 0, ok: true},
		{name: "integer", value: 1, want: 0, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := length(tt.value)
			if got != tt.want || ok != tt.ok {
				t.Errorf("length(%#v) = (%d, %t), want (%d, %t)", tt.value, got, ok, tt.want, tt.ok)
			}
		})
	}

	Len(t, []int{1, 2}, 2)
}

func TestUnit_ElementsMatch(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		want     bool
		ok       bool
	}{
		{name: "same elements different order", expected: []int{1, 2, 3}, actual: []int{3, 1, 2}, want: true, ok: true},
		{name: "duplicate elements", expected: []int{1, 1, 2}, actual: []int{1, 2, 1}, want: true, ok: true},
		{name: "different multiplicity", expected: []int{1, 1, 2}, actual: []int{1, 2, 2}, want: false, ok: true},
		{name: "different element type", expected: []int{1}, actual: []int64{1}, want: false, ok: true},
		{name: "nil and empty", expected: []int(nil), actual: []int{}, want: true, ok: true},
		{name: "unsupported value", expected: "value", actual: "value", want: false, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := elementsMatch(tt.expected, tt.actual)
			if got != tt.want || ok != tt.ok {
				t.Errorf("elementsMatch(%#v, %#v) = (%t, %t), want (%t, %t)", tt.expected, tt.actual, got, ok, tt.want, tt.ok)
			}
		})
	}

	ElementsMatch(t, []int{1, 2}, []int{2, 1})
}

func TestUnit_Contains(t *testing.T) {
	Contains(t, "casecheck", "check")
	Contains(t, []byte("casecheck"), []byte("check"))
	Contains(t, map[string]int{"key": 1}, "key")
	Contains(t, []string{"first", "second"}, "second")
	Contains(t, [2]int{1, 2}, 2)
	NotContains(t, []int{1, 2}, 3)

	tests := []struct {
		name       string
		searchData interface{}
		need       interface{}
		want       bool
		ok         bool
	}{
		{name: "slice element", searchData: []int{1, 2}, need: 2, want: true, ok: true},
		{name: "array element", searchData: [2]string{"first", "second"}, need: "second", want: true, ok: true},
		{name: "missing element", searchData: []int{1, 2}, need: 3, want: false, ok: true},
		{name: "unsupported type", searchData: 1, need: 1, want: false, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := contains(tt.searchData, tt.need)
			if got != tt.want || ok != tt.ok {
				t.Errorf("contains(%#v, %#v) = (%t, %t), want (%t, %t)", tt.searchData, tt.need, got, ok, tt.want, tt.ok)
			}
		})
	}
}
