/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// JSONEqual fails the test unless expected and actual contain equivalent JSON.
// Inputs must be strings, byte slices, or json.RawMessage values. Object key
// order and insignificant whitespace do not affect equality.
func JSONEqual(t testing.TB, expected interface{}, actual interface{}, args ...interface{}) {
	expectedValue, expectedErr := decodeJSON(expected)
	actualValue, actualErr := decodeJSON(actual)
	if expectedErr == nil && actualErr == nil && reflect.DeepEqual(expectedValue, actualValue) {
		return
	}
	t.Helper()
	switch {
	case expectedErr != nil:
		t.Error(errorMessage(args, "Invalid expected JSON: %v", expectedErr))
	case actualErr != nil:
		t.Error(errorMessage(args, "Invalid actual JSON: %v", actualErr))
	default:
		t.Error(errorMessage(args, "JSON values are not identical\nExpected: %s\nActual: %s", expected, actual))
	}
	t.FailNow()
}

func decodeJSON(value interface{}) (interface{}, error) {
	var data []byte
	switch input := value.(type) {
	case string:
		data = []byte(input)
	case []byte:
		data = input
	case json.RawMessage:
		data = input
	default:
		return nil, fmt.Errorf("unsupported JSON input type %T", value)
	}

	var decoded interface{}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}
