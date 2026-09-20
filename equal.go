/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

// Equal fails the test unless expected and actual have the same type and value.
func Equal(t testing.TB, expected interface{}, actual interface{}, args ...interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Helper()
		t.Error(errorMessage(args, "Different type\nExpected: %T\nActual: %T", expected, actual))
		t.FailNow()
		return
	}
	if valuesEqual(expected, actual) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Value is not identical\nExpected: %+v\nActual: %+v", expected, actual))
	t.FailNow()
}

// NotEqual fails the test when expected and actual have the same type and value.
func NotEqual(t testing.TB, expected interface{}, actual interface{}, args ...interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Helper()
		t.Error(errorMessage(args, "Different type\nExpected: %T\nActual: %T", expected, actual))
		t.FailNow()
		return
	}
	if !valuesEqual(expected, actual) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Value is not identical\nExpected: %+v\nActual: %+v", expected, actual))
	t.FailNow()
}

func valuesEqual(expected interface{}, actual interface{}) bool {
	return reflect.DeepEqual(expected, actual)
}

// True fails the test when actual is false.
func True(t testing.TB, actual bool, args ...interface{}) {
	if actual {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Want <true>, but got: %+v", actual))
	t.FailNow()
}

// False fails the test when actual is true.
func False(t testing.TB, actual bool, args ...interface{}) {
	if !actual {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Want <false>, but got: %+v", actual))
	t.FailNow()
}

// Contains fails the test unless searchData contains need. Strings and byte
// slices use substring matching; maps are searched by key; slices and arrays
// are searched by element.
func Contains(t testing.TB, searchData interface{}, need interface{}, args ...interface{}) {
	found, supported := contains(searchData, need)
	if !supported {
		t.Helper()
		t.Error(errorMessage(args, "Unsupported types\nSearchData: %T\nNeed: %T", searchData, need))
		t.FailNow()
		return
	}

	if found {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Not found\nSearchData: %+v\nNeed: %+v", searchData, need))
	t.FailNow()
}

// NotContains fails the test when searchData contains need. Strings and byte
// slices use substring matching; maps are searched by key; slices and arrays
// are searched by element.
func NotContains(t testing.TB, searchData interface{}, need interface{}, args ...interface{}) {
	found, supported := contains(searchData, need)
	if !supported {
		t.Helper()
		t.Error(errorMessage(args, "Unsupported types\nSearchData: %T\nNeed: %T", searchData, need))
		t.FailNow()
		return
	}

	if !found {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Found\nSearchData: %+v\nNeed: %+v", searchData, need))
	t.FailNow()
}

func contains(searchData interface{}, need interface{}) (bool, bool) {
	if s1, s2, ok := asString(searchData, need); ok {
		return strings.Contains(s1, s2), true
	}
	if b1, b2, ok := asBytes(searchData, need); ok {
		return bytes.Contains(b1, b2), true
	}

	data := reflect.ValueOf(searchData)
	switch data.Kind() {
	case reflect.Map:
		for _, key := range data.MapKeys() {
			if reflect.DeepEqual(key.Interface(), need) {
				return true, true
			}
		}
		return false, true
	case reflect.Array, reflect.Slice:
		for index := 0; index < data.Len(); index++ {
			if reflect.DeepEqual(data.Index(index).Interface(), need) {
				return true, true
			}
		}
		return false, true
	default:
		return false, false
	}
}

func asString(v0 interface{}, v1 interface{}) (string, string, bool) {
	sv0, ok0 := v0.(string)
	sv1, ok1 := v1.(string)
	return sv0, sv1, ok0 && ok1
}

func asBytes(v0 interface{}, v1 interface{}) ([]byte, []byte, bool) {
	sv0, ok0 := v0.([]byte)
	sv1, ok1 := v1.([]byte)
	return sv0, sv1, ok0 && ok1
}
