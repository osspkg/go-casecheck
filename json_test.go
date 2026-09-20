package casecheck

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnit_JSONEqual(t *testing.T) {
	JSONEqual(t, `{"name":"casecheck","items":[1,2]}`, []byte(`{ "items": [1, 2], "name": "casecheck" }`))
}

func TestUnit_DecodeJSON(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  interface{}
		ok    bool
	}{
		{name: "string", value: `{"name":"casecheck"}`, want: map[string]interface{}{"name": "casecheck"}, ok: true},
		{name: "bytes", value: []byte(`[1,2]`), want: []interface{}{float64(1), float64(2)}, ok: true},
		{name: "raw message", value: json.RawMessage(`true`), want: true, ok: true},
		{name: "invalid JSON", value: `{`, ok: false},
		{name: "unsupported type", value: 1, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJSON(tt.value)
			if (err == nil) != tt.ok {
				t.Fatalf("decodeJSON(%#v) error = %v, want success %t", tt.value, err, tt.ok)
			}
			if tt.ok && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeJSON(%#v) = %#v, want %#v", tt.value, got, tt.want)
			}
		})
	}
}
