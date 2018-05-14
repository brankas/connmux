package netmux

import (
	"bytes"
)

// radixNode is a simple radix (patricia) tree node.
type radixNode struct {
	// prefix is the prefix for this node.
	prefix []byte

	// next are any successive prefixes.
	next map[byte]*radixNode

	// terminal indicates the end of a prefix.
	terminal bool
}

// radixTree is a simple radix (patricia) tree.
type radixTree struct {
	radixNode
	max int
}

// newRadixTree creates a radix (patricia) tree for the supplied prefixes.
func newRadixTree(prefixes ...[]byte) *radixTree {
	var max int
	for _, prefix := range prefixes {
		if n := len(prefix); max < n {
			max = n
		}
	}
	return &radixTree{*newRadixNode(prefixes), max + 1}
}

// newRadixTreeString creates a radix (patricia) tree for the supplied
// prefixes.
func newRadixTreeString(prefixStrings ...string) *radixTree {
	prefixes := make([][]byte, len(prefixStrings))
	for i, prefix := range prefixStrings {
		prefixes[i] = []byte(prefix)
	}
	return newRadixTree(prefixes...)
}

// match recursively searches for buf, returning true when the end of buf
// matches a terminal node of the tree or begins with a prefix in the tree.
//
// When prefix is true, and match has reached a node with no
// children, returns true if buf begins with the specified prefix.
//
// Matches prefixes only when prefix is true.
func (rt *radixTree) match(buf []byte, prefix bool) bool {
	return rt.radixNode.match(buf[:max(len(buf), rt.max)], prefix)
}

// matchPrefix determines if r matches a prefix in the tree.
// newRadixNode creates a new radix (patricia) tree node.
func newRadixNode(prefixes [][]byte) *radixNode {
	if len(prefixes) == 0 {
		return nil
	}

	if len(prefixes) == 1 {
		return &radixNode{
			prefix:   prefixes[0],
			terminal: true,
		}
	}

	var prefix []byte
	prefix, prefixes = splitPrefix(prefixes)
	rn := &radixNode{
		prefix: prefix,
	}

	next := make(map[byte][][]byte)
	for _, prefix := range prefixes {
		if len(prefix) == 0 {
			rn.terminal = true
			continue
		}
		next[prefix[0]] = append(next[prefix[0]], prefix[1:])
	}

	rn.next = make(map[byte]*radixNode)
	for first, rest := range next {
		rn.next[first] = newRadixNode(rest)
	}

	return rn
}

// splitPrefix splits prefixes.
func splitPrefix(prefixes [][]byte) ([]byte, [][]byte) {
	if len(prefixes) == 0 || len(prefixes[0]) == 0 {
		return nil, prefixes
	}

	if len(prefixes) == 1 {
		return prefixes[0], [][]byte{{}}
	}

	var prefix []byte
	for i := 0; ; i++ {
		var cur byte
		eq := true
		for j, p := range prefixes {
			if len(p) <= i {
				eq = false
				break
			}

			if j == 0 {
				cur = p[i]
				continue
			}

			if cur != p[i] {
				eq = false
				break
			}
		}

		if !eq {
			break
		}

		prefix = append(prefix, cur)
	}

	rest := make([][]byte, 0, len(prefix))
	for _, b := range prefixes {
		rest = append(rest, b[len(prefix):])
	}

	return prefix, rest
}

// match recursively searches for buf, returning true when the end of buf
// matches a terminal node of the tree or begins with a prefix in the tree.
//
// When prefix is true, and match has reached a node with no
// children, returns true if buf begins with the specified prefix.
//
// Matches prefixes only when prefix is true.
func (rn *radixNode) match(buf []byte, prefix bool) bool {
	l := len(rn.prefix)
	if l > 0 {
		if l > len(buf) {
			l = len(buf)
		}
		if !bytes.Equal(buf[:l], rn.prefix) {
			return false
		}
	}

	if rn.terminal && (prefix || len(rn.prefix) == len(buf)) {
		return true
	}

	if l >= len(buf) {
		return false
	}

	next, ok := rn.next[buf[l]]
	if !ok {
		return false
	}

	if l == len(buf) {
		buf = buf[l:l]
	} else {
		buf = buf[l+1:]
	}

	return next.match(buf, prefix)
}

// max returns the maximum of a, b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
