package chunk

import "io"

// Oren ...
type Oren struct {
	header *Header
	data   []byte

	Uent *Uent
}

// Header returns the header of the chunk.
func (o *Oren) Header() *Header {
	return o.header
}

// Type returns the type of the chunk.
func (o *Oren) Type() string {
	return o.header.Type()
}

// NewOrenChunk ...
func NewOrenChunk(header *Header, r io.Reader) (*Oren, error) {
	return nil, nil
}
