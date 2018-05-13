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

// Listen creates a listener that matches any of the supplied matchers.
func (cm *ConnMux) Listen(matchers ...Matcher) *Listener {
	l := &Listener{
		cm: cm,
	}

	return l
}

// Listener
type Listener struct {
	cm *ConnMux
}

// Accept satisfies the net.Listener interface.
func (l *Listener) Accept() (net.Conn, error) {
	return nil, nil
}

// Close satisfies the net.Listener interface.
func (l *Listener) Close() error {
	return nil
}

// Addr satisfies the net.Listener interface.
func (l *Listener) Addr() net.Addr {
	return nil
}
