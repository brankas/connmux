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

func TestRadixTreeMatch(t *testing.T) {
	return
	tests := []struct {
		prefixes []string
		val      string
		prefix   bool
		exp      bool
	}{
		{nil, "a", false, false}, // 0
		{nil, "", false, false},
		{nil, "a", true, false},
		{[]string{}, "", false, false},
		{[]string{}, "", true, false},

		{[]string{""}, "", false, false}, // 5
		{[]string{""}, "", true, false},

		{[]string{"a"}, "", false, false}, // 7
		{[]string{"a"}, "", true, false},
		{[]string{"a", "aa"}, "", false, false},
		{[]string{"a", "aa"}, "", true, false},

		{[]string{"a"}, "a", false, true}, // 11
		{[]string{"a"}, "a", true, true},
		{[]string{"b", "aa"}, "a", false, true},
		{[]string{"b", "aa"}, "a", true, true},
	}
	for i, test := range tests {
		rt := newRadixTreeString(test.prefixes...)
		res := rt.match([]byte(test.val), test.prefix)
		if res != test.exp {
			t.Errorf("test %d for %v radixTree.match(%q, %t) should be %t", i, test.prefixes, test.val, test.prefix, test.exp)
		}
	}
}
