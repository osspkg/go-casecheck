/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"testing"
	"time"
)

// WithinDuration fails the test unless expected and actual differ by at most
// delta. A negative delta never matches.
func WithinDuration(t testing.TB, expected time.Time, actual time.Time, delta time.Duration, args ...interface{}) {
	if withinDuration(expected, actual, delta) {
		return
	}
	t.Helper()
	t.Error(errorMessage(args, "Times differ by more than delta\nExpected: %s\nActual: %s\nDelta: %s", expected, actual, delta))
	t.FailNow()
}

func withinDuration(expected time.Time, actual time.Time, delta time.Duration) bool {
	return delta >= 0 && expected.Sub(actual).Abs() <= delta
}
