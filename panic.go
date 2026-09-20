/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import "testing"

// Panics fails the test unless fn panics.
func Panics(t testing.TB, fn func(), args ...interface{}) {
	if panics(fn) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Expected panic, but function completed normally"))
	t.FailNow()
}

// NotPanics fails the test when fn panics.
func NotPanics(t testing.TB, fn func(), args ...interface{}) {
	if !panics(fn) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Unexpected panic"))
	t.FailNow()
}

func panics(fn func()) (didPanic bool) {
	completed := false
	defer func() {
		didPanic = !completed
		_ = recover()
	}()
	fn()
	completed = true
	return false
}
