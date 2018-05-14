package netmux

import (
	"testing"
)

func TestRadixTreeMatch(t *testing.T) {
	tests := []struct {
		prefixes []string
		test     string
		prefix   bool
		exp      bool
	}{
		{nil, "a", false, false},
		{nil, "", false, true},
		{nil, "a", true, true},
		{[]string{}, "", false, true},
		{[]string{}, "", true, true},
	}
	for i, test := range tests {
		i, test = i, test
	}
}
