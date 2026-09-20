package casecheck

import "testing"

func TestUnit_Match(t *testing.T) {
	Match(t, "casecheck", "^case")
	NotMatch(t, "casecheck", "^test")
}
