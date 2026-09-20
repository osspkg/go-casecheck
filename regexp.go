/*
 *  Copyright (c) 2024-2026 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package casecheck

import (
	"regexp"
	"testing"
)

// Match fails the test unless value matches pattern.
func Match(t testing.TB, value string, pattern string, args ...interface{}) {
	matched, err := regexp.MatchString(pattern, value)
	if err == nil && matched {
		return
	}
	t.Helper()
	if err != nil {
		t.Error(errorMessage(args, "Invalid regular expression %q: %v", pattern, err))
	} else {
		t.Error(errorMessage(args, "Value does not match pattern\nValue: %q\nPattern: %q", value, pattern))
	}
	t.FailNow()
}

// NotMatch fails the test when value matches pattern.
func NotMatch(t testing.TB, value string, pattern string, args ...interface{}) {
	matched, err := regexp.MatchString(pattern, value)
	if err == nil && !matched {
		return
	}
	t.Helper()
	if err != nil {
		t.Error(errorMessage(args, "Invalid regular expression %q: %v", pattern, err))
	} else {
		t.Error(errorMessage(args, "Value matches pattern\nValue: %q\nPattern: %q", value, pattern))
	}
	t.FailNow()
}
