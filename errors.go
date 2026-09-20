/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

var errorType = reflect.TypeFor[error]()

// NoError fails the test when err is not nil.
func NoError(t testing.TB, err error, args ...interface{}) {
	if err == nil {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Want <nil>, but got error: %+v", err.Error()))
	t.FailNow()
}

// Error fails the test when err is nil.
func Error(t testing.TB, err error, args ...interface{}) {
	if err != nil {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Want error, but got <nil>"))
	t.FailNow()
}

// ErrorContains fails the test unless err contains need in its message.
func ErrorContains(t testing.TB, err error, need string, args ...interface{}) {
	Error(t, err, args...)

	if strings.Contains(err.Error(), need) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Not found\nSearchData: %+v\nNeed: %+v", err.Error(), need))
	t.FailNow()
}

// ErrorIs fails the test unless err matches target through errors.Is.
func ErrorIs(t testing.TB, err error, target error, args ...interface{}) {
	if errors.Is(err, target) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Error does not match target\nError: %+v\nTarget: %+v", err, target))
	t.FailNow()
}

// ErrorAs fails the test unless err contains a value assignable to target.
// target must be a non-nil pointer to an error type or an interface type.
func ErrorAs(t testing.TB, err error, target interface{}, args ...interface{}) {
	matched, valid := errorAs(err, target)
	if valid && matched {
		return
	}
	t.Helper()
	if !valid {
		t.Error(errorMessage(args, "Invalid error target: %T", target))
	} else {
		t.Error(errorMessage(args, "Error does not match target\nError: %+v\nTarget: %T", err, target))
	}
	t.FailNow()
}

func errorAs(err error, target interface{}) (bool, bool) {
	if target == nil {
		return false, false
	}

	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return false, false
	}

	targetType := value.Elem().Type()
	if targetType.Kind() != reflect.Interface && !targetType.Implements(errorType) {
		return false, false
	}

	return errors.As(err, target), true
}
