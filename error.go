package zinc

import "fmt"

// UnknownIndexError indicates that the given index is not known to the server.
type UnknownIndexError struct {
	index string
}

// NewUnknownIndexError returns a new UnknownIndexError.
func NewUnknownIndexError(name string) UnknownIndexError {
	return UnknownIndexError{
		index: name,
	}
}

// Error implements the errors interface.
func (u UnknownIndexError) Error() string {
	return fmt.Sprintf("unknown index with name %q", u.index)
}
