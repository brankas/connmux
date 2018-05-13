// Package connmux provides a simple way to multiplex (mux) net connections
// based on content.
package connmux

import "net"

// ConnMux is a connection multiplexor.
type ConnMux struct {
	l net.Listener

	// Default is the default listener.
	Default *Listener
}

// New creates a connection multiplexor for the supplied listener.
func New(l net.Listener, opts ...Option) (*ConnMux, error) {
	var err error

	cm := &ConnMux{
		l: l,
	}

	// apply opts
	for _, o := range opts {
		if err = o(cm); err != nil {
			return nil, err
		}
	}

	// set default listener
	if cm.Default == nil {
		cm.Default = cm.Listen(Any())
	}

	return cm, nil
}

// Listen wraps net.Listen and New, passing the network, address, and options.
func Listen(network, address string, opts ...Option) (*ConnMux, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return New(l, opts...)
}

// Listen creates a listener that matches any of the supplied matchers.
func (cm *ConnMux) Listen(matchers ...Matcher) *Listener {
	l := &Listener{
		cm: cm,
	}

	return l
}

type Listener struct {
	cm *ConnMux
}
