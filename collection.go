/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"reflect"
	"testing"
)

// Empty fails the test unless value is nil or an empty string, slice, map,
// array, or channel.
func Empty(t testing.TB, value interface{}, args ...interface{}) {
	empty, supported := isEmpty(value)
	if supported && empty {
		return
	}
	t.Helper()
	if !supported {
		t.Error(errorMessage(args, "Unsupported type: %T", value))
	} else {
		t.Error(errorMessage(args, "Want empty value, but got: %+v", value))
	}
	t.FailNow()
}

// NotEmpty fails the test unless value is a non-empty string, slice, map,
// array, or channel.
func NotEmpty(t testing.TB, value interface{}, args ...interface{}) {
	empty, supported := isEmpty(value)
	if supported && !empty {
		return
	}
	t.Helper()
	if !supported {
		t.Error(errorMessage(args, "Unsupported type: %T", value))
	} else {
		t.Error(errorMessage(args, "Want non-empty value, but got: %+v", value))
	}
	t.FailNow()
}

// Len fails the test unless value has the expected length. It supports strings,
// slices, maps, arrays, and channels.
func Len(t testing.TB, value interface{}, want int, args ...interface{}) {
	got, supported := length(value)
	if supported && got == want {
		return
	}
	t.Helper()
	if !supported {
		t.Error(errorMessage(args, "Unsupported type: %T", value))
	} else {
		t.Error(errorMessage(args, "Unexpected length\nGot: %d\nWant: %d", got, want))
	}
	t.FailNow()
}

// ElementsMatch fails the test unless expected and actual contain the same
// elements with the same multiplicities, regardless of order. It supports
// slices and arrays; nil and empty collections are equivalent when both have
// zero length.
func ElementsMatch(t testing.TB, expected interface{}, actual interface{}, args ...interface{}) {
	matched, supported := elementsMatch(expected, actual)
	if supported && matched {
		return
	}
	t.Helper()
	if !supported {
		t.Error(errorMessage(args, "Unsupported types\nExpected: %T\nActual: %T", expected, actual))
	} else {
		t.Error(errorMessage(args, "Elements are not identical\nExpected: %+v\nActual: %+v", expected, actual))
	}
	t.FailNow()
}

func isEmpty(value interface{}) (bool, bool) {
	if value == nil {
		return true, true
	}
	length, supported := length(value)
	return length == 0, supported
}

func length(value interface{}) (int, bool) {
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return reflectValue.Len(), true
	default:
		return 0, false
	}
}

func elementsMatch(expected interface{}, actual interface{}) (bool, bool) {
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)
	if !isSliceOrArray(expectedValue) || !isSliceOrArray(actualValue) {
		return false, false
	}
	if expectedValue.Type().Elem() != actualValue.Type().Elem() || expectedValue.Len() != actualValue.Len() {
		return false, true
	}

	matched := make([]bool, actualValue.Len())
	for expectedIndex := 0; expectedIndex < expectedValue.Len(); expectedIndex++ {
		found := false
		for actualIndex := 0; actualIndex < actualValue.Len(); actualIndex++ {
			if !matched[actualIndex] && reflect.DeepEqual(expectedValue.Index(expectedIndex).Interface(), actualValue.Index(actualIndex).Interface()) {
				matched[actualIndex] = true
				found = true
				break
			}
		}
		if !found {
			return false, true
		}
	}
	return true, true
}

func isSliceOrArray(value reflect.Value) bool {
	return value.IsValid() && (value.Kind() == reflect.Array || value.Kind() == reflect.Slice)
}
