package chunk

import "io"

// Uent ...
type Uent struct {
	header *Header
	data   []byte
}

// Header returns the header of the chunk.
func (u *Uent) Header() *Header {
	return u.header
}

// Type returns the type of the chunk.
func (u *Uent) Type() string {
	return u.header.Type()
}

// NewUentChunk ...
func NewUentChunk(header *Header, r io.Reader) (*Uent, error) {
	return nil, nil
}
