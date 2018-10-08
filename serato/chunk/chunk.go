package chunk

import "errors"

// ErrUnexpectedIdentifier is an error representing that a constructor received
// the wrong chunk identifier for the chunk type being created.
var ErrUnexpectedIdentifier = errors.New("unexpected chunk identifier")

// Chunk ...
type Chunk interface {
	Header() *Header
	Type() string
}
