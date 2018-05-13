// Package netmux provides a simple way to multiplex (mux) net connections
// based on content.
package netmux

import (
	"net"
)

// Netmux is a connection multiplexor.
type Netmux struct {
	l net.Listener

	// Default is the default listener.
	Default *Listener
}

// New creates a connection multiplexor for the supplied listener.
func New(l net.Listener, opts ...Option) (*Netmux, error) {
	var err error

	nm := &Netmux{
		l: l,
	}

	// apply opts
	for _, o := range opts {
		if err = o(nm); err != nil {
			return nil, err
		}
	}

	// set default listener
	if nm.Default == nil {
		nm.Default = nm.Listen(Any())
	}

	return nm, nil
}

// Listen wraps net.Listen and New, passing the network, address, and options.
func Listen(network, address string, opts ...Option) (*Netmux, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return New(l, opts...)
}

// Listen creates a listener that matches any of the supplied matchers.
func (nm *Netmux) Listen(matchers ...Matcher) *Listener {
	l := &Listener{
		nm: nm,
	}

	return l
}
