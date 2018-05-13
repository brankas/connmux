package connmux

// Error is a connection mux error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// Error values.
const (
	// ErrListenerClosed is the listener closed error.
	ErrListenerClosed Error = "listener closed"
)
