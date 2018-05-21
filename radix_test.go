package netmux

import (
	"testing"
)

func TestRadixTree(t *testing.T) {
	strs := []string{"foo", "far", "farther", "boo", "ba", "bar"}
	pt := newRadixTreeString(strs...)
	for _, s := range strs {
		if !pt.match([]byte(s), false) {
			t.Errorf("%s is not matched by %s", s, s)
		}

		if !pt.match([]byte(s+s), true) {
			t.Errorf("%s is not matched as a prefix by %s", s+s, s)
		}

		if pt.match([]byte(s+s), false) {
			t.Errorf("%s matches %s", s+s, s)
		}

		// The following tests are just to catch index out of
		// range and off-by-one errors and not the functionality.
		pt.match([]byte(s[:len(s)-1]), true)
		pt.match([]byte(s[:len(s)-1]), false)
		pt.match([]byte(s+"$"), true)
		pt.match([]byte(s+"$"), false)
	}
}
