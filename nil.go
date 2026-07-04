/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"reflect"
	"testing"
)

func Nil(t testing.TB, actual interface{}, args ...interface{}) {
	if isNil(actual) {
		return
	}
	t.Helper()
	t.Errorf(errorMessage(args, "Want <nil>, but got %+v", actual))
	t.FailNow()
}

func NotNil(t testing.TB, actual interface{}, args ...interface{}) {
	if !isNil(actual) {
		return
	}
	t.Helper()
	t.Errorf(errorMessage(args, "Want not <nil>, but got %+v", actual))
	t.FailNow()
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}
	return reflect.ValueOf(value).IsNil()
}
