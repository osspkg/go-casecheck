/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"math"
	"testing"
)

// InDelta fails the test unless expected and actual differ by at most delta.
// It supports float32 and float64 values. A negative delta never matches.
func InDelta[T ~float32 | ~float64](t testing.TB, expected T, actual T, delta T, args ...interface{}) {
	if inDelta(expected, actual, delta) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Values differ by more than delta\nExpected: %+v\nActual: %+v\nDelta: %+v", expected, actual, delta))
	t.FailNow()
}

func inDelta[T ~float32 | ~float64](expected T, actual T, delta T) bool {
	if delta < 0 {
		return false
	}
	if expected == actual {
		return true
	}
	return math.Abs(float64(expected)-float64(actual)) <= float64(delta)
}
